package http

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigosscode/easy-user/internal/interface/delivery/controller"
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
		updateUserCtrl      *controller.UpdateUserController
	}
)

func NewGinEngine(
	router *gin.Engine,
	deleteUserByIdCtrl *controller.DeleteUserByIdController,
	findUserByIdCtrl *controller.FindUserByIdController,
	findUsersPagingCtrl *controller.FindUsersPagingController,
	saveUserCtrl *controller.SaveUserController,
	updateUserCtrl *controller.UpdateUserController,
) *ginEngine {
	return &ginEngine{
		router:              router,
		deleteUserByIdCtrl:  deleteUserByIdCtrl,
		findUserByIdCtrl:    findUserByIdCtrl,
		findUsersPagingCtrl: findUsersPagingCtrl,
		saveUserCtrl:        saveUserCtrl,
		updateUserCtrl:      updateUserCtrl,
	}
}

func (e *ginEngine) SetAppHandlers() {
	e.router.DELETE(RouteDeleteUserById, e.deleteUserById())
	e.router.GET(RouteFindUserById, e.findUserById())
	e.router.GET(RouteFindUsersPaging, e.findUsersPaging())
	e.router.POST(RouteSaveUser, e.saveUser())
	e.router.PUT(RouteUpdateUser, e.updateUser())
}

func (e *ginEngine) GetRouter() *gin.Engine {
	return e.router
}
