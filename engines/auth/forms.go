package auth

//FmSignIn sign-in form
type FmSignIn struct {
	Email    string `form:"email" binding:"required" validate:"email,max=255"`
	Password string `form:"password" binding:"required"`
}

//FmSignUp sign-up form
type FmSignUp struct {
	Name       string `form:"name" binding:"required" validate:"min=2,max=255"`
	Email      string `form:"email" binding:"required" validate:"email,max=255"`
	Password   string `form:"password" binding:"required" validate:"min=6,max=255"`
	RePassword string `form:"re_password" binding:"required"`
}

//FmEmail email form
type FmEmail struct {
	Email string `form:"email" binding:"required" validate:"email,max=255"`
}

//FmToken email form
type FmToken struct {
	Token string `form:"token" binding:"required"`
}

//FmChangePassword sign-up form
type FmChangePassword struct {
	Token      string `form:"token" binding:"required"`
	Password   string `form:"password" binding:"required" validate:"min=6,max=255"`
	RePassword string `form:"re_password" binding:"required"`
}
