package requests

type RegisterRequest struct {
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	FullName    string `json:"full_name" validate:"required"`
	AccountType string `json:"account_type" validate:"required"`
}
