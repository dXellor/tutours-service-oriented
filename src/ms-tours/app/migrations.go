package app

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

func InsertInfo(client *mongo.Client) {
	// Create a context with a timeout of 5 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Access the "tourService" database and "tours" collection.
	toursDatabase := client.Database("tours")
	// Retrieve the list of collections in the database.
	collections, err := toursDatabase.ListCollectionNames(ctx, bson.D{})
	if err != nil {
		log.Println("Error listing collections:", err)
	}

	// Drop each collection.
	for _, collectionName := range collections {
		collection := toursDatabase.Collection(collectionName)
		if err := collection.Drop(ctx); err != nil {
			log.Printf("Error dropping collection %s: %v\n", collectionName, err)
		} else {
			log.Printf("Collection %s dropped successfully\n", collectionName)
		}
	}

	toursCollection := toursDatabase.Collection("tours")
	keypointsCollection := toursDatabase.Collection("keypoints")
	counter := toursDatabase.Collection("counter")

	// Define the documents to insert.
	tours := []interface{}{
		map[string]interface{}{
			"_id":                1,
			"user_id":            16,
			"name":               "Zlatibor Nature Escape",
			"description":        "Discover the natural beauty of Zlatibor.",
			"price":              1500,
			"duration":           37,
			"distance":           7,
			"difficulty":         2,
			"transport_type":     3,
			"status":             1,
			"status_update_time": time.Date(2024, time.February, 16, 0, 0, 0, 0, time.UTC),
			"tags":               []string{"nature", "escape", "Zlatibor"},
		},
		map[string]interface{}{
			"_id":                2,
			"user_id":            16,
			"name":               "Zlatibor Nature Escape2",
			"description":        "Natural beauty of Zlatibor.",
			"price":              1400,
			"duration":           33,
			"distance":           7,
			"difficulty":         2,
			"transport_type":     3,
			"status":             0,
			"status_update_time": time.Date(2024, time.February, 16, 0, 0, 0, 0, time.UTC),
			"tags":               []string{"nature", "escape", "Zlatibor"},
		},
	}

	// Insert documents into MongoDB collection.
	_, err = toursCollection.InsertMany(ctx, tours)
	if err != nil {
		log.Fatal(err)
	}

	keypoints := []interface{}{
		map[string]interface{}{
			"_id":         1,
			"tour_id":     1,
			"name":        "Danube Park",
			"latitude":    45.25329954841971,
			"longitude":   19.829717564246433,
			"description": "An urban park in the downtown of Novi Sad",
			"position":    1,
			"image":       "image danube park",
			"secret":      "Secret of Danube Park",
		},
		map[string]interface{}{
			"_id":         2,
			"tour_id":     2,
			"name":        "Mary Church",
			"latitude":    45.2532995484197,
			"longitude":   19.8297175642465,
			"description": "The largest church in Novi Sad",
			"position":    1,
			"image":       "image mary church",
			"secret":      "Secret of Mary Church",
		},
		map[string]interface{}{
			"_id":         3,
			"tour_id":     1,
			"name":        "Mary Church1",
			"latitude":    45.2532995484197,
			"longitude":   19.8297175642465,
			"description": "The largest church in Novi Sad",
			"position":    2,
			"image":       "image mary church",
			"secret":      "Secret of Mary Church",
		},
	}

	// Insert documents into MongoDB collection.
	_, err = keypointsCollection.InsertMany(ctx, keypoints)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Documents inserted successfully.")
	counterDocs := []interface{}{
		map[string]interface{}{
			"_id":   1,
			"value": 3,
			"name":  "tour",
		},
		map[string]interface{}{
			"_id":   2,
			"value": 1,
			"name":  "review",
		},
		map[string]interface{}{
			"_id":   3,
			"value": 4,
			"name":  "keypoint",
		},
		map[string]interface{}{
			"_id":   4,
			"value": 1,
			"name":  "position",
		},
		// Add more counter documents as needed
	}

	_, err = counter.InsertMany(ctx, counterDocs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Document inserted into 'counter' collection successfully.")
}
