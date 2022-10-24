package response

type UserResponse struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type UpdateUserResponse struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Age       int    `json:"age"`
	UpdatedAt string `json:"updated_at"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
