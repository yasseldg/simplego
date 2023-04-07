package sMongo

import (
	"github.com/yasseldg/simplego/sLog"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create
func (c *Collection) Create(obj mgm.Model) bool {

	err := c.Collection.Create(obj)

	if err != nil {
		sLog.Error("sMongo: %sCreate(obj): %s  ..  obj: %#v", c.prefix(), err, obj)
		return false
	}
	return true
}

// CreateMany
func (c *Collection) CreateMany(docs []mgm.Model) bool {

	if len(docs) > 0 {
		err := c.Collection.CreateMany(docs)
		if err != nil {
			sLog.Error("sMongo: %sCreateMany(objs): %s  ..  objs: %#v", c.prefix(), err, docs)
			return false
		}
	}
	return true
}

// Update
func (c *Collection) Update(obj mgm.Model) bool {

	err := c.Collection.Update(obj)
	if err != nil {
		sLog.Error("sMongo: %sUpdate(&obj): %s  ..  obj: %#v", c.prefix(), err, obj)
		return false
	}
	return true
}

// Count
func (c *Collection) Count(filter interface{}, opts *options.CountOptions) (int64, error) {

	count, err := c.Collection.SimpleCount(filter, opts)
	if err != nil {
		sLog.Error("sMongo: %sSimpleCount(filter, opts): %s  ..  opts: %#v", c.prefix(), err, opts)
		return 0, err
	}
	return count, nil
}

// Find
func (c *Collection) Find(filters interface{}, opts *options.FindOptions, objs interface{}) error {

	err := c.Collection.SimpleFind(objs, filters, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			sLog.Debug("sMongo: %sSimpleFind(objs, filter, opts): %s", c.prefix(), err)
			return err
		}
		sLog.Error("sMongo: %sSimpleFind(objs, filter, opts): %s  ..  filter: %#v  ..  opts: %#v", c.prefix(), err, filters, opts)
		return err
	}
	return nil
}

// FindOne
func (c *Collection) FindOne(filters interface{}, opts options.FindOneOptions, obj mgm.Model) error {

	err := c.Collection.First(filters, obj, &opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			sLog.Debug("sMongo: %sFirst(filters, obj, &opts): %s", c.prefix(), err)
			return err
		}
		sLog.Error("sMongo: %sFirst(filters, obj, &opts): %s  ..  filter: %#v  ..  opts: %#v", c.prefix(), err, filters, opts)
		return err
	}
	return nil
}

// FindById
func (c *Collection) FindById(id interface{}, obj mgm.Model) error {

	err := c.Collection.FindByID(id, obj)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			sLog.Debug("sMongo: %sFindByID(id, obj): %s", c.prefix(), err)
			return err
		}
		sLog.Error("sMongo: %sFindByID(id, obj): %s  ..  id: %s", c.prefix(), err, id)
		return err
	}
	return nil
}
