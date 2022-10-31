package getuser

type GetUserQuery struct {
	Username string `json:"username" binding:"required"`
}
