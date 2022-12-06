package dto

type RegisterDTO struct {
	Email    string `json:"email" form:"email" binding:"required,email" validate:"email"`
	Name     string `json:"name" form:"name" binding:"required,min=6"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}
