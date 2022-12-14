package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDB string) (db *mongo.Database, err error) {
	var mongoDBUlr string
	var isAuth bool
	if username == "" && password == "" {
		mongoDBUlr = fmt.Sprintf("mongodb://%s:%s", host, port)
	} else {
		isAuth = true
		mongoDBUlr = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	}

	clientOptions := options.Client().ApplyURI(mongoDBUlr)
	if isAuth {
		if authDB == "" {
			authDB = database
		}
		clientOptions.SetAuth(options.Credential{
			AuthSource: authDB,
			Username:   username,
			Password:   password,
		})
	}
	//Connect
	client, _ := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	//Ping
	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping to MongoDB: %v", err)
	}
	return client.Database(database), nil
}
