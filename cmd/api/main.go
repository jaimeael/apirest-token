package main

import (
    "log"
    "task-api/internal/bootstrap"                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                
)

func main() {
    app, err := bootstrap.NewApp()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Server running on :8080")
    app.Router.Run(":8080")
}