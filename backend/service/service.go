package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/radityaqb/tgtc/backend/database"
	"github.com/radityaqb/tgtc/backend/dictionary"
)

func SampleFunction() {
	fmt.Printf("My Service!")

	// // you can connect and
	// // get current database connection
	// db := database.GetDB()

	// // construct query
	// query := `
	// SELECT something FROM table_something WHERE id = $1
	// `
	// // actual query process
	// row = db.QueryRow(query, paramID)

	// // read query result, and assign to variable(s)
	// err = row.Scan(&ID, &name)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	idstring := r.URL.Query().Get("id")

	idInt64, err := strconv.ParseInt(idstring, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	db := database.GetDB()

	query := `
		SELECT
			product_id,
			product_name,
			product_price,
			product_image,
			shop_name
		FROM
			products
		WHERE
			product_id = $1
	`

	row := db.QueryRow(query, idInt64)

	var p dictionary.Product
	err = row.Scan(
		&p.ID,
		&p.Name,
		&p.ProductPrice,
		&p.ImageURL,
		&p.ShopName,
	)
	var response dictionary.APIResponse = dictionary.APIResponse{}
	if err != nil {
		response.Error = err.Error()
		json.NewEncoder(w).Encode(response)
	} else {
		response.Data = p
		json.NewEncoder(w).Encode(response)
	}
}
