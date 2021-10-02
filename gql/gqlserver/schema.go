package gqlserver

import (
	"github.com/graphql-go/graphql"
)

type SchemaWrapper struct {
	productResolver *Resolver
	Schema          graphql.Schema
}

func NewSchemaWrapper() *SchemaWrapper {
	return &SchemaWrapper{}
}

func (s *SchemaWrapper) WithProductResolver(pr *Resolver) *SchemaWrapper {
	s.productResolver = pr

	return s
}

func (s *SchemaWrapper) Init() error {
	// add gql schema as needed
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "ProductGetter",
			Description: "All query related to getting product data",
			Fields: graphql.Fields{
				"ProductDetail": &graphql.Field{
					Type:        ProductType,
					Description: "Get product by ID",
					Args: graphql.FieldConfigArgument{
						"product_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: s.productResolver.GetProduct(),
				},
				"Products": &graphql.Field{
					Type:        graphql.NewList(ProductType),
					Description: "Get products",
					Resolve:     s.productResolver.GetProducts(),
				},
			},
		}),
		// uncomment this and add objects for mutation
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:        "ProductSetter",
			Description: "All query related to modify product data",
			Fields: graphql.Fields{
				"CreateProduct": &graphql.Field{
					Type:        graphql.Boolean,
					Description: "Create product",
					Args: graphql.FieldConfigArgument{
						"product_name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"product_shop_name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"product_price": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"product_image": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: s.productResolver.CreateProducts(),
				},
				"UpdateProduct": &graphql.Field{
					Type:        graphql.Boolean,
					Description: "Update product by ID",
					Args: graphql.FieldConfigArgument{
						"product_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"product_name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"product_shop_name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"product_price": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"product_image": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: s.productResolver.UpdateProducts(),
				},
			},
		}),
	})

	if err != nil {
		return err
	}

	s.Schema = schema

	return nil
}
