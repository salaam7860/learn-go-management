package controller

import (
	"basic_management/learn-go-management/model"
	"basic_management/learn-go-management/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)




func (server Server)CreateUser(ctx *gin.Context){

	// PASING JSON OBJ INTO STRUCT 
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.CreateUser, "CREATING NEW USER", nil)
	user := model.User{}
	err := ctx.Bind(&user)
	if err != nil{
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateUser, "ERROR WHILE JSON BINDING.", err)
		ctx.JSON(http.StatusInternalServerError,err)
		return 
	}

	
	// CREATING USER IN THE DATABASE 
	user.ID = uuid.New()
	user.CreatedAt = time.Now()

	err = server.PostgresDb.CreateUser(&user)
	if err != nil{
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateUser, "ERROR WHILE JSON BINDING.", err)
		ctx.JSON(http.StatusInternalServerError,err)
		return 
	}
	ctx.JSON(http.StatusCreated, user)

}