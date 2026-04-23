func AuthMiddleware(tokenService TokenService) gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")

        userID, role, err := tokenService.Validate(token)
        if err != nil {
            c.AbortWithStatus(401)
            return
        }

        c.Set("userID", userID)
        c.Set("role", role)
        c.Next()
    }
}