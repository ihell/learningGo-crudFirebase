package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var Client *firestore.Client

func InitFirestore() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Gagal memuat .env: %v", err)
	}

	credJSON := os.Getenv("FIREBASE_CREDENTIAL")
	projectID := os.Getenv("FIREBASE_PROJECT_ID")

	// Convert string to io.Reader
	r := strings.NewReader(credJSON)

	ctx := context.Background()
	opt := option.WithCredentialsJSON([]byte(credJSON))
	client, err := firestore.NewClient(ctx, projectID, opt)
	if err != nil {
		log.Fatalf("Gagal membuat Firestore client: %v", err)
	}

	Client = client
}
