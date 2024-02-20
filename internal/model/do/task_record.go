// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskRecord is the golang structure of table task_record for DAO operations like Where/Data.
type TaskRecord struct {
	g.Meta   `orm:"table:task_record, do:true"`
	TaskId   interface{} // 任务 ID
	WorkedAt *gtime.Time // 工作日期
	WorkerId interface{} // 员工 ID
	Detail   interface{} // 工作内容
	StateAct interface{} // 状态轮转，如不设置则不进行轮转
}
