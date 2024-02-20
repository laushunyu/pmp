package hello

import (
	"context"
	"github.com/laushunyu/pmp/internal/model"
	"github.com/laushunyu/pmp/internal/service"

	"github.com/laushunyu/pmp/api/task/v1"
)

func (c *ControllerV1) UpdateTask(ctx context.Context, req *v1.UpdateTaskReq) (res *v1.UpdateTaskRes, err error) {
	if _, err := service.Pmp().UpdateTask(ctx, &model.UpdateTaskIn{
		TaskId:   req.TaskId,
		WorkedAt: req.WorkedAt,
		WorkerId: req.WorkerId,
		Detail:   req.Detail,
		StateAct: req.StateAct,
	}); err != nil {
		return nil, err
	}

	return new(v1.UpdateTaskRes), nil
}
