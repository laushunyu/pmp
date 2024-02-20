package pmp

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
	"github.com/laushunyu/pmp/internal/consts"
	"github.com/laushunyu/pmp/internal/dao"
	"github.com/laushunyu/pmp/internal/model"
	"github.com/laushunyu/pmp/internal/model/entity"
	"github.com/laushunyu/pmp/internal/service"
)

func init() {
	service.RegisterPmp(&sPmp{})
}

type sPmp struct {
}

func (p *sPmp) CreateTask(ctx context.Context, in *model.CreateTaskIn) (*model.CreateTaskOut, error) {
	_, err := dao.Task.Ctx(ctx).Insert(&entity.Task{
		Id:        uuid.New().String(),
		WorkerId:  in.WorkerId,
		CreatedAt: gtime.Now(),
		Title:     in.Title,
		Url:       in.Url,
		Sprint:    in.Sprint,
		ParentId:  in.ParentId,
		State:     string(consts.TaskStateTodo),
	})
	return new(model.CreateTaskOut), err
}

func (p *sPmp) UpdateTask(ctx context.Context, in *model.UpdateTaskIn) (*model.UpdateTaskOut, error) {
	var task entity.Task
	if err := dao.Task.Ctx(ctx).Where(dao.Task.Columns().Id, in.TaskId).Scan(&task); err != nil {
		return nil, err
	}

	if _, err := dao.TaskRecord.Ctx(ctx).Insert(&entity.TaskRecord{
		TaskId:   in.TaskId,
		WorkedAt: in.WorkedAt,
		WorkerId: in.WorkerId,
		Detail:   in.Detail,
		StateAct: string(in.StateAct),
	}); err != nil {
		return nil, err
	}

	updateColumn := make(g.Map)
	switch in.StateAct {
	case consts.TaskStateTodo:
		updateColumn[dao.Task.Columns().State] = in.StateAct
	case consts.TaskStateWaitDesign:
		updateColumn[dao.Task.Columns().State] = in.StateAct
	case consts.TaskStateInProgress:
		updateColumn[dao.Task.Columns().State] = in.StateAct
	case consts.TaskStateWaitMerge:
		updateColumn[dao.Task.Columns().State] = in.StateAct
	case consts.TaskStateDone:
		updateColumn[dao.Task.Columns().State] = in.StateAct
		updateColumn[dao.Task.Columns().FinishedAt] = gtime.Now()
	}

	if _, err := dao.Task.Ctx(ctx).Where(dao.Task.Columns().Id, in.TaskId).Update(updateColumn); err != nil {
		return nil, err
	}

	return new(model.UpdateTaskOut), nil
}
