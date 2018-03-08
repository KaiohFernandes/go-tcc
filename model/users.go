package model

import (
	"database"
)

type UsersModel struct{
	Model
}

func (usersModel *UsersModel) Init() *UsersModel {
	usersModel.db = database.FirebaseInit()
	usersModel.client = usersModel.db.OpenConnection()

	return usersModel
}

func (usersModel *UsersModel) GetAll() [] map[string]interface{} {

	defer usersModel.db.CloseConnection(usersModel.client)
	return usersModel.db.Get(usersModel.client.Collection("users"), "users")
} 

func (usersModel *UsersModel) GetOne() map[string]interface{} {

	defer usersModel.db.CloseConnection(usersModel.client)
	return usersModel.db.GetById(usersModel.document, "users")
} 

func (usersModel *UsersModel) Create( body map[string]interface{} ) map[string] string {

	defer usersModel.db.CloseConnection(usersModel.client)
	return usersModel.db.Create(usersModel.client, "users", body)
} 

func (usersModel *UsersModel) Update( id string, body map[string]interface{} ) map[string] string {

	defer usersModel.db.CloseConnection(usersModel.client)
	return usersModel.db.Update(usersModel.client, "users", id, body)
}

func (usersModel *UsersModel) Delete( id string ) map[string] string {

	defer usersModel.db.CloseConnection(usersModel.client)
	return usersModel.db.Delete(usersModel.client, "users", id)
} 