package users

type UserInviteRequest struct {
	Email string `json:"email" binding:"required,min=10"`
}

type UserActivateRequest struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

type UserAuthenticateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
