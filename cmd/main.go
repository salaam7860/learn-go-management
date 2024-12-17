package main

import (
	"basic_management/learn-go-management/api"
	"basic_management/learn-go-management/controller"

	"fmt"
)

func main() {

	api := api.ApiRoutes{}

	api.StartApp(controller.Server{})

	fmt.Printf("MAIN SERVER: %v\n", api)
}
