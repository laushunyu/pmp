// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskRecordDao is the data access object for table task_record.
type TaskRecordDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns TaskRecordColumns // columns contains all the column names of Table for convenient usage.
}

// TaskRecordColumns defines and stores column names for table task_record.
type TaskRecordColumns struct {
	TaskId   string // 任务 ID
	WorkedAt string // 工作日期
	WorkerId string // 员工 ID
	Detail   string // 工作内容
	StateAct string // 状态轮转，如不设置则不进行轮转
}

// taskRecordColumns holds the columns for table task_record.
var taskRecordColumns = TaskRecordColumns{
	TaskId:   "task_id",
	WorkedAt: "worked_at",
	WorkerId: "worker_id",
	Detail:   "detail",
	StateAct: "state_act",
}

// NewTaskRecordDao creates and returns a new DAO object for table data access.
func NewTaskRecordDao() *TaskRecordDao {
	return &TaskRecordDao{
		group:   "default",
		table:   "task_record",
		columns: taskRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskRecordDao) Columns() TaskRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
