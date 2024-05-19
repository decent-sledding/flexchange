type UserService interface {
    New(user NewUserRequest) error
}

// User repository is what lets our service do db operations without knowing anything about the implementation
type UserRepository interface {
    CreateUser(NewUserRequest) error
}