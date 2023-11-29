package sMongo

import (
	"fmt"

	"github.com/yasseldg/mgm/v4"

	"github.com/yasseldg/simplego/sEnv"
	"github.com/yasseldg/simplego/sLog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	connection ConnectionParams

	client *mongo.Client

	databases DatabasesMap
}
type ClientsMap map[string]*Client

type Database struct {
	database *mongo.Database

	collections CollectionsMap
}
type DatabasesMap map[string]*Database

type CollectionsMap map[string]*CollManager

type Manager struct {
	clients ClientsMap
}

func NewManager() Manager {
	return Manager{clients: make(ClientsMap)}
}

func (m Manager) Log() {
	println()
	for _, client := range m.clients {
		for _, database := range client.databases {
			for _, collection := range database.collections {
				sLog.Info("client: %s  ..  database: %s  ..  coll: %s \n", client.connection.Environment, database.database.Name(), collection.Collection.Collection.Name())
			}
		}
	}
}

func (m *Manager) GetColl(env, conn_name, db_name, coll_name string, indexes ...Index) (CollManager, error) {

	client, err := m.getClient(env, conn_name)
	if err != nil {
		return CollManager{}, err
	}

	return client.getColl(db_name, coll_name, indexes...)
}

func (m *Manager) getClient(env, conn_name string) (*Client, error) {

	client := m.clients.get(conn_name)
	if client != nil {
		return client, nil
	}

	return m.setClient(env, conn_name)
}

func (m *Manager) setClient(env, conn_name string) (*Client, error) {

	mgm.SetDefaultConfig(getCtx(env))

	conn := getConnection(sEnv.Get(fmt.Sprint("CONN_", env), conn_name))

	client, err := mgm.NewClient(conn.getClientOpt())
	if err != nil {
		err := fmt.Errorf(" mgm.NewClient() for env: %s  ..  conn_name: %s  ..  err: %s", env, conn_name, err)
		return nil, err
	}

	m.clients[conn_name] = &Client{
		connection: *conn,
		client:     client,
		databases:  make(DatabasesMap),
	}
	return m.clients[conn_name], nil
}

func (c *Client) getColl(db_name, coll_name string, indexes ...Index) (CollManager, error) {

	db, err := c.getDatabase(db_name)
	if err != nil {
		return CollManager{}, err
	}

	coll, err := db.getCollection(coll_name, &c.connection)
	if err != nil {
		return CollManager{}, err
	}

	if c.connection.Environment == "write" {
		coll.CreateIndexes(indexes)
	}

	return coll, nil
}

func (c *Client) getDatabase(db_name string) (*Database, error) {
	db := c.databases.get(db_name)
	if db != nil {
		return db, nil
	}

	return c.setDatabase(db_name)
}

func (c *Client) setDatabase(db_name string) (*Database, error) {

	db := c.client.Database(db_name)
	if db == nil {
		err := fmt.Errorf(" client.Database( %s ) is nil", db_name)
		return nil, err
	}

	c.databases[db.Name()] = &Database{
		database:    db,
		collections: make(CollectionsMap),
	}
	return c.databases[db.Name()], nil
}

func (db *Database) getCollection(coll_name string, conn *ConnectionParams) (CollManager, error) {
	coll := db.collections.get(coll_name)
	if coll != nil {
		return *coll, nil
	}

	return db.setCollection(coll_name, conn)
}

func (db *Database) setCollection(coll_name string, conn *ConnectionParams) (CollManager, error) {

	coll := mgm.NewCollection(db.database, coll_name)
	if coll == nil {
		err := fmt.Errorf(" mgm.NewCollection( %s , %s ) is nil", db.database.Name(), coll_name)
		return CollManager{}, err
	}

	db.collections[coll.Name()] = &CollManager{
		Collection: Collection{Collection: coll, Database: db.database, Connection: conn},
		filters:    Filters{Fields: bson.D{}},
		sorts:      Sorts{Fields: bson.D{}},
		limit:      0,
		pipeline:   mongo.Pipeline{},
	}

	return *db.collections[coll.Name()], nil
}

func (cs ClientsMap) get(env string) *Client {
	if c, ok := cs[env]; ok {
		return c
	}
	return nil
}

func (dbs DatabasesMap) get(db_name string) *Database {
	if db, ok := dbs[db_name]; ok {
		return db
	}
	return nil
}

func (colls CollectionsMap) get(coll_name string) *CollManager {
	if c, ok := colls[coll_name]; ok {
		return c
	}
	return nil
}
