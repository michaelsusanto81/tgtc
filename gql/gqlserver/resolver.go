package gqlserver

import (
	"github.com/graphql-go/graphql"
	"github.com/radityaqb/tgtc/backend/dictionary"
	"github.com/radityaqb/tgtc/backend/service"
)

type Resolver struct {
}

func NewResolver() *Resolver {
	return &Resolver{}
}

func (r *Resolver) GetProduct() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["product_id"].(int)

		// update to use Usecase from previous session
		return service.GetProduct(id)
	}
}

func (r *Resolver) GetProducts() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		return service.GetProducts()
	}
}

func (r *Resolver) CreateProducts() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["product_name"].(string)
		shop, _ := p.Args["product_shop_name"].(string)
		price, _ := p.Args["product_price"].(float64)
		image, _ := p.Args["product_image"].(string)

		req := dictionary.Product{
			Name:         name,
			ShopName:     shop,
			ProductPrice: price,
			ImageURL:     image,
		}

		_, err := service.CreateProduct(req)
		if err != nil {
			return false, err
		}

		return true, nil
	}
}

func (r *Resolver) UpdateProducts() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["product_id"].(int)
		name, _ := p.Args["product_name"].(string)
		shop, _ := p.Args["product_shop_name"].(string)
		price, _ := p.Args["product_price"].(float64)
		image, _ := p.Args["product_image"].(string)

		req := dictionary.Product{
			ID:           int64(id),
			Name:         name,
			ShopName:     shop,
			ProductPrice: price,
			ImageURL:     image,
		}

		_, err := service.UpdateProduct(req)
		if err != nil {
			return false, err
		}

		return true, nil
	}
}
