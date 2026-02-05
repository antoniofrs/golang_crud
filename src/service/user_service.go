package service

import (
	"context"
	"golang_crud/src/dto"
	mapper "golang_crud/src/factory"
	models "golang_crud/src/model"
	"golang_crud/src/plugin/logger"
	"golang_crud/src/repository"

	"go.uber.org/zap"
)

type userService struct {
	repo repository.UserRepository
}

type UserService interface {
	Create(ctx context.Context, input dto.InsertUserDto) (dto.UserDto, error)
	GetByID(ctx context.Context, id string) (dto.UserDto, error)
	GetAll(ctx context.Context) ([]dto.UserDto, error)
	Delete(ctx context.Context, id string) error
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// Create implements [UserService].
func (s *userService) Create(ctx context.Context, input dto.InsertUserDto) (dto.UserDto, error) {
	logger.Log.Info("service: create user",
		zap.String("email", input.Email),
	)

	user := models.User{
		Name:  input.Name,
		Email: input.Email,
	}

	createdUser, err := s.repo.Create(ctx, &user)
	if err != nil {
		return dto.UserDto{}, err
	}

	return mapper.ToUserDto(*createdUser), nil
}

// Delete implements [UserService].
func (s *userService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

// GetAll implements [UserService].
func (s *userService) GetAll(ctx context.Context) ([]dto.UserDto, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]dto.UserDto, 0, len(users))
	for _, user := range users {
		result = append(result, mapper.ToUserDto(user))
	}

	return result, nil
}


// GetByID implements [UserService].
func (s *userService) GetByID(ctx context.Context, id string) (dto.UserDto, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return dto.UserDto{}, err
	}

	return mapper.ToUserDto(*user), nil
}
