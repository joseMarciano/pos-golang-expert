package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joseMarciano/pos-golang-expert/apis/internal/entity/product"
	"github.com/joseMarciano/pos-golang-expert/apis/internal/entity/user"
	"github.com/joseMarciano/pos-golang-expert/apis/internal/infra/database"
	"github.com/joseMarciano/pos-golang-expert/apis/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func init() {
	println("Main")
}

func main() {
	//fmt.Println(databaseconfig.GetDbConfig())
	//fmt.Println(authconfig.GetAuthConfig())
	//fmt.Println(serverconfig.GetServerConfig())

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&user.User{}, &product.Product{})
	productDB := database.NewProductDB(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger) // add a Logger Middleware
	r.Post("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8080", r)

}
