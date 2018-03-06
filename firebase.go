package main

import (
	"log"
	"fmt"
  
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"

  )
  
func main(){

	log.Printf("ae")

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

	// _, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
    //     "first": "Ada",
    //     "last":  "Lovelace",
    //     "born":  1815,
	// })

	// if err != nil {
	// 	log.Fatalf("Failed adding alovelace: %v", err)
	// }

	
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

	dsnap, err := client.Collection("users").Doc("akBPrKubfrmB5uIp6b4g").Get(ctx)
	if err != nil {
		return
	}

	m := dsnap.Data()
	fmt.Printf("Document data: %#v\n", m)

	iter = client.Collection("users").Where("first", "==", "Kaio").Documents(ctx)
	for {
			doc, err := iter.Next()
			if err == iterator.Done {
					break
			}
			if err != nil {
					return
			}
			fmt.Println(doc.Data())
	}


	defer client.Close()
}