package bootstrap

import (
    "database/sql"

    "github.com/gin-gonic/gin"

    httpAdapter "task-api/internal/adapters/inbound/http"
    "task-api/internal/adapters/outbound/persistence"
    "task-api/internal/adapters/outbound/security"
    "task-api/internal/application"
)

type App struct {
    Router *gin.Engine
    DB     *sql.DB
}

func NewApp() (*App, error) {
    cfg := LoadConfig()

    // DB
    db, err := NewDB(cfg.DBUrl)
    if err != nil {
        return nil, err
    }

    // Repositories
    userRepo := &persistence.UserRepo{DB: db}
    taskRepo := &persistence.TaskRepo{DB: db}

    // Services
    jwtSvc := &security.JWTService{
        Secret: cfg.JWTSecret,
    }

    // UseCases
    authUC := &application.AuthUseCase{
        UserRepo: userRepo,
        TokenSvc: jwtSvc,
    }

    userUC := &application.UserUseCase{
        UserRepo: userRepo,
    }

    taskUC := &application.TaskUseCase{
        TaskRepo: taskRepo,
        UserRepo: userRepo,
    }

    // Handlers
    handler := &httpAdapter.Handler{
        AuthUC: authUC,
        UserUC: userUC,
        TaskUC: taskUC,
    }

    // Router
    r := gin.Default()

    // Public
    r.POST("/login", handler.Login)

    // Protected
    authMiddleware := httpAdapter.AuthMiddleware(jwtSvc)

    api := r.Group("/api", authMiddleware)
    {
        api.GET("/tasks", handler.GetTasks)
        api.POST("/tasks", handler.CreateTask)
        api.PUT("/tasks/:id/status", handler.UpdateTaskStatus)

        api.POST("/users", handler.CreateUser)
    }

    return &App{
        Router: r,
        DB:     db,
    }, nil
}