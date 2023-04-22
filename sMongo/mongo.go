package sMongo

import (
	"context"
	"fmt"
	"time"

	"github.com/yasseldg/simplego/sConv"
	"github.com/yasseldg/simplego/sEnv"
	"github.com/yasseldg/simplego/sLog"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	Connection *ConnectionParams
	Database   *mongo.Database
	Collection *mgm.Collection
}

// New
func New(env, connection, database, collection, interval string) *Collection {
	var c Collection

	mgm.SetDefaultConfig(getCtx(env))

	c.Connection = getConnection(sEnv.Get(fmt.Sprint("CONN_", env), connection))

	client, err := mgm.NewClient(c.Connection.getClientOpt())
	if err != nil {
		sLog.Error("sMongo.New: mgm.NewClient() for %s.%s: %s", database, collection, err)
		return nil
	}

	c.Database = client.Database((sEnv.Get(fmt.Sprint("DB_", env), database)))
	if c.Database == nil {
		sLog.Error("sMongo.New: client.Database() for %s.%s", database, collection)
		return nil
	}

	coll_name := sEnv.Get(fmt.Sprint("COLL_", env), collection)
	if len(interval) > 0 {
		coll_name = fmt.Sprintf("%s_%s", coll_name, interval)
	}

	c.Collection = mgm.NewCollection(c.Database, coll_name)
	if c.Collection == nil {
		sLog.Error("sMongo.New: mgm.NewCollection() for %s.%s", database, collection)
		return nil
	}

	c.print()

	return &c
}

func getCtx(env string) *mgm.Config {
	return &mgm.Config{
		CtxTimeout: time.Duration(sConv.GetInt(sEnv.Get(fmt.Sprint("CTX_", env), "10"))) * time.Second,
	}
}

func (c *Collection) print() {
	fmt.Println()
	time.Sleep(1 * time.Second)
	sLog.Info("sMongo: Host: %s  ..  Env: %s  ..  AuthDB: %s  ..  User: %s  ..  Database: %s  ..  Collection: %s \n",
		c.Connection.Host, c.Connection.Environment, c.Connection.AuthDatabase, c.Connection.Username, c.Database.Name(), c.Collection.Name())
}

func (c *Collection) prefix() string {
	return fmt.Sprintf("%s.%s.", c.Database.Name(), c.Collection.Name())
}

// Clone, create Connection to a new collection in the same DB
func (c *Collection) Clone(name string) *Collection {
	clone := *c
	clone.Collection = mgm.NewCollection(c.Database, name)
	if clone.Collection == nil {
		sLog.Error("sMongo.Clone: mgm.NewCollection() for %s.%s", c.Database.Name(), name)
		return nil
	}
	return &clone
}

// CreateIndex, create an index for a specific field in a collectionName
func (c *Collection) CreateIndex(fields interface{}, unique bool) {
	if c.Connection.Environment == "write" {
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
			index, err := c.Collection.Indexes().CreateOne(ctx, mod)
			if err == nil {
				sLog.Info("sMongo: Index %s%s created \n", c.prefix(), index)
				return
			}

			sLog.Error("sMongo: %sCreateIndex(): %s", c.prefix(), err)

			if count > 15 {
				sLog.Fatal("sMongo: restart App")
			}

			time.Sleep(time.Second)
			count++
		}
	}
}
