package auth

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type APIError struct {
	Message string `json:"msg"`
	Error   string `json:"error"`
}
