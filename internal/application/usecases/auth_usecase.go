func (uc *AuthUseCase) Login(username, password string) (string, error) {
    user, err := uc.userRepo.GetByUsername(username)
    if err != nil {
        return "", err
    }

    if !CheckPassword(password, user.PasswordHash) {
        return "", errors.New("invalid credentials")
    }

    token, err := uc.tokenService.Generate(user.ID, string(user.Role))
    if err != nil {
        return "", err
    }

    return token, nil
}