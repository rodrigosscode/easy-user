package delivery

import (
	"github.com/gin-gonic/gin"
)

func (e *ginEngine) deleteUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		query.Add(QueryParamUserId, ctx.Param(QueryParamUserId))
		ctx.Request.URL.RawQuery = query.Encode()
		e.deleteUserByIdCtrl.Execute(ctx.Writer, ctx.Request)
	}
}

func (e *ginEngine) findUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		query.Add(QueryParamUserId, ctx.Param(QueryParamUserId))
		ctx.Request.URL.RawQuery = query.Encode()
		e.findUserByIdCtrl.Execute(ctx.Writer, ctx.Request)
	}
}

func (e *ginEngine) findUsersPaging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		e.findUsersPagingCtrl.Execute(ctx.Writer, ctx.Request)
	}
}

func (e *ginEngine) saveUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		e.saveUserCtrl.Execute(ctx.Writer, ctx.Request)
	}
}

func (e *ginEngine) updateUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		e.updateUserByIdCtrl.Execute(ctx.Writer, ctx.Request)
	}
}
