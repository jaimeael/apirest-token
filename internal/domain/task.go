type TaskStatus string

const (
    StatusAssigned  TaskStatus = "ASSIGNED"
    StatusCompleted TaskStatus = "COMPLETED"
    StatusExpired   TaskStatus = "EXPIRED"
)

type Task struct {
    ID          string
    Title       string
    Description string
    DueDate     time.Time
    Status      TaskStatus
    AssignedTo  string
}