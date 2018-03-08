package model

import (
	"database"
)

type ListsModel struct{
	Model
}

func (listsModel *ListsModel) Init() *ListsModel {
	listsModel.db = database.FirebaseInit()
	listsModel.client = listsModel.db.OpenConnection()

	return listsModel
}

func (listsModel *ListsModel) GetAll() [] map[string]interface{} {

	defer listsModel.db.CloseConnection(listsModel.client)
	return listsModel.db.Get(listsModel.collection, "lists")
} 

func (listsModel *ListsModel) GetOne() map[string]interface{} {

	defer listsModel.db.CloseConnection(listsModel.client)
	return listsModel.db.GetById(listsModel.document, "lists")
} 

func (listsModel *ListsModel) Create( body map[string]interface{} ) map[string] string {

	defer listsModel.db.CloseConnection(listsModel.client)
	return listsModel.db.Create(listsModel.client, "lists", body)
} 

func (listsModel *ListsModel) Update( id string, body map[string]interface{} ) map[string] string {

	defer listsModel.db.CloseConnection(listsModel.client)
	return listsModel.db.Update(listsModel.client, "lists", id, body)
}

func (listsModel *ListsModel) Delete( id string ) map[string] string {

	defer listsModel.db.CloseConnection(listsModel.client)
	return listsModel.db.Delete(listsModel.client, "lists", id)
} 