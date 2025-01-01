package main

import (
	"basic_management/learn-go-management/api"
	"basic_management/learn-go-management/controller"

	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	api := api.ApiRoutes{}

	controller := controller.Server{}
	routes := gin.Default()
	api.StartApp(routes, controller)

	routes.Run(":8000")

	fmt.Printf("MAIN SERVER: %v\n", api)
}
