package api

import (
	"basic_management/learn-go-management/controller"
	"basic_management/learn-go-management/store"
)

type ApiRoutes struct {
	Server controller.ServerOperations
}

func (api *ApiRoutes) StartApp(server controller.Server) {

	api.Server = &server
	api.Server.NewServer(store.Connect{})

}
