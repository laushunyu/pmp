// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Task is the golang structure for table task.
type Task struct {
	Id         string      `json:"id"         ` // 任务 ID
	WorkerId   string      `json:"workerId"   ` // 员工 ID
	CreatedAt  *gtime.Time `json:"createdAt"  ` // 创建时间
	FinishedAt *gtime.Time `json:"finishedAt" ` // 结束时间
	Title      string      `json:"title"      ` // 任务标题
	Url        string      `json:"url"        ` // 任务链接
	Sprint     string      `json:"sprint"     ` // 任务开发 SPRINT
	ParentId   string      `json:"parentId"   ` // 主任务 ID
	State      string      `json:"state"      ` // 任务状态
}
