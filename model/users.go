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

func (usersModel *UsersModel) GetUsers() [] map[string] interface{} {

	defer usersModel.db.CloseConnection(usersModel.client)
	return usersModel.db.Get(usersModel.client, "users")
} 