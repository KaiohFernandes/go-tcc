package model

import (
	"database"
	"cloud.google.com/go/firestore"
)

type ListsModel struct{
	db *database.Firebase
	client *firestore.Client
}

func ListsInit() *ListsModel{

	listsModel :=  &ListsModel{}
	listsModel.db = database.FirebaseInit()
	listsModel.client = listsModel.db.OpenConnection()

	return listsModel;

}

func (listsModel *ListsModel) GetLists() [] map[string]interface{} {

	defer listsModel.db.CloseConnection(listsModel.client)
	return listsModel.db.Get(listsModel.client.Collection("lists").Doc("G6sFnGbdbPxKo7GMbYr9").Collection("cards"), "lists")
} 

func (listsModel *ListsModel) GetList( documentId string ) map[string]interface{} {

	defer listsModel.db.CloseConnection(listsModel.client)
	return listsModel.db.GetById(listsModel.client, "lists", documentId)
} 

func (listsModel *ListsModel) CreateList( body map[string]interface{} ) map[string] string {

	defer listsModel.db.CloseConnection(listsModel.client)
	return listsModel.db.Create(listsModel.client, "lists", body)
} 

func (listsModel *ListsModel) UpdateList( id string, body map[string]interface{} ) map[string] string {

	defer listsModel.db.CloseConnection(listsModel.client)
	return listsModel.db.Update(listsModel.client, "lists", id, body)
}

func (listsModel *ListsModel) DeleteList( id string ) map[string] string {

	defer listsModel.db.CloseConnection(listsModel.client)
	return listsModel.db.Delete(listsModel.client, "lists", id)
} 