package main

import (
	"context"
	"fmt"
	"log"

	"io"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/google/uuid"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	f, err := Connect()
	if err != nil {
		log.Println(err)
		return
	}
	f.Create(&Bin{
		ID:             uuid.New(),
		BIN:            "1234567",
		CardNetwork:    "Visa",
		CardType:       "Credit",
		CardLevel:      "Classic",
		IssuerWebsite:  "https://www.visa.com",
		IssuerBankName: "Visa",
		CountryName:    "United States",
		CountryCode:    "US",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})
}

type FireDB struct {
	*db.Client
}

type Bin struct {
	ID             uuid.UUID `json:"id"`
	BIN            string    `json:"bin"`
	CardNetwork    string    `json:"card_network"`
	CardType       string    `json:"card_type"`
	CardLevel      string    `json:"card_level"`
	IssuerWebsite  string    `json:"issuer_website"`
	IssuerBankName string    `json:"issuer_bank_name"`
	CountryName    string    `json:"country_name"`
	CountryCode    string    `json:"country_code"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func Connect() (*FireDB, error) {
	opt := option.WithCredentialsFile("./fcm-poc-b3506-firebase-adminsdk-ffxgg-16607542af.json")
	config := &firebase.Config{DatabaseURL: "https://fcm-poc-b3506.firebaseio.com"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Println("error initializing app")
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	client, err := app.Database(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error initializing database: %v", err)
	}
	log.Println("Firebase app initialized")
	return &FireDB{
		Client: client,
	}, nil
}

func (f *FireDB) Create(b *Bin) error {
	if err := f.NewRef("bins/"+b.BIN).Set(context.Background(), b); err != nil {
		return err
	}
	return nil
}

// listenDocument listens to a single document.
func (f *FireDB) listenDocument(ctx context.Context, w io.Writer, projectID, collection string) error {
	// projectID := "project-id"
	// Сontext with timeout stops listening to changes.
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("firestore.NewClient: %w", err)
	}
	defer client.Close()

	it := client.Collection(collection).Doc("SF").Snapshots(ctx)
	for {
		snap, err := it.Next()
		// DeadlineExceeded will be returned when ctx is cancelled.
		if status.Code(err) == codes.DeadlineExceeded {
			return nil
		}
		if err != nil {
			return fmt.Errorf("Snapshots.Next: %w", err)
		}
		if !snap.Exists() {
			fmt.Fprintf(w, "Document no longer exists\n")
			return nil
		}
		fmt.Fprintf(w, "Received document snapshot: %v\n", snap.Data())
	}
}
