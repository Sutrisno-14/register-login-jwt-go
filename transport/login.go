package transport

//login -> di gunakan ketika users melakukan login
type Login struct {
	Email    string `json:"email" form:"email" binding:"required,email" `
	Password string `json:"password" form:"password" binding:"required"`
}