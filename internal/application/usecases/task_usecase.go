func (uc *TaskUseCase) CreateTask(admin *User, task *Task) error {
    if admin.Role != RoleAdmin {
        return errors.New("unauthorized")
    }

    user, err := uc.userRepo.GetByID(task.AssignedTo)
    if err != nil {
        return err
    }

    if user.Role != RoleExecutor {
        return errors.New("task must be assigned to executor")
    }

    task.Status = StatusAssigned
    return uc.taskRepo.Create(task)
}

func (uc *TaskUseCase) UpdateStatus(user *User, taskID string, status TaskStatus) error {
    if user.Role != RoleExecutor {
        return errors.New("unauthorized")
    }

    task, err := uc.taskRepo.GetByID(taskID)
    if err != nil {
        return err
    }

    if time.Now().After(task.DueDate) {
        return errors.New("task expired")
    }

    if task.AssignedTo != user.ID {
        return errors.New("not your task")
    }

    task.Status = status
    return uc.taskRepo.Update(task)
}

func (uc *TaskUseCase) GetTasks(role, userID string) ([]domain.Task, error) {
    switch role {
    case "ADMIN", "AUDITOR":
        return uc.TaskRepo.GetAll()
    case "EXECUTOR":
        return uc.TaskRepo.GetByUser(userID)
    default:
        return nil, errors.New("unauthorized")
    }
}