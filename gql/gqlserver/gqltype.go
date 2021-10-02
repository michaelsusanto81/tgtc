package gqlserver

import "github.com/graphql-go/graphql"

var ProductType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Product",
		Description: "Detail of the product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"product_price": &graphql.Field{
				Type: graphql.Float,
			},
			"image_url": &graphql.Field{
				Type: graphql.String,
			},
			"shop_name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var ResultType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Result",
		Description: "Result of operation",
		Fields: graphql.Fields{
			"success": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)
