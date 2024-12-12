package controller

import (
	"basic_management/learn-go-management/store"
	"fmt"
)


type Server struct{
	PostgresDb store.Connect
}


func (s *Server) NewServer(){
	
	
	s.PostgresDb.NewStore()
	fmt.Printf("CONTROLLER SERVER: %v\n", s)
}