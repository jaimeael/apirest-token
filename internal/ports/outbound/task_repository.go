type TaskRepository interface {
    Create(task *Task) error
    Update(task *Task) error
    Delete(id string) error
    GetByID(id string) (*Task, error)
    ListByUser(userID string) ([]Task, error)
    ListAll() ([]Task, error)
}