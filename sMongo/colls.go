package sMongo

type Colls map[string]CollManager

func NewColls() Colls {
	return make(Colls)
}

func (colls Colls) Get(env, connection, database, collection, interval string, indexes Indexes) CollManager {
	if coll, ok := colls[collection]; ok {
		return coll
	}

	c := Coll(env, connection, database, collection, interval)
	c.CreateIndexes(indexes)
	colls[collection] = c
	return c
}
