type UserRepository interface {
    Create(user *User) error
    GetByUsername(username string) (*User, error)
    GetByID(id string) (*User, error)
    Update(user *User) error
    Delete(id string) error
    List() ([]User, error)
}