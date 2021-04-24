package goerror_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func query() error {
	return sql.ErrNoRows
}

func handle() error {
	// do nothing
	return query()
}

// 1. 如果仅仅需要处理结果可以这样处理
func TestHandle(t *testing.T) {
	err := handle()
	if err != nil {
		fmt.Printf(" handle err,%+v\n", err)
	}
}

//输出:  handle err,sql: no rows in result se

// 2. 有时要明确一些错误信息，某些应用场景下
//    这种"错误"反应的始终状态, 不需要进行错误处理
func TestHandle2(t *testing.T) {
	err := handle()
	if err == sql.ErrNoRows {
		fmt.Printf(" data not found, ,%+v\n", err)
		return
	}
	if err != nil {
	}
}

//输出: data not found, ,sql: no rows in result set

// 3. 有时要想在执行sql的函数中打印一下详细信息
//    方便出错时处理, 这时发现错误已经无法匹配了
func query3() error {
	return fmt.Errorf("query error %v, sql detail: ...", sql.ErrNoRows)
}

func handle3() error {
	// do nothing
	return query3()
}

func TestHandle3(t *testing.T) {
	err := handle3()
	if err == sql.ErrNoRows {
		fmt.Printf(" data not found, ,%+v\n", err)
		return
	}
	if err != nil {
	}
}

//输出： 无

// 4. 解决方式： github.com/pkg/errors 他有三个方法
//     1)Wrap用于包装基础错误，添加上下文文本信息以及附加调用堆栈。
//       通常，它用于包装其他人（标准库或第三方库）对API的调用。
//     2)WithMessage用于在不附加调用堆栈的情况下将上下文文本信息添加到基础错误中。
//       仅将此方法应用于“包装的错误”。 注意：不要重复Wrap，它会记录冗余调用堆栈
//     3)Cause用于确定潜在的错误

func query4() error {
	return errors.Wrap(sql.ErrNoRows, "query error , sql detail: ...")
}
func handle4() error {
	// do nothing
	return errors.WithMessage(query4(), "handle error")
}

func TestHandle4(t *testing.T) {
	err := handle4()
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf(" data not found, ,%+v\n", err)
		return
	}
	if err != nil {
	}
}

//输出：成功定位到了错误，打印了错误信息及相关调用堆栈
/* data not found, ,sql: no rows in result set
query error , sql detail: ...
golesson/week02_test.query4
	f:/GO/golesson/week02/goerror_test.go:75
golesson/week02_test.handle4
	f:/GO/golesson/week02/goerror_test.go:79
golesson/week02_test.TestHandle4
	f:/GO/golesson/week02/goerror_test.go:83
testing.tRunner
	C:/Program Files/Go/src/testing/testing.go:1194
runtime.goexit
	C:/Program Files/Go/src/runtime/asm_amd64.s:1371
handle error
*/
