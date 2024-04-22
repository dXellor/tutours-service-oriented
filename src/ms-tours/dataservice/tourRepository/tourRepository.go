package TourRepository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"strconv"
	"time"
	"tutours/soa/ms-tours/model"
	"tutours/soa/ms-tours/model/enum"
)

type TourRepository struct {
	cli    *mongo.Client
	logger *log.Logger
}

func New(ctx context.Context, logger *log.Logger) (*TourRepository, error) {
	dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &TourRepository{
		cli:    client,
		logger: logger,
	}, nil
}

func (pr *TourRepository) Disconnect(ctx context.Context) error {
	err := pr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (pr *TourRepository) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := pr.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		pr.logger.Println(err)
	}

	databases, err := pr.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		pr.logger.Println(err)
	}
	fmt.Println(databases)
}

func (pr *TourRepository) getCollection() *mongo.Collection {
	patientDatabase := pr.cli.Database("tours")
	patientsCollection := patientDatabase.Collection("tours")
	return patientsCollection
}

func (tourRepository *TourRepository) GetAll() (model.Tours, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	toursCollection := tourRepository.getCollection()

	var tours model.Tours
	toursCursor, err := toursCollection.Find(ctx, bson.M{})
	if err != nil {
		tourRepository.logger.Println(err)
		return nil, err
	}
	if err = toursCursor.All(ctx, &tours); err != nil {
		tourRepository.logger.Println(err)
		return nil, err
	}
	return tours, nil

}

func (tourRepository *TourRepository) Get(id int) (*model.Tour, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	toursCollection := tourRepository.getCollection()

	var tour model.Tour
	objID, _ := primitive.ObjectIDFromHex(strconv.Itoa(id))
	err := toursCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&tour)
	if err != nil {
		tourRepository.logger.Println(err)
		return nil, err
	}
	return &tour, nil
}

func (tourRepository *TourRepository) Create(tour *model.Tour) (*model.Tour, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	toursCollection := tourRepository.getCollection()

	result, err := toursCollection.InsertOne(ctx, &tour)
	if err != nil {
		tourRepository.logger.Println(err)
		return tour, err
	}
	tourRepository.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil, nil
}

func (tourRepository *TourRepository) Update(tour *model.Tour) (*model.Tour, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	toursCollection := tourRepository.getCollection()

	objID, _ := primitive.ObjectIDFromHex(strconv.Itoa(tour.Id))
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"name":             tour.Name,
		"price":            tour.Price,
		"duration":         tour.Duration,
		"distance":         tour.Distance,
		"difficulty":       tour.Difficulty,
		"transportType":    tour.TransportType,
		"status":           tour.Status,
		"statusUpdateTime": tour.StatusUpdateTime,
		"tags":             tour.Tags,
	}}
	result, err := toursCollection.UpdateOne(ctx, filter, update)
	tourRepository.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	tourRepository.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		tourRepository.logger.Println(err)
		return tour, err
	}
	return nil, nil
}

func (tourRepository *TourRepository) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	toursCollection := tourRepository.getCollection()

	objID, _ := primitive.ObjectIDFromHex(strconv.Itoa(id))
	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := toursCollection.DeleteOne(ctx, filter)
	if err != nil {
		tourRepository.logger.Println(err)
		return err
	}
	tourRepository.logger.Printf("Documents deleted: %v\n", result.DeletedCount)
	return nil
}

func (tourRepository *TourRepository) GetByAuthor(authorId int) ([]model.Tour, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	toursCollection := tourRepository.getCollection()

	filter := bson.M{"authorId": authorId} // Assuming the field name in MongoDB is authorId
	cursor, err := toursCollection.Find(ctx, filter)
	if err != nil {
		tourRepository.logger.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var tours []model.Tour
	if err := cursor.All(ctx, &tours); err != nil {
		tourRepository.logger.Println(err)
		return nil, err
	}

	return tours, nil
}

func (tourRepository *TourRepository) GetPublished() ([]model.Tour, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	toursCollection := tourRepository.getCollection()

	filter := bson.M{"status": enum.PUBLISHED}
	cursor, err := toursCollection.Find(ctx, filter)
	if err != nil {
		tourRepository.logger.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var tours []model.Tour
	if err := cursor.All(ctx, &tours); err != nil {
		tourRepository.logger.Println(err)
		return nil, err
	}

	/*for i, tour := range tours {
		var keypoints []model.Keypoint
		//filter := bson.M{"id": tour.Id}
		//cursor, err := keypointsCollection.Find(ctx, filter)
		if err != nil {
			tourRepository.logger.Println(err)
			return nil, err
		}
		defer cursor.Close(ctx)
		if err := cursor.All(ctx, &keypoints); err != nil {
			tourRepository.logger.Println(err)
			return nil, err
		}
		tours[i].Keypoints = keypoints
	}*/

	return tours, nil
}

func (tourRepository *TourRepository) GetPublishedByAuthor(authorId int) ([]model.Tour, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	toursCollection := tourRepository.getCollection()

	filter := bson.M{"status": enum.PUBLISHED, "authorId": authorId}
	cursor, err := toursCollection.Find(ctx, filter)
	if err != nil {
		tourRepository.logger.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var tours []model.Tour
	if err := cursor.All(ctx, &tours); err != nil {
		tourRepository.logger.Println(err)
		return nil, err
	}

	return tours, nil
}
