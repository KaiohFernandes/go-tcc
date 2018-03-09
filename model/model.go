package model

import (
	"database"
	"cloud.google.com/go/firestore"
)

// padrão terminar com er
type Modeler interface {
	Init() Modeler
	SetCollection(query ...string) Modeler
	GetAll() [] map[string]interface{}
	GetOne() map[string]interface{}
	Create( body map[string]interface{} ) map[string] string
	Update( id string, body map[string]interface{} ) map[string] string
	Delete( id string ) map[string] string
}

type Model struct {
	db *database.Firebase
	client *firestore.Client
	collection *firestore.CollectionRef
	document *firestore.DocumentRef
}

func (m *Model) Init() Modeler {
	m.db = database.FirebaseInit()
	m.client = m.db.OpenConnection()

	return m
}

func (m *Model) GetAll() [] map[string]interface{} {

	defer m.db.CloseConnection(m.client)
	return m.db.Get(m.collection, "lists")
} 

func (m *Model) GetOne() map[string]interface{} {

	defer m.db.CloseConnection(m.client)
	return m.db.GetById(m.document, "lists")
} 

func (m *Model) Create( body map[string]interface{} ) map[string] string {

	defer m.db.CloseConnection(m.client)
	return m.db.Create(m.client, "lists", body)
} 

func (m *Model) Update( id string, body map[string]interface{} ) map[string] string {

	defer m.db.CloseConnection(m.client)
	return m.db.Update(m.client, "lists", id, body)
}

func (m *Model) Delete( id string ) map[string] string {

	defer m.db.CloseConnection(m.client)
	return m.db.Delete(m.client, "lists", id)
} 

func (m *Model) SetCollection(query ...string) Modeler {

	if query == nil {
		panic("Error, not have query values")
	}

	for index, value := range query {
		
		if index % 2 == 0 {

			if index == 0 {
				m.collection = m.client.Collection( value )
			} else {
				m.collection = m.document.Collection( value )
			}

		} else {
			m.document = m.collection.Doc( value )
		}
		
	}

	return m

}