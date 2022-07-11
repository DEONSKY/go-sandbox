package request

//RegisterDTO is used when client post from /register url
type RegisterRequest struct {
	Name     string `json:"name" form:"name" validate:"required,max=32"`
	Email    string `json:"email" form:"email" validate:"required,email" `
	Password string `json:"password" form:"password" validate:"required,min=8,max=32"`
}
