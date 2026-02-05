package mapper

import (
	"golang_crud/src/dto"
	models "golang_crud/src/model"
)

func ToUserDto(user models.User) dto.UserDto {
	return dto.UserDto{
		ID:    user.ID.Hex(),
		Name:  user.Name,
		Email: user.Email,
	}
}
