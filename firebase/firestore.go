package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"

	"google.golang.org/api/option"
)

var Client *firestore.Client

func InitFirestore() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("serviceAccountKey.json")
	client, err := firestore.NewClient(ctx, "kalo erorro lupa ganti-id", sa)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	Client = client
}
