package controller

import (
	"basic_management/learn-go-management/model"
	"basic_management/learn-go-management/store"
	"basic_management/learn-go-management/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	PostgresDb store.StoreOperations
}

type ServerOperations interface {
	NewServer(pgStore store.Connect)
	CreateUser(ctx *gin.Context)
}

func (s *Server) NewServer(pgStore store.Connect) {

	util.SetLogger()
	util.Logger.Info("Logger setup Done.\n")

	s.PostgresDb = &pgStore

	err := s.PostgresDb.NewStore()
	if err != nil {
		//util.Logger.Errorf("Error while creating Store connect %v\n", err)
		util.Log(model.LogLevelError, model.Controller, model.NewServer, "Error while creating Store connection", err)
	} else {
		util.Logger.Info("Store Connected\n")
		util.Log(model.LogLevelInfo, model.Controller, model.NewServer, "Store Connected", nil)
	}

	fmt.Printf("CONTROLLER SERVER: %v\n", s)
}
