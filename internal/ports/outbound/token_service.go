type TokenService interface {
    Generate(userID string, role string) (string, error)
    Validate(token string) (userID string, role string, err error)
}