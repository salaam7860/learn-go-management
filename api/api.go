package api

import (
	"basic_management/learn-go-management/controller"
	"basic_management/learn-go-management/store"

	"github.com/gin-gonic/gin"
)

type ApiRoutes struct {
	Server controller.ServerOperations
}

func (api *ApiRoutes) StartApp(routes *gin.Engine, server controller.Server) {

	api.Server = &server
	api.Server.NewServer(store.Connect{})


	api.UserRoutes(routes)
}
