package database

import (
	"log"

	firebase "firebase.google.com/go"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

type Firebase struct{
	DatabaseConfig
}

func (fb *Firebase) Initialize() *firestore.Client{
	fb.config()
	
	return fb.OpenConnection()
	//data := fb.Get(client, "users")
	//dataId := fb.GetById(client, "users", "akBPrKubfrmB5uIp6b4g")
	//dataWhere := fb.GetWhere(client, "users", "first", "==", "Ada")
	//create := fb.Create(client, "users", data)
	//delete := fb.Delete(client, "users", "akBPrKubfrmB5uIp6b4g")
	//update := fb.Update(client, "users", "G7aS2gAvJE1VHhVQKpaM", data)
}

func (fb *Firebase) Get(client *firestore.Client, collection string) interface{}{

	ctx := context.Background()
	
	iter := client.Collection(collection).Documents(ctx)

	var items []interface{}
	
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		
		items = append(items, doc.Data())
	}

	return items

}

func (fb *Firebase) GetById(client *firestore.Client, collection string, id string) interface{} {

	ctx := context.Background()

	item, err := client.Collection(collection).Doc(id).Get(ctx)
	if err != nil {
		return collection + " data not found"
	}

	return item.Data()
}

func (fb *Firebase) GetWhere(client *firestore.Client, collection string, query ...string) interface{} {

	ctx := context.Background()
	iter := client.Collection(collection).Where(query[0], query[1], query[2]).Documents(ctx)
	var items []interface{}

	for {
		doc, err := iter.Next()
		
		if err == iterator.Done {
			break
		}

		if err != nil {
			return collection + " data not found"
		}

		items = append(items, doc.Data())
	}

	return items
}

func (fb *Firebase) Create(client *firestore.Client, collection string, data map[string]interface{}) string{

	ctx := context.Background()

	_, _, err := client.Collection(collection).Add(ctx, data)

	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	return collection + " created successfully"

}

func (fb *Firebase) Update(client *firestore.Client, collection string, id string, data map[string]interface{}) string{

	ctx := context.Background()

	_, err := client.Collection(collection).Doc(id).Set(ctx, data, firestore.MergeAll)

	if err != nil {
		log.Fatalln(err)
	}

	return collection + " updated successfully"

}

func (fb *Firebase) Delete(client *firestore.Client, collection string, id string) string{

	ctx := context.Background()

	_, err := client.Collection(collection).Doc(id).Delete(ctx)

	if err != nil {
		return "Erro"
	}

	return collection + " deleted successfully"

}

func (fb *Firebase) OpenConnection() *firestore.Client{
	
	ctx := context.Background()
	
	serviceAccount := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, serviceAccount)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return client
} 

func (fb *Firebase) CloseConnection(client *firestore.Client) {
	client.Close()
}