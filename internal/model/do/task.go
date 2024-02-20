// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Task is the golang structure of table task for DAO operations like Where/Data.
type Task struct {
	g.Meta     `orm:"table:task, do:true"`
	Id         interface{} // 任务 ID
	WorkerId   interface{} // 员工 ID
	CreatedAt  *gtime.Time // 创建时间
	FinishedAt *gtime.Time // 结束时间
	Title      interface{} // 任务标题
	Url        interface{} // 任务链接
	Sprint     interface{} // 任务开发 SPRINT
	ParentId   interface{} // 主任务 ID
	State      interface{} // 任务状态
}
