package converter

import (
	"golectro-user/internal/entity"
	"golectro-user/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:           user.ID,
		Email:        user.Email,
		Username:     user.Username,
		AvatarObject: user.AvatarObject,
	}
}

func UserSyncToResponse(user *entity.User) *model.UserSyncResponse {
	return &model.UserSyncResponse{
		ID:           user.ID,
		Email:        user.Email,
		Username:     user.Username,
		AvatarObject: user.AvatarObject,
	}
}
