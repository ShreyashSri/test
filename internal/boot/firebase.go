package boot

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var App *firebase.App

func InitFirebase() {
	ctx := context.Background()

	serviceAccountKeyPath := GetEnv("FIREBASE_SERVICE_ACCOUNT_KEY", "serviceAccountKey.json")

	opt := option.WithCredentialsFile(serviceAccountKeyPath)

	var err error
	App, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		// Fallback to default credentials
		log.Printf("Failed to init with service account file: %v. Trying default credentials...", err)
		App, err = firebase.NewApp(ctx, nil)
		if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
		}
	}

	log.Println("Firebase initialized successfully")
}
