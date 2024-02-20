package hack

import "time"

type Task struct {
	ID         string    // 任务 ID
	WorkerID   string    // 员工 ID
	CreatedAt  time.Time // 创建时间
	FinishedAt time.Time // 结束时间
	Title      string    // 任务标题
	Url        string    // 任务链接
	Sprint     string    // 任务开发 SPRINT
	MainTaskID string    // 主任务 ID
}

type TaskRecord struct {
	TaskID   string    // 任务 ID
	WorkAt   time.Time // 工作日期
	WorkerID string    // 员工 ID
	Detail   string    // 工作内容
}

type Worker struct {
	ID   string // 员工 ID
	Name string // 员工名
}
