package biz 


type TodoItem struct{
  Id int `gorm:"primary_key"`
  Description string
  Completed bool
}

type TodoItemRepo interface {
  createItem(todoitem *TodoItem) error
}

type TodoUsecase struct {
  repo TodoItemRepo 
}

func NewTodoUsecase(repo TotoItemRepo ) *TodoUsecase {
  return &TodoUsecase{ repo :repo}
}

func (uc *TodoUsecase) Create ( todoitem *TodoItem ) error {
  return uc.repo.CreateItem(todoitem)
}
