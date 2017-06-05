package gq

import (
	"github.com/graphql-go/graphql"
	"live/gq/resolvers"
	"live/gq/types"
)

var query = graphql.NewObject(graphql.ObjectConfig{
	Name: "episode",
	Fields: graphql.Fields{
		"episode": &graphql.Field{
			Type:        types.Episode,
			Description: "get episode",
			Args: graphql.FieldConfigArgument{
				"entity_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: resolvers.GetEpisode,
		},
	},
})
