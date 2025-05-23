package repository

import (
	"database/sql"
	"fmt"

	"github.com/rochaandrey/restful-API.git/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func newProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "select * from product p"
	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObject model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObject.ID,
			&productObject.Name,
			&productObject.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObject)
	}

	rows.Close()

	return productList, nil
}
