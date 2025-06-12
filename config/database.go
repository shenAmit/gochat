package config

import (
	"log"
	"os"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() {
	mongoURI := os.Getenv("MONGO_URI")
	err := mgm.SetDefaultConfig(nil, "gochat", options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("❌ Failed to connect to MongoDB: %v", err)
	}
	log.Println("✅ Connected to MongoDB using mgm")
}
