import "time"

type Todo struct {
	ID        int `gorm:"primaryKey"`
	Task      string
	Status    TaskStatus
	CreatedAt time.Time `gorm:"<-:false"`
	UpdatedAt time.Time `gorm:"<-:false"`
}

func NewTodo(task string) *Todo {
	return &Todo{
		Task:   task,
		Status: Created,
	}
}

// Task Status の独自型の定義
type TaskStatus string

const (
	Created    = TaskStatus("created")
	Processing = TaskStatus("processing")
	Done       = TaskStatus("done")
)

func NewUpdateTodo(id int, task string, status TaskStatus) *Todo {
	return &Todo{
		ID:     id,
		Task:   task,
		Status: status,
	}
}
