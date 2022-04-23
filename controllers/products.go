package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/Gabriel-Rabeloo/golang-web-api/models"
)

var temps = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		temps.ExecuteTemplate(w, "NotFound", nil)
		return
	}
	products := models.FindProducts()
	temps.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temps.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		priceStr := r.FormValue("price")
		quantityStr := r.FormValue("quantity")

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Println("Error in convertion price:", err)
		}

		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			log.Println("Error in quantity price:", err)
		}

		models.InsertProduct(name, description, price, quantity)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProduct(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := models.FindProductById(id)
	temps.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idStr := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		priceStr := r.FormValue("price")
		quantityStr := r.FormValue("quantity")

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Println("Error in convertion price:", err)
		}

		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			log.Println("Error in convertion quantity:", err)
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println("Error in convertion id:", err)
		}

		fmt.Println(id, name, description, price, quantity)
		models.UpdateProduct(id, name, description, price, quantity)
	}

	http.Redirect(w, r, "/", 301)
}
