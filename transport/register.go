package transport

//Register i-> di gunakan ketika user melakukan register
type Register struct {
	Email    string `json:"email" form:"email" binding:"required,email" `
	Password string `json:"password" form:"password" binding:"required"`
}