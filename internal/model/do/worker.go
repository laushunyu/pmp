// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Worker is the golang structure of table worker for DAO operations like Where/Data.
type Worker struct {
	g.Meta `orm:"table:worker, do:true"`
	Id     interface{} // 员工 ID
	Name   interface{} // 员工名
}
