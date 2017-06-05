package resolvers

import (
	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
	"live/db"
)

type Episode struct {
	Entity_id int    `json:"entity_id"`
	Title     string `json:"title"`
}

func GetEpisode(params graphql.ResolveParams) (interface{}, error) {
	entity_id := params.Args["entity_id"].(int)
	var res Episode

	_ = db.Mongo.DB("demo").C("episode").FindId(entity_id).Select(bson.M{
		"entity_id": 1,
		"title":     1}).One(&res)

	return res, nil
}
