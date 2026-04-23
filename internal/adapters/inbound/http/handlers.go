package http

func (h *Handler) Login(c *gin.Context) {
    var req LoginRequest
    c.BindJSON(&req)

    token, err := h.authUseCase.Login(req.Username, req.Password)
    if err != nil {
        c.JSON(401, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"token": token})
}

func (h *Handler) GetTasks(c *gin.Context) {
    role := c.GetString("role")
    userID := c.GetString("userID")

    tasks, err := h.TaskUC.GetTasks(role, userID)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, tasks)
}