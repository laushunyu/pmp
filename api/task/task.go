// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package task

import (
	"context"

	"github.com/laushunyu/pmp/api/task/v1"
)

type ITaskV1 interface {
	CreateTask(ctx context.Context, req *v1.CreateTaskReq) (res *v1.CreateTaskRes, err error)
	UpdateTask(ctx context.Context, req *v1.UpdateTaskReq) (res *v1.UpdateTaskRes, err error)
}
