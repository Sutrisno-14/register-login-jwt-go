package transport

//UserUpdate -> digunakan ketika user update
type UserUpdate struct {
	ID       uint64 `json:"id_user" form:"id_user"`
	Email    string `json:"email" form:"email" binding:"required,email" `
	Password string `json:"password" form:"password" binding:"required"`
}