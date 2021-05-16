package data

import (
  log "github.com/sirupsen/logrus"
)

type todoitemRepo struct {
  data *Data
}

func NewTodoRepo(data *Data) biz.TodoItemRepo {
  return &todoitemRepo {
           data: data
  } 
}

func (to *todoitemRepo) CreateTodoItem(todoitem *biz.TodoItem) error {
  _, err := ar.data.db..Create(todoitem)
}
