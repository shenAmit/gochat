package config

import (
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() {
	err := mgm.SetDefaultConfig(nil, "gochat", options.Client().ApplyURI("mongodb://localhost:27017/gochat"))
	if err != nil {
		log.Fatalf("❌ Failed to connect to MongoDB: %v", err)
	}
	log.Println("✅ Connected to MongoDB using mgm")
}
