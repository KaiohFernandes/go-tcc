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
	ctx context.Context
	Database
}

func FirebaseInit() *Firebase{
	fb := &Firebase{}

	fb.config()
	fb.ctx = context.Background()
	
	return fb;
}

func (fb *Firebase) Get(client *firestore.CollectionRef) [] map[string]interface{} {

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

func (fb *Firebase) GetById(client *firestore.DocumentRef) map[string] interface{} {

	item, err := client.Get(fb.ctx)

	if err != nil {
		return map[string]interface{}{"message": "Data not found"}
	}

	return item.Data()
}

func (fb *Firebase) GetWhere(client *firestore.CollectionRef, query ...string) [] map[string]interface{} {

	iter := client.Where(query[0], query[1], query[2]).Documents(fb.ctx)

	var items [] map[string]interface{}

	for {
		doc, err := iter.Next()
		
		if err == iterator.Done {
			break
		}

		if err != nil {
			panic(err)
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

func (fb *Firebase) Create(client *firestore.CollectionRef, data map[string]interface{}) map[string]string{

	_, _, err := client.Add(fb.ctx, data)

	if err != nil {
		panic(err)
	}

	res := map[string] string{
		"message" : "Created successfully",
	}

	return res
}

func (fb *Firebase) Update(client *firestore.DocumentRef, data map[string]interface{}) map[string]string{

	_, err := client.Set(fb.ctx, data, firestore.MergeAll)

	if err != nil {
		panic(err)
	}

	return map[string] string {"message": "Updated successfully"}
}

func (fb *Firebase) Delete(client *firestore.DocumentRef) map[string]string{

	_, err := client.Delete(fb.ctx)

	if err != nil {
		panic( err )
	}

	return map[string] string {"message": " deleted successfully"}
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