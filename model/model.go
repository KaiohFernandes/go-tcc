package model

import (
	"database"
	"cloud.google.com/go/firestore"
)

type modelInterface interface {
	Init() *Model
	GetAll() [] map[string]interface{}
	GetOne() map[string]interface{}
	Create( body map[string]interface{} ) map[string] string
	Update( body map[string]interface{} ) map[string] string
	Delete() map[string] string
}

type Model struct {
	db *database.Firebase
	client *firestore.Client
	collection *firestore.CollectionRef
	document *firestore.DocumentRef
}

func (model *Model) SetCollection(query ...string) *Model {

	if query == nil {
		panic("Error, not have query values")
	}

	for index, value := range query {
		
		if index % 2 == 0 {
			model.collection = model.client.Collection( value )
		} else {
			model.document = model.collection.Doc( value )
		}
		
	}

	return model

}