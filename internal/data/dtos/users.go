package dtos

type UserRegistrationDTO struct {
	ID       int    `json:"-"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// this DTO
type UserFetchingDTO struct {
	ID       int    `json:"id"`
	Role     string `json:"-"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
