package controller

import (
	"basic_management/learn-go-management/store"
	"fmt"
)

type Server struct {
	PostgresDb store.StoreOperations
}

type ServerOperations interface {
	NewServer(pgStore store.Connect)
}

func (s *Server) NewServer(pgStore store.Connect) {

	s.PostgresDb = &pgStore

	s.PostgresDb.NewStore()
	fmt.Printf("CONTROLLER SERVER: %v\n", s)
}
