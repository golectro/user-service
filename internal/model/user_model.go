package model

import "github.com/google/uuid"

type VerifyUserRequest struct {
	Token string `validate:"required,max=500"`
}

type UserResponse struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	AvatarObject string    `json:"avatar_object,omitempty"`
}

type UserSyncResponse struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	AvatarObject string    `json:"avatar_object,omitempty"`
}

type UploadAvatarResponse struct {
	AvatarObject string `json:"avatar_object"`
}
