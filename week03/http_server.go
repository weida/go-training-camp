package main

import (
	"context"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)


var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {

	g, ctx := errgroup.WithContext(context.Background())

	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
        srv := &http.Server{Addr: ":1718"}

        g.Go(func() error {
	    sig := make(chan os.Signal)
            signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
            select {
              case <-ctx.Done():
                  fmt.Printf("ctx done 1...\n")
                  return ctx.Err()
	      case  <-sig:
                  fmt.Printf("sig ...%+v\n", sig)
                  return srv.Shutdown(ctx)
            }
        })

        g.Go(func() error{
           <-ctx.Done()
           fmt.Printf("ctx done 2...\n")
           return srv.Shutdown(ctx)
        })

        g.Go(func() error{
	    return  srv.ListenAndServe()
        })
        

       g.Wait()
}

func QR(w http.ResponseWriter, req *http.Request) {
	templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET">
    <input maxLength=1024 size=70 name=s value="" title="Text to QR Encode">
    <input type=submit value="Show QR" name=qr>
</form>
</body>
</html>
`

