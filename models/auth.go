package models

const AUTHED_NAME = "authed"
const PASSWORD_REGEX = "^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$"

type SignInDTO struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type SignUpDTO struct {
	Username       string `form:"username" binding:"required"`
	Password       string `form:"password" binding:"required"`
	Email          string `form:"email" binding:"required"`
	PasswordRepeat string `form:"repeat-password" binding:"required"`
}
