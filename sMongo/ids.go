package sMongo

import (
	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetID(id interface{}) primitive.ObjectID {
	if strId, ok := id.(string); ok {
		objId, _ := primitive.ObjectIDFromHex(strId)
		return objId
	}
	return id.(primitive.ObjectID)
}

func GetIDs(objs []mgm.DefaultModel) []primitive.ObjectID {
	objIds := []primitive.ObjectID{}
	for _, obj := range objs {
		objIds = append(objIds, obj.ID)
	}
	return objIds
}
