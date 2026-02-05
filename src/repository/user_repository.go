package repository

import (
	"context"
	"errors"
	models "golang_crud/src/model"
	"golang_crud/src/plugin/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type userRepository struct {
	collection *mongo.Collection
}

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	FindByID(ctx context.Context, id string) (*models.User, error)
	FindAll(ctx context.Context) ([]models.User, error)
	Update(ctx context.Context, id string, user *models.User) error
	Delete(ctx context.Context, id string) error
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}


func (r *userRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		logger.Log.Error("insert user failed", zap.Error(err))
		return nil, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}


func (u *userRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

func (u *userRepository) FindAll(ctx context.Context) ([]models.User, error) {
	panic("unimplemented")
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	var user models.User
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}



func (u *userRepository) Update(ctx context.Context, id string, user *models.User) error {
	panic("unimplemented")
}
