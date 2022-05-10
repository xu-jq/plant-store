package messages

type UserRegisterReq struct {
	Username string `form:"username"`
	Email    string `form:"email"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

type UserLoginReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type UserUpdate struct {
	Email string `form:"email"`
	Phone string `form:"phone"`
}
