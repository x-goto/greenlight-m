package users

type Role string

const (
	RoleUser      Role = "USER"
	RoleModerator Role = "MODERATOR"
	RoleAdmin     Role = "ADMIN"
)

type User struct {
	ID       int    `json:"id"`
	Role     Role   `json:"-"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
