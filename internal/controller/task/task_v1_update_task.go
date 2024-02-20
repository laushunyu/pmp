package task

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/laushunyu/pmp/api/task/v1"
)

func (c *ControllerV1) UpdateTask(ctx context.Context, req *v1.UpdateTaskReq) (res *v1.UpdateTaskRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
