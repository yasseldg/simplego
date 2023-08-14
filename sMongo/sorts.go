package sMongo

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Sorts struct {
	Fields bson.D
}

func Sort() *Sorts { return &Sorts{Fields: bson.D{}} }

// Append agrega un nuevo campo de ordenación con su dirección.
func (s *Sorts) Append(key string, value int) *Sorts {
	s.Fields = append(s.Fields, bson.E{Key: key, Value: value})
	return s
}

// Ascending agrega un campo para ordenar en dirección ascendente.
func (s *Sorts) Asc(key string) *Sorts {
	return s.Append(key, 1)
}

// Descending agrega un campo para ordenar en dirección descendente.
func (s *Sorts) Desc(key string) *Sorts {
	return s.Append(key, -1)
}

func (s *Sorts) TsAsc() *Sorts {
	return s.Asc("ts")
}

func (s *Sorts) TsDesc() *Sorts {
	return s.Desc("ts")
}
