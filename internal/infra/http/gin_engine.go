package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigosscode/easy-user/internal/delivery/controller"
)

type (
	GinRouter interface {
		SetAppHandlers()
		GetRouter() *gin.Engine
	}

	ginEngine struct {
		router              *gin.Engine
		deleteUserByIdCtrl  *controller.DeleteUserByIdController
		findUserByIdCtrl    *controller.FindUserByIdController
		findUsersPagingCtrl *controller.FindUsersPagingController
		saveUserCtrl        *controller.SaveUserController
		updateUserByIdCtrl  *controller.UpdateUserByIdController
	}
)

func NewGinEngine(
	router *gin.Engine,
	deleteUserByIdCtrl *controller.DeleteUserByIdController,
	findUserByIdCtrl *controller.FindUserByIdController,
	findUsersPagingCtrl *controller.FindUsersPagingController,
	saveUserCtrl *controller.SaveUserController,
	updateUserByIdCtrl *controller.UpdateUserByIdController,
) *ginEngine {
	return &ginEngine{
		router:              router,
		deleteUserByIdCtrl:  deleteUserByIdCtrl,
		findUserByIdCtrl:    findUserByIdCtrl,
		findUsersPagingCtrl: findUsersPagingCtrl,
		saveUserCtrl:        saveUserCtrl,
		updateUserByIdCtrl:  updateUserByIdCtrl,
	}
}

func (e *ginEngine) SetAppHandlers() {
	e.router.DELETE(RouteDeleteUserById, e.deleteUserById())
	e.router.GET(RouteFindUserById, e.findUserById())
	e.router.GET(RouteFindUsersPaging, e.findUsersPaging())
	e.router.POST(RouteSaveUser, e.saveUser())
	e.router.PUT(RouteUpdateUserById, e.updateUserById())
}
