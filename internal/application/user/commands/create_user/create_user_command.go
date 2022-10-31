package createuser

type CreateUserCommand struct {
	Username    string `json:"username"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phonenumber" binding:"required"`
}
