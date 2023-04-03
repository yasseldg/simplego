package sMongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetIdsFilter
func GetIdsFilter(ids []string) interface{} {
	objIds := []primitive.ObjectID{}
	for _, id := range ids {
		objId, _ := primitive.ObjectIDFromHex(id)
		objIds = append(objIds, objId)
	}
	return bson.M{"_id": bson.M{"$in": objIds}}
}

// GetTsFilter, $gte: tsFrom  $lt: tsTo
func GetTsFilter(tsFrom, tsTo int64) bson.M {
	if tsFrom > 0 {
		if tsTo > 0 {
			return bson.M{"ts": bson.M{"$gte": tsFrom, "$lt": tsTo}}
		}
		return bson.M{"ts": bson.M{"$gte": tsFrom}}
	}

	if tsTo > 0 {
		return bson.M{"ts": bson.M{"$lt": tsTo}}
	}
	return bson.M{}
}
