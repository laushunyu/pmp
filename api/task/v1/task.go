package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/laushunyu/pmp/internal/consts"
)

type CreateTaskReq struct {
	g.Meta `path:"/task" tags:"PMP" method:"post" summary:"创建任务"`

	WorkerId  string      `json:"workerId" ` // 员工 ID
	CreatedAt *gtime.Time `json:"createdAt"` // 创建时间
	Title     string      `json:"title"    ` // 任务标题
	Url       string      `json:"url"      ` // 任务链接
	Sprint    string      `json:"sprint"   ` // 任务开发 SPRINT
	ParentId  string      `json:"parentId" ` // 主任务 ID
}

type CreateTaskRes struct {
	g.Meta `mime:"application/json"`

	Id string `json:"id"` // 任务 ID
}

type UpdateTaskReq struct {
	g.Meta `path:"/task" tags:"PMP" method:"put" summary:"更新任务状态"`

	TaskId   string           `json:"taskId"  ` // 任务 ID
	WorkedAt *gtime.Time      `json:"workedAt"` // 工作日期
	WorkerId string           `json:"workerId"` // 员工 ID
	Detail   string           `json:"detail"  ` // 工作内容
	StateAct consts.TaskState `json:"stateAct"` // 状态轮转
}

type UpdateTaskRes struct {
	g.Meta `mime:"application/json"`
}

type GetTaskSummaryReq struct {
	g.Meta `path:"/task/summary" tags:"PMP" method:"get" summary:"获取任务概要"`

	Sprint string `json:"sprint"`
}

type GetTaskSummaryRes struct {
	g.Meta `mime:"application/json"`

	Gantt GetTaskSummaryGantt
}

type GetTaskSummaryGantt struct {
}

type GetTaskSummaryGanttTask struct {
}
