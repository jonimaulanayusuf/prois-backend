package requests

type RegisterRequest struct {
	Username             string `json:"username"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
