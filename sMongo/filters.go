package sMongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Filters struct {
	Fields bson.D
}

func Filter() *Filters { return &Filters{Fields: bson.D{}} }

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

// Set Ts, $gte: ts_from  $lt: ts_to
func (f *Filters) Ts(ts_from, ts_to int64) *Filters {
	return f.TsField(ts_from, ts_to, "ts")
}

func (f *Filters) TsIn(tss ...int64) *Filters {
	f.Int64_in("ts", tss)
	return f
}

// string
func (f *Filters) String_in(field string, values []string) *Filters {
	return f.Append(field, bson.D{{"$in", values}})
}

// ObjectId
func (f *Filters) ObjectId_in(field string, values []primitive.ObjectID) *Filters {
	return f.Append(field, bson.D{{"$in", values}})
}

// int
func (f *Filters) Int_in(field string, values []int) *Filters {
	return f.Append(field, bson.D{{"$in", values}})
}

func (f *Filters) Int_nin(field string, values []int) *Filters {
	return f.Append(field, bson.D{{"$nin", values}})
}

func (f *Filters) Int_gt(field string, value int) *Filters {
	return f.Append(field, bson.D{{"$gt", value}})
}

func (f *Filters) Int_gte(field string, value int) *Filters {
	return f.Append(field, bson.D{{"$gte", value}})
}

func (f *Filters) Int_lt(field string, value int) *Filters {
	return f.Append(field, bson.D{{"$lt", value}})
}

func (f *Filters) Int_lte(field string, value int) *Filters {
	return f.Append(field, bson.D{{"$lte", value}})
}

func (f *Filters) Int_gte_lte(field string, value_1, value_2 int) *Filters {
	return f.Append(field, bson.D{{"$gte", value_1}, {"$lte", value_2}})
}

// int64
func (f *Filters) Int64_in(field string, values []int64) *Filters {
	return f.Append(field, bson.D{{"$in", values}})
}

func (f *Filters) Int64_nin(field string, values []int64) *Filters {
	return f.Append(field, bson.D{{"$nin", values}})
}

func (f *Filters) Int64_gt(field string, value int64) *Filters {
	return f.Append(field, bson.D{{"$gt", value}})
}

func (f *Filters) Int64_gte(field string, value int64) *Filters {
	return f.Append(field, bson.D{{"$gte", value}})
}

func (f *Filters) Int64_lt(field string, value int64) *Filters {
	return f.Append(field, bson.D{{"$lt", value}})
}

func (f *Filters) Int64_lte(field string, value int64) *Filters {
	return f.Append(field, bson.D{{"$lte", value}})
}

func (f *Filters) Int64_gte_lt(field string, value_1, value_2 int64) *Filters {
	return f.Append(field, bson.D{{"$gte", value_1}, {"$lt", value_2}})
}

func (f *Filters) Int64_gte_lte(field string, value_1, value_2 int64) *Filters {
	return f.Append(field, bson.D{{"$gte", value_1}, {"$lte", value_2}})
}

// float64
func (f *Filters) Float64_in(field string, values []float64) *Filters {
	return f.Append(field, bson.D{{"$in", values}})
}

func (f *Filters) Float64_nin(field string, values []float64) *Filters {
	return f.Append(field, bson.D{{"$nin", values}})
}

func (f *Filters) Float64_gt(field string, value float64) *Filters {
	return f.Append(field, bson.D{{"$gt", value}})
}

func (f *Filters) Float64_gte(field string, value float64) *Filters {
	return f.Append(field, bson.D{{"$gte", value}})
}

func (f *Filters) Float64_lt(field string, value float64) *Filters {
	return f.Append(field, bson.D{{"$lt", value}})
}

func (f *Filters) Float64_lte(field string, value float64) *Filters {
	return f.Append(field, bson.D{{"$lte", value}})
}

func (f *Filters) Float64_gte_lte(field string, value_1, value_2 float64) *Filters {
	return f.Append(field, bson.D{{"$gte", value_1}, {"$lte", value_2}})
}

// *** Older code TODO: remove

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
