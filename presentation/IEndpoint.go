package presentation

type IRequest struct{}

type IResponse struct {
	Status int
	Body   interface{}
}

type IEndpoint interface {
	Handle(IRequest) *IResponse
}
