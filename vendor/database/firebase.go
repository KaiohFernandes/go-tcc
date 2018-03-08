package database

import (
	"log"
	_ "reflect"

	firebase "firebase.google.com/go"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

type Firebase struct{
	ctx context.Context
	Database
}

func FirebaseInit() *Firebase{
	fb := &Firebase{}

	fb.config()
	fb.ctx = context.Background()
	
	return fb;
}

func (fb *Firebase) Get(client *firestore.CollectionRef, collection string) [] map[string]interface{} {
	
	iter := client.Documents(fb.ctx)

	var items [] map[string]interface{}
	
	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		item := doc.Data();
		item["id"] = doc.Ref.ID
		
		items = append(items, item)
	}

	if items == nil {
		items = append(items, map[string]interface{}{"message": "users not found"})
	}

	return items
}

func (fb *Firebase) GetById(client *firestore.Client, collection string, id string) map[string] interface{} {

	item, err := client.Collection(collection).Doc(id).Get(fb.ctx)

	if err != nil {
		return map[string]interface{}{"message": collection + " data not found"}
	}

	return item.Data()
}

func (fb *Firebase) GetWhere(client *firestore.Client, collection string, query ...string) [] map[string]interface{} {

	iter := client.Collection(collection).Where(query[0], query[1], query[2]).Documents(fb.ctx)

	var items [] map[string]interface{}

	for {
		doc, err := iter.Next()
		
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatal(collection, " data not found")
		}

		item := doc.Data();
		item["id"] = doc.Ref.ID

		items = append(items, doc.Data())
	}

	if items == nil {
		items = append(items, map[string]interface{}{"message": "users not found"})
	}

	return items
}

func (fb *Firebase) Create(client *firestore.Client, collection string, data map[string]interface{}) map[string]string{

	_, _, err := client.Collection(collection).Add(fb.ctx, data)

	if err != nil {
		panic(err)
	}

	res := map[string] string{
		"message" : collection + " created successfully",
	}

	return res
}

func (fb *Firebase) Update(client *firestore.Client, collection string, id string, data map[string]interface{}) map[string]string{

	_, err := client.Collection(collection).Doc(id).Set(fb.ctx, data, firestore.MergeAll)

	if err != nil {
		panic(err)
	}

	return map[string] string {"message": collection + " updated successfully"}
}

func (fb *Firebase) Delete(client *firestore.Client, collection string, id string) map[string]string{

	_, err := client.Collection(collection).Doc(id).Delete(fb.ctx)

	if err != nil {
		panic( err )
	}

	return map[string] string {"message": collection + " deleted successfully"}
}

func (fb *Firebase) OpenConnection() *firestore.Client{
	
	serviceAccount := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(fb.ctx, nil, serviceAccount)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(fb.ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return client
} 

func (fb *Firebase) CloseConnection(client *firestore.Client) {
	client.Close()
}