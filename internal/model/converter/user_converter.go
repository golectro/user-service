package converter

import (
	"golectro-user/internal/entity"
	"golectro-user/internal/model"
)

func UserToResponse(user *entity.User) *model.UserSyncResponse {
	return &model.UserSyncResponse{
		Email:    user.Email,
		Username: user.Username,
	}
}
