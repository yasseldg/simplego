package sMongo

import (
	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type TsModel struct {
	mgm.DefaultModel `bson:",inline"`
	UnixTs           int64 `bson:"ts" json:"ts"`
}

// First, $gte: tsFrom  $lt: tsTo, tsFrom = tsTo = 0 for "ts" first object,
func (c *Collection) First(tsFrom, tsTo int64, obj mgm.Model) error {
	return c.FindOne(Filter().Ts(tsFrom, tsTo).Fields, *options.FindOne().SetSort(Sort().TsAsc().Fields), obj)
}

// Last, $gte: tsFrom  $lt: tsTo, tsFrom = tsTo = 0 for first
func (c *Collection) FirstTs(tsFrom, tsTo int64) int64 {
	var obj TsModel
	err := c.First(tsFrom, tsTo, &obj)
	if err != nil {
		return 0
	}
	return int64(obj.UnixTs)
}

// Last, $gte: tsFrom  $lt: tsTo, tsFrom = tsTo = 0 for "ts" Last object,
func (c *Collection) Last(tsFrom, tsTo int64, obj mgm.Model) error {
	return c.FindOne(Filter().Ts(tsFrom, tsTo).Fields, *options.FindOne().SetSort(Sort().TsDesc().Fields), obj)
}

// Last, $gte: tsFrom  $lt: tsTo, tsFrom = tsTo = 0 for last
func (c *Collection) LastTs(tsFrom, tsTo int64) int64 {
	var obj TsModel
	err := c.Last(tsFrom, tsTo, &obj)
	if err != nil {
		return 0
	}
	return int64(obj.UnixTs)
}
