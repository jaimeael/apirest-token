func (uc *UserUseCase) CreateUser(admin *User, newUser *User) error {
    if admin.Role != RoleAdmin {
        return errors.New("unauthorized")
    }

    if newUser.Role == RoleAdmin {
        return errors.New("cannot create admin users")
    }

    newUser.MustChangePassword = true
    newUser.PasswordHash = HashPassword("temp123")

    return uc.userRepo.Create(newUser)
}