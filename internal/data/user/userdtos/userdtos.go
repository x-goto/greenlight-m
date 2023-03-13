package userdtos

type CreateUserDTO struct {
	ID       int    `json:"-"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// this DTO
type GetUserDTO struct {
	ID       int    `json:"id"`
	Role     string `json:"-"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UpdateUserDTO struct {
	ID       int    `json:"-"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
