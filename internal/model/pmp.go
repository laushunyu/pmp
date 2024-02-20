package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/laushunyu/pmp/internal/consts"
)

type CreateTaskIn struct {
	WorkerId  string      `json:"workerId"   ` // 员工 ID
	CreatedAt *gtime.Time `json:"createdAt"  ` // 创建时间
	Title     string      `json:"title"      ` // 任务标题
	Url       string      `json:"url"        ` // 任务链接
	Sprint    string      `json:"sprint"     ` // 任务开发 SPRINT
	ParentId  string      `json:"parentId"   ` // 主任务 ID
}

type CreateTaskOut struct {
}

type UpdateTaskIn struct {
	TaskId   string           `json:"taskId"   ` // 任务 ID
	WorkedAt *gtime.Time      `json:"workedAt" ` // 工作日期
	WorkerId string           `json:"workerId" ` // 员工 ID
	Detail   string           `json:"detail"   ` // 工作内容
	StateAct consts.TaskState `json:"stateAct" ` // 状态轮转
}

type UpdateTaskOut struct {
}
