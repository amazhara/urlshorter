package urlshort

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	storage        *mongo.Client
	db             *mongo.Database
	urlsCollection *mongo.Collection
)

func saveURLs(urls urlsShort) error {
	var data []interface{}

	// Explisitly convert urls to []interface
	for _, url := range urls {
		data = append(data, url)
	}

	if _, err := urlsCollection.InsertMany(context.Background(), data); err != nil {
		return err
	}

	return nil
}

func findURLs() (urlsShort, error) {
	var urls urlsShort

	curr, err := urlsCollection.Find(context.Background(), nil)

	if err != nil {
		return nil, err
	}

	if curr.Decode(&urls); err != nil {
		return nil, err
	}

	return urls, nil
}

// Init can panic when there is no connect to db
func init() {
	var err error

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	storage, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(errors.Wrap(err, "Mongo-driver cannot create Client:"))
	}

	if err := storage.Ping(context.TODO(), nil); err != nil {
		panic(errors.Wrap(err, "Mongo-driver ping func:"))
	}

	// Set db
	db = storage.Database("urlshorter")
	// Set collection
	urlsCollection = db.Collection("urls")
}
