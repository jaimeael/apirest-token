type Role string

const (
    RoleAdmin    Role = "ADMIN"
    RoleExecutor Role = "EXECUTOR"
    RoleAuditor  Role = "AUDITOR"
)

type User struct {
    ID                string
    Username          string
    PasswordHash      string
    Role              Role
    MustChangePassword bool
}