package tokens

type Token struct {
	UserID       string `json:"-"`
	RefreshToken string `json:"refresh_token"`
}
