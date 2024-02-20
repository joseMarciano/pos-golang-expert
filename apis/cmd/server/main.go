package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/joseMarciano/pos-golang-expert/apis/configs/authconfig"
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

	authConfiguration := authconfig.GetAuthConfig()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&user.User{}, &product.Product{})
	productDB := database.NewProductDB(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUserDB(db)
	userHandler := handlers.NewUserHandler(userDB, jwtauth.New("HS256", []byte(authConfiguration.JwtSecret()), nil), authConfiguration.JwtExpiresIn())

	r := chi.NewRouter()
	r.Use(middleware.Logger) // add a Logger Middleware
	r.Get("/products", productHandler.ListProducts)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	r.Post("/users", userHandler.Create)
	r.Post("/users/jwt", userHandler.GetJwt)
	http.ListenAndServe(":8080", r)

}
