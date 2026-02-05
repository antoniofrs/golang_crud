package dependencies

import (
	"golang_crud/src/config/database"
	"golang_crud/src/repository"
	"golang_crud/src/service"
	"os"
)


type AppDependencies struct {
	UserService service.UserService
}

func Init() *AppDependencies {

	db := database.MongoConnect(
		os.Getenv("MONGO_URI"),
		os.Getenv("MONGO_DB_NAME"),
	)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	return &AppDependencies{
		UserService: userService,
	}

}