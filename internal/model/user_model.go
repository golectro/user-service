package model

type VerifyUserRequest struct {
	Token string `validate:"required,max=500"`
}

type UserSyncResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	CreatedAt int64  `json:"created_at"`
}
