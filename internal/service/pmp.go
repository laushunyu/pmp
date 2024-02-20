// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"github.com/laushunyu/pmp/internal/model"
)

type (
	IPmp interface {
		CreateTask(ctx context.Context, in *model.CreateTaskIn) (*model.CreateTaskOut, error)
		UpdateTask(ctx context.Context, in *model.UpdateTaskIn) (*model.UpdateTaskOut, error)
	}
)

var (
	localPmp IPmp
)

func Pmp() IPmp {
	if localPmp == nil {
		panic("implement not found for interface IPmp, forgot register?")
	}
	return localPmp
}

func RegisterPmp(i IPmp) {
	localPmp = i
}
