package service 

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewtodoService)

type TodoService struct {
	pb.UnimplementedBlogServiceServer

	log *log.Helper

	article *biz.TodoUsecase
}
