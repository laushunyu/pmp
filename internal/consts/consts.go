package consts

type TaskState string

const (
	TaskStateTodo       TaskState = "todo"
	TaskStateWaitDesign TaskState = "wait_design"
	TaskStateInProgress TaskState = "in_progress"
	TaskStateWaitMerge  TaskState = "wait_merge"
	TaskStateDone       TaskState = "done"
)
