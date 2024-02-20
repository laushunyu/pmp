package hello

import (
	"context"
	"github.com/laushunyu/pmp/internal/model"
	"github.com/laushunyu/pmp/internal/service"

	"github.com/laushunyu/pmp/api/task/v1"
)

func (c *ControllerV1) CreateTask(ctx context.Context, req *v1.CreateTaskReq) (res *v1.CreateTaskRes, err error) {
	in := &model.CreateTaskIn{
		WorkerId:  req.WorkerId,
		CreatedAt: req.CreatedAt,
		Title:     req.Title,
		Url:       req.Url,
		Sprint:    req.Sprint,
		ParentId:  req.ParentId,
	}
	if _, err := service.Pmp().CreateTask(ctx, in); err != nil {
		return nil, err
	}
	return new(v1.CreateTaskRes), nil
}
