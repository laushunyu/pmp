// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskRecord is the golang structure for table task_record.
type TaskRecord struct {
	TaskId   string      `json:"taskId"   ` // 任务 ID
	WorkedAt *gtime.Time `json:"workedAt" ` // 工作日期
	WorkerId string      `json:"workerId" ` // 员工 ID
	Detail   string      `json:"detail"   ` // 工作内容
	StateAct string      `json:"stateAct" ` // 状态轮转，如不设置则不进行轮转
}
