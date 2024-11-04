package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/ThanawatPtd/SAProject/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client // Singleton MongoDB client

func connectMongoDB(ctx context.Context, config *config.Config) *mongo.Client {
	if mongoClient != nil {
		return mongoClient // Reuse the existing client
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(config.MongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	mongoClient = client
	if err != nil {
		panic(err)
	}

	// Check the connection
	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")
	return mongoClient
}

func GetClient() *mongo.Client {
	ctx := context.Background()
	config := config.ProvideConfig()
	client := connectMongoDB(ctx, config)
	return client
} // CloseMongoDB disconnects the client. This should be called when the application shuts down.

func CloseMongoDB() {
	if mongoClient != nil {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			log.Fatal("Error while disconnecting MongoDB:", err)
		}
		fmt.Println("Disconnected from MongoDB")
	}
}
