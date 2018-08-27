package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
)

func main() {
	ctx := context.Background()
	app := initFirebase()
	client, err := app.Database(ctx)
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
	if _, err = client.NewRef("accounts/alice").Push(ctx, acc); err != nil {
		log.Fatal(err)
	}
}
func initFirebase() *firebase.App {
	// gcp上で鍵を作成し、設定する。
	ctx := context.Background()
	config := &firebase.Config{
		DatabaseURL: "https://study-firebase-c7244.firebaseio.com/",
	}
	app, err := firebase.NewApp(ctx, config)
	if err != nil {
		log.Fatal(err)
	}
	return app
}
