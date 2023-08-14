package sMongo

import (
	"github.com/yasseldg/mgm/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollManager struct {
	Collection
	filters  Filters
	sorts    Sorts
	limit    int64
	pipeline mongo.Pipeline
}

func (c *CollManager) Filters(f Filters) *CollManager {
	c.filters = f
	return c
}

func (c *CollManager) Sorts(s Sorts) *CollManager {
	c.sorts = s
	return c
}

func (c *CollManager) Limit(l int64) *CollManager {
	c.limit = l
	return c
}

func (c *CollManager) Pipeline(p mongo.Pipeline) *CollManager {
	c.pipeline = p
	return c
}

func (c *CollManager) Count() (int64, error) {
	return c.Collection.Count(c.filters.Fields, options.Count())
}

func (c *CollManager) Find(objs interface{}) error {
	opts := options.Find().SetSort(c.sorts.Fields)
	if c.limit > 0 {
		opts.SetLimit(c.limit)
	}
	return c.Collection.Find(c.filters.Fields, opts, objs)
}

func (c *CollManager) FindOne(obj mgm.Model) error {
	return c.Collection.FindOne(c.filters.Fields, *options.FindOne().SetSort(c.sorts.Fields), obj)
}

func (c *CollManager) Agregates(docs interface{}) error {
	return c.Collection.Agregates(c.pipeline, docs)
}
