package database

import (
	"log"
	"fmt"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

type Firebase struct{
	Database *Database
}

func (fb *Firebase) Connect(dbConfig *database.Database){

	ctx := context.Background()

	// Use a service account
	sa := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}
		
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		fmt.Println(doc.Data())
	}

}