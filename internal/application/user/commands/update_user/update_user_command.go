package updateuser

type UpdateUserCommand struct {
	Username    string `json:"username"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phonenumber" binding:"required"`
}
