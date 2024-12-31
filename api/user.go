package api

import (
	"basic_management/learn-go-management/model"
	"basic_management/learn-go-management/util"

	"github.com/gin-gonic/gin"
)

func (api ApiRoutes) UserRoutes(routes *gin.Engine) {
	group := routes.Group("user")
	{
		group.POST("/create", api.CreateUser)
	}

}

func (api ApiRoutes) CreateUser(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreateUser, "CREATING NEW USER", nil)

	api.Server.CreateUser(ctx)
}
