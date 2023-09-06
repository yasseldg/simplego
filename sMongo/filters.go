package sMongo

import (
	"github.com/yasseldg/simplego/sLog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Filters struct {
	Fields bson.D
}

func Filter() *Filters { return &Filters{Fields: bson.D{}} }

func (f *Filters) Log(msg string) {
	sLog.Debug("%s: Filters: %v", msg, f.Fields)
}

// Append adds a new filter field.
func (f *Filters) Append(key string, value interface{}) *Filters {
	f.Fields = append(f.Fields, bson.E{Key: key, Value: value})
	return f
}

// Field like Ts, $gte: ts_from  $lt: ts_to
func (f *Filters) TsField(ts_from, ts_to int64, field string) *Filters {
	if ts_from > 0 {
		if ts_to > 0 {
			return f.Int64_gte_lt(field, ts_from, ts_to)
		}
		return f.Int64_gte(field, ts_from)
	}
	if ts_to > 0 {
		return f.Int64_lt(field, ts_to)
	}
	return f
}

// ----- Ts Filters

// Set Ts, $gte: ts_from  $lt: ts_to
func (f *Filters) Ts(ts_from, ts_to int64) *Filters {
	return f.TsField(ts_from, ts_to, "ts")
}

func (f *Filters) TsIn(tss ...int64) *Filters {
	f.Int64_in("ts", tss...)
	return f
}

// ----- States Filters

func (f *Filters) States(states ...string) *Filters { f.String_in("st", states...); return f }

func (f *Filters) NotStates(states ...string) *Filters { f.String_nin("st", states...); return f }

// -----

// string
func (f *Filters) String_in(field string, values ...string) *Filters {
	return f.Append(field, bson.D{{Key: "$in", Value: values}})
}

func (f *Filters) String_nin(field string, values ...string) *Filters {
	return f.Append(field, bson.D{{Key: "$nin", Value: values}})
}

// ObjectId
func (f *Filters) ObjectId(field string, value primitive.ObjectID) *Filters {
	return f.Append(field, value)
}

func (f *Filters) ObjectId_in(field string, values []primitive.ObjectID) *Filters {
	return f.Append(field, bson.D{{Key: "$in", Value: values}})
}

// int
func (f *Filters) Int(field string, value int) *Filters {
	return f.Append(field, value)
}

func (f *Filters) Int_in(field string, values ...int) *Filters {
	return f.Append(field, bson.D{{Key: "$in", Value: values}})
}

func (f *Filters) Int_nin(field string, values ...int) *Filters {
	return f.Append(field, bson.D{{Key: "$nin", Value: values}})
}

func (f *Filters) Int_gt(field string, value int) *Filters {
	return f.Append(field, bson.D{{Key: "$gt", Value: value}})
}

func (f *Filters) Int_gte(field string, value int) *Filters {
	return f.Append(field, bson.D{{Key: "$gte", Value: value}})
}

func (f *Filters) Int_lt(field string, value int) *Filters {
	return f.Append(field, bson.D{{Key: "$lt", Value: value}})
}

func (f *Filters) Int_lte(field string, value int) *Filters {
	return f.Append(field, bson.D{{Key: "$lte", Value: value}})
}

func (f *Filters) Int_gte_lte(field string, value_1, value_2 int) *Filters {
	return f.Append(field, bson.D{{Key: "$gte", Value: value_1}, {Key: "$lte", Value: value_2}})
}

// int64
func (f *Filters) Int64(field string, value int64) *Filters {
	return f.Append(field, value)
}

func (f *Filters) Int64_in(field string, values ...int64) *Filters {
	return f.Append(field, bson.D{{Key: "$in", Value: values}})
}

func (f *Filters) Int64_nin(field string, values ...int64) *Filters {
	return f.Append(field, bson.D{{Key: "$nin", Value: values}})
}

func (f *Filters) Int64_gt(field string, value int64) *Filters {
	return f.Append(field, bson.D{{Key: "$gt", Value: value}})
}

func (f *Filters) Int64_gte(field string, value int64) *Filters {
	return f.Append(field, bson.D{{Key: "$gte", Value: value}})
}

func (f *Filters) Int64_lt(field string, value int64) *Filters {
	return f.Append(field, bson.D{{Key: "$lt", Value: value}})
}

func (f *Filters) Int64_lte(field string, value int64) *Filters {
	return f.Append(field, bson.D{{Key: "$lte", Value: value}})
}

func (f *Filters) Int64_gte_lt(field string, value_1, value_2 int64) *Filters {
	return f.Append(field, bson.D{{Key: "$gte", Value: value_1}, {Key: "$lt", Value: value_2}})
}

func (f *Filters) Int64_gte_lte(field string, value_1, value_2 int64) *Filters {
	return f.Append(field, bson.D{{Key: "$gte", Value: value_1}, {Key: "$lte", Value: value_2}})
}

// float64
func (f *Filters) Float64(field string, value float64) *Filters {
	return f.Append(field, value)
}

func (f *Filters) Float64_in(field string, values ...float64) *Filters {
	return f.Append(field, bson.D{{Key: "$in", Value: values}})
}

func (f *Filters) Float64_nin(field string, values ...float64) *Filters {
	return f.Append(field, bson.D{{Key: "$nin", Value: values}})
}

func (f *Filters) Float64_gt(field string, value float64) *Filters {
	return f.Append(field, bson.D{{Key: "$gt", Value: value}})
}

func (f *Filters) Float64_gte(field string, value float64) *Filters {
	return f.Append(field, bson.D{{Key: "$gte", Value: value}})
}

func (f *Filters) Float64_lt(field string, value float64) *Filters {
	return f.Append(field, bson.D{{Key: "$lt", Value: value}})
}

func (f *Filters) Float64_lte(field string, value float64) *Filters {
	return f.Append(field, bson.D{{Key: "$lte", Value: value}})
}

func (f *Filters) Float64_gte_lte(field string, value_1, value_2 float64) *Filters {
	return f.Append(field, bson.D{{Key: "$gte", Value: value_1}, {Key: "$lte", Value: value_2}})
}

// *** Older code TODO: remove

// GetIdsFilter
func GetIdsFilter(ids []string) bson.D {
	objIds := []primitive.ObjectID{}
	for _, id := range ids {
		objId, _ := primitive.ObjectIDFromHex(id)
		objIds = append(objIds, objId)
	}
	return bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: objIds}}}}
}

// GetTsFilter, $gte: tsFrom  $lt: tsTo
func GetTsFilter(tsFrom, tsTo int64) bson.D {
	if tsFrom > 0 {
		if tsTo > 0 {
			return bson.D{{Key: "ts", Value: bson.D{{Key: "$gte", Value: tsFrom}, {Key: "$lt", Value: tsTo}}}}
		}
		return bson.D{{Key: "ts", Value: bson.D{{Key: "$gte", Value: tsFrom}}}}
	}

	if tsTo > 0 {
		return bson.D{{Key: "ts", Value: bson.D{{Key: "$lt", Value: tsTo}}}}
	}
	return bson.D{}
}
