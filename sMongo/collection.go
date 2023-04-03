package sMongo

import (
	"fmt"

	"github.com/yasseldg/simplego/sEnv"
	"github.com/yasseldg/simplego/sLog"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	Collection     *mgm.Collection
	Connection     *ConnectionParams
	DatabaseName   string
	CollectionName string
}

// New
func New(env, connection, databaseName, collectionName string) *Collection {
	var c Collection

	c.Connection = getConnection(sEnv.Get(fmt.Sprint("CONN_", env), connection))
	c.DatabaseName = sEnv.Get(fmt.Sprint("DB_", env), databaseName)
	c.CollectionName = sEnv.Get(fmt.Sprint("COLL_", env), collectionName)

	c.Collection = getMgmColl(c.Connection, c.DatabaseName, c.CollectionName)

	return &c
}

func (c *Collection) prefix() string {
	return fmt.Sprintf("%s.%s.", c.DatabaseName, c.CollectionName)
}

// Clone, create Connection to a new collectionName in the same DB
func (c *Collection) Clone(collectionName string) *Collection {
	clone := *c
	clone.CollectionName = collectionName
	clone.Collection = getMgmColl(c.Connection, c.DatabaseName, c.CollectionName)
	return &clone
}

// CreateIndex, create an index for a specific field in a collectionName
func (c *Collection) CreateIndex(fields interface{}, unique bool) {

	err := cicleCreateIndex(c.Collection, fields, unique)
	if err != nil {
		sLog.Fatal("Mongo DB: %s CreateIndex failure %s", c.prefix(), err)
	}
	sLog.Info("Init %s completed ", c.prefix())
}

// Create
func (c *Collection) Create(obj mgm.Model) bool {

	err := c.Collection.Create(obj)

	if err != nil {
		sLog.Error("%sCreate(obj): %s  ..  obj: %#v", c.prefix(), err, obj)
		return false
	}
	return true
}

// CreateMany
func (c *Collection) CreateMany(docs []mgm.Model) bool {

	if len(docs) > 0 {
		err := c.Collection.CreateMany(docs)
		if err != nil {
			sLog.Error("%sCreateMany(objs): %s  ..  objs: %#v", c.prefix(), err, docs)
			return false
		}
	}
	return true
}

// Update
func (c *Collection) Update(obj mgm.Model) bool {

	err := c.Collection.Update(obj)
	if err != nil {
		sLog.Error("%sUpdate(&obj): %s  ..  obj: %#v", c.prefix(), err, obj)
		return false
	}
	return true
}

// Count
func (c *Collection) Count(filter interface{}, opts *options.CountOptions) (int64, error) {

	count, err := c.Collection.SimpleCount(filter, opts)
	if err != nil {
		sLog.Error("%sSimpleCount(filter, opts): %s  ..  opts: %#v", c.prefix(), err, opts)
		return 0, err
	}
	return count, nil
}

// Find
func (c *Collection) Find(filters interface{}, opts *options.FindOptions, objs interface{}) error {

	err := c.Collection.SimpleFind(objs, filters, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			sLog.Debug("%sSimpleFind(objs, filter, opts): %s", c.prefix(), err)
			return err
		}
		sLog.Error("%sSimpleFind(objs, filter, opts): %s  ..  filter: %#v  ..  opts: %#v", c.prefix(), err, filters, opts)
		return err
	}
	return nil
}

// FindOne
func (c *Collection) FindOne(filters interface{}, opts options.FindOneOptions, obj mgm.Model) error {

	err := c.Collection.First(filters, obj, &opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			sLog.Debug("%sFirst(filters, obj, &opts): %s", c.prefix(), err)
			return err
		}
		sLog.Error("%sFirst(filters, obj, &opts): %s  ..  filter: %#v  ..  opts: %#v", c.prefix(), err, filters, opts)
		return err
	}
	return nil
}

// FindById
func (c *Collection) FindById(id interface{}, obj mgm.Model) error {

	err := c.Collection.FindByID(id, obj)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			sLog.Debug("%sFindByID(id, obj): %s", c.prefix(), err)
			return err
		}
		sLog.Error("%sFindByID(id, obj): %s  ..  id: %s", c.prefix(), err, id)
		return err
	}
	return nil
}
