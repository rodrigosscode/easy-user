package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	usecase "github.com/rodrigosscode/easy-user/internal/application/usecase/user"
	"github.com/rodrigosscode/easy-user/internal/infrastructure/db"
	"github.com/rodrigosscode/easy-user/internal/infrastructure/http"
	"github.com/rodrigosscode/easy-user/internal/infrastructure/mapper"
	"github.com/rodrigosscode/easy-user/internal/infrastructure/repository"
	"github.com/rodrigosscode/easy-user/internal/interface/delivery/controller"
)

func main() {
	ginDefault := gin.Default()

	userDbSource, _ := db.NewDbConnection()
	userMapper := mapper.NewUserMapper()
	userRepository := repository.NewUserRepository(userDbSource, userMapper)

	findUserByIdUseCase := usecase.NewFindByIdUseCase(userRepository)
	findUserByIdCtrl := controller.NewFindUserByIdController(findUserByIdUseCase)

	deleteUserByIdUseCase := usecase.NewDeleteByIdUseCase(userRepository)
	deleteUserByIdCtrl := controller.NewDeleteUserByIdController(deleteUserByIdUseCase)

	saveUserUseCase := usecase.NewSaveUseCase(userRepository)
	saveUserCtrl := controller.NewSaveUserController(saveUserUseCase)

	updateUserUseCase := usecase.NewUpdateUseCase(userRepository)
	updateUserByIdCtrl := controller.NewUpdateUserByIdController(updateUserUseCase)

	findUsersPagingCtrl := controller.FindUsersPagingController{}

	appEngine := http.NewGinEngine(ginDefault, deleteUserByIdCtrl, findUserByIdCtrl, &findUsersPagingCtrl, saveUserCtrl, updateUserByIdCtrl)
	appEngine.SetAppHandlers()

	ginDefault.Run(":8080")

	fmt.Println("Iniciando...")
}
