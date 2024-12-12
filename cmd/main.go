package main

import (
	"basic_management/learn-go-management/controller"
	"fmt"
)


func main(){
	server := controller.Server{}

	server.NewServer()

	// PRINT SERVER FOR TESTING 

	fmt.Printf("MAIN SERVER: %v\n", server)
}