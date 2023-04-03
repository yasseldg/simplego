package sMongo

import (
	"context"
	"fmt"
	"time"

	"github.com/yasseldg/simplego/sLog"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getClientOpt(conn *ConnectionParams) *options.ClientOptions {

	Uri, Credentials := getConnectionUri(conn)

	switch conn.Environment {
	case "prod":
		return options.Client().ApplyURI(Uri)

	default: // dev
		return options.Client().ApplyURI(Uri).SetAuth(Credentials)
	}
}

func getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), (5 * time.Second))
}

func getDatabase(conn *ConnectionParams, database string) *mongo.Database {

	ctx, cancel := getContext()
	defer cancel()

	clientOpt := getClientOpt(conn)

	client, err := mongo.Connect(ctx, clientOpt)
	if err != nil {
		sLog.Fatal("getDatabase: mongo.Connect(ctx, ClientOpt): %s", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		sLog.Fatal(err.Error())
	}

	return client.Database(database)
}

func getMgmColl(conn *ConnectionParams, database, collection string) *mgm.Collection {

	fmt.Println()
	time.Sleep(1 * time.Second)
	sLog.Info("Host: %s  ..  Env: %s  ..  AuthDB: %s  ..  User: %s  ..  Database: %s  ..  Collection: %s", conn.Host, conn.Environment, conn.AuthDatabase, conn.Username, database, collection)

	return &mgm.Collection{Collection: getDatabase(conn, database).Collection(collection)}
}

// CreateIndex - creates an index for a specific field in a collection
func cicleCreateIndex(coll *mgm.Collection, fields interface{}, unique bool) error {

	// 1. Lets define the keys for the index we want to create
	mod := mongo.IndexModel{
		Keys:    fields, // index in ascending order or -1 for descending order
		Options: options.Index().SetUnique(unique),
	}

	ctx, cancel := context.WithTimeout(context.Background(), (35 * time.Second))
	defer cancel()

	// 4. Create a single index
	count := 0
	for {
		_, err := coll.Indexes().CreateOne(ctx, mod)
		if err == nil {
			return nil
		} else {
			sLog.Error("coll.Indexes().CreateOne(ctx, mod): %s", err)
		}

		time.Sleep(time.Second)
		count++

		if count > 15 {
			sLog.Error("Mongo DB: %s.%s CreateIndex failure ", coll.Database().Name(), coll.Name())
			sLog.Fatal(err.Error())
		}
	}
}
