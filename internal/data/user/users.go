package data

type password string

type Role string

const (
	RoleUser      Role = "USER"
	RoleModerator Role = "MODERATOR"
	RoleAdmin     Role = "ADMIN"
)

type User struct {
	ID       int      `json:"id"`
	Role     Role     `json:"role"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password password `json:"-"`
}
