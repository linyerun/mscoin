package dbinit

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func CreateMongoDBClient(username, password, url string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// create connection
	credential := options.Credential{Username: username, Password: password}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url).SetAuth(credential))
	if err != nil {
		return nil, err
	}

	// ping
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
