package model

import (
	"database"
	"cloud.google.com/go/firestore"
)

type UsersModel struct{
	db *database.Firebase
	client *firestore.Client
}

func Init() *UsersModel{

	usersModel :=  &UsersModel{}
	usersModel.db = database.FirebaseInit()
	usersModel.client = usersModel.db.OpenConnection()

	return usersModel;

}

func (usersModel *UsersModel) GetUsers() [] map[string]interface{} {

	defer usersModel.db.CloseConnection(usersModel.client)
	return usersModel.db.Get(usersModel.client, "users")
} 

func (usersModel *UsersModel) GetUser( documentId string ) map[string]interface{} {

	defer usersModel.db.CloseConnection(usersModel.client)
	return usersModel.db.GetById(usersModel.client, "users", documentId)
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