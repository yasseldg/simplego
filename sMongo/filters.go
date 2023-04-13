package sMongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetIdsFilter
func GetIdsFilter(ids []string) bson.D {
	objIds := []primitive.ObjectID{}
	for _, id := range ids {
		objId, _ := primitive.ObjectIDFromHex(id)
		objIds = append(objIds, objId)
	}
	return bson.D{{"_id", bson.D{{"$in", objIds}}}}
}

// GetTsFilter, $gte: tsFrom  $lt: tsTo
func GetTsFilter(tsFrom, tsTo int64) bson.D {
	if tsFrom > 0 {
		if tsTo > 0 {
			return bson.D{{"ts", bson.D{{"$gte", tsFrom}, {"$lt", tsTo}}}}
		}
		return bson.D{{"ts", bson.D{{"$gte", tsFrom}}}}
	}

	if tsTo > 0 {
		return bson.D{{"ts", bson.D{{"$lt", tsTo}}}}
	}
	return bson.D{}
}
