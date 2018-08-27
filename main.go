package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func main() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "ProjectID")
	if err != nil {
		log.Fatal(err)
	}

	type Account struct {
		Name    string
		Balance string
	}
	acc := Account{
		Name:    "Alice",
		Balance: "ababa",
	}
	if _, _, err = client.Collection("chatroom").Doc("roomC").Collection("message").Add(ctx, acc); err != nil {
		log.Fatal(err)
	}
}
