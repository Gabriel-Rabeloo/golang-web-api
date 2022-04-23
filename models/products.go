package models

import (
	"github.com/Gabriel-Rabeloo/golang-web-api/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func FindProducts() []Product {
	db := db.ConnectDataBase()

	selectAllProducts, err := db.Query("SELECT * FROM products ORDER BY id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity
		p.Id = id

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func InsertProduct(name, description string, price float64, quantity int) {
	db := db.ConnectDataBase()

	queryInsert, err := db.Prepare("INSERT INTO PRODUCTS(name, description, price, quantity) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	queryInsert.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectDataBase()

	deleteQuery, err := db.Prepare("DELETE FROM products WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteQuery.Exec(id)
	defer db.Close()
}

func FindProductById(id string) Product {
	db := db.ConnectDataBase()

	dbProduct, err := db.Query("SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for dbProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = dbProduct.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Quantity = quantity
		product.Price = price
		product.Description = description
	}
	defer db.Close()
	return product
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := db.ConnectDataBase()

	queryUpdate, err := db.Prepare("UPDATE products SET name=$2, description=$3, price=$4, quantity=$5 WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	queryUpdate.Exec(id, name, description, price, quantity)
}
