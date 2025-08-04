package model

import "github.com/google/uuid"

type VerifyUserRequest struct {
	Token string `validate:"required,max=500"`
}

type UserSyncResponse struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
}
