package service 

func NewTodoService(todo *biz.TodoUsecase) *TodoService {
  return &TodoService {
    todo : todo
  }
}

func (s *TodoService) CreateTodoService(w *http.ResponseWriter, req *http.Request)  {
   err := s.article.Create(&biz.TodoItem {
              Id : req.Id
              Description : req.Description  
              Completed : req.Completed
   })
   w.Header().Set("Content-Type", "application/json")
   json.NewEncoder(w).Encode(result.Value)
}
