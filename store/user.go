package store

import (
	"basic_management/learn-go-management/model"
	"basic_management/learn-go-management/util"
)



func (store Connect)CreateUser(user *model.User) error{
	 util.Log(model.LogLevelInfo, model.StorePackage, model.CreateUser, "Creating New User.", nil)
	 respone := store.DB.Create(user)
	 if respone.Error != nil{
		util.Log(model.LogLevelError, model.StorePackage, model.CreateUser, "ERROR WHILE CREATING USER.", respone.Error)
		return respone.Error
	 }
	 util.Log(model.LogLevelInfo, model.StorePackage, model.CreateUser, "Created New User.", user)
	 return nil

}
