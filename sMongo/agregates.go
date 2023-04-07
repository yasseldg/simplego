package sMongo

import (
	"context"

	"github.com/yasseldg/simplego/sLog"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *Collection) Agregates(pipeline mongo.Pipeline, docs interface{}) error {
	return c.AgregatesWithCtx(mgm.Ctx(), pipeline, docs)
}

func (c *Collection) AgregatesWithCtx(ctx context.Context, pipeline mongo.Pipeline, docs interface{}) error {

	cursor, err := c.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		sLog.Error("sMongo: %sAgregatesWithCtx: ", c.prefix(), err.Error())
	} else {
		err = cursor.All(ctx, docs)
		if err != nil {
			sLog.Error("sMongo: %sAgregatesWithCtx: cursor.All(): ", c.prefix(), err.Error())
		}
	}
	return err
}

func (c *Collection) AgregatesCount(pipeline mongo.Pipeline) ([]bson.M, error) {
	return c.AgregatesCountWithCtx(mgm.Ctx(), pipeline)
}

func (c *Collection) AgregatesCountWithCtx(ctx context.Context, pipeline mongo.Pipeline) ([]bson.M, error) {

	var result []bson.M

	cursor, err := c.Collection.Aggregate(mgm.Ctx(), pipeline)
	if err != nil {
		sLog.Error("sMongo: %sAgregatesCountWithCtx: ", c.prefix(), err.Error())
	} else {
		err = cursor.All(mgm.Ctx(), &result)
		if err != nil {
			sLog.Error("sMongo: %sAgregatesCountWithCtx: cursor.All(): ", c.prefix(), err.Error())
		}
	}
	return result, err
}
