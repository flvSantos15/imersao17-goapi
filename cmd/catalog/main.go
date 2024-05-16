package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/flvsantos15/imersao17/goapi/internal/database"
	"github.com/flvsantos15/imersao17/goapi/internal/services"
	"github.com/flvsantos15/imersao17/goapi/internal/webserver"
	"github.com/go-chi/chi/v5"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersao17")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := services.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productService := services.NewProductService(*productDB)

	webCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	webProductHandler := webserver.NewProductHandler(productService)

	c := chi.NewRouter()
	c.Get("/category/{id}", webCategoryHandler.GetCategory)
	c.Get("/category", webCategoryHandler.GetCategories)
	c.Post("/category", webCategoryHandler.CreateCategory)

	c.Get("/product/{id}", webProductHandler.GetProduct)
	c.Get("/product", webProductHandler.GetProducts)
	c.Get("product/category/{categoryID}", webProductHandler.GetProductByCategoryID)
	c.Post("/product", webProductHandler.CreateProduct)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)
}
