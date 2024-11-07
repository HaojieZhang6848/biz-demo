package service

import (
	"context"
	pbapi "github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi"
)

type ByeService struct {
	ctx context.Context
} // NewByeService new ByeService
func NewByeService(ctx context.Context) *ByeService {
	return &ByeService{ctx: ctx}
}

// Run create note info
func (s *ByeService) Run(req *pbapi.Request) (resp *pbapi.Response, err error) {
	// Finish your business logic.
	return &pbapi.Response{Message: "Bye " + req.Message}, nil
}
