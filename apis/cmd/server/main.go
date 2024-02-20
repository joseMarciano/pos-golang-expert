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
	auth := jwtauth.New("HS256", []byte(authConfiguration.JwtSecret()), nil)
	userHandler := handlers.NewUserHandler(userDB, auth, authConfiguration.JwtExpiresIn())

	r := chi.NewRouter()
	r.Use(middleware.Recoverer) // gracefull handle panic errors in my route
	r.Use(middleware.Logger)    // add a Logger Middleware
	r.Use(MyOwnMiddleware)      // using my own middleware

	r.Route("/products", func(router chi.Router) {
		router.Use(jwtauth.Verifier(auth)) // insert the token in the Context
		router.Use(jwtauth.Authenticator)  // authenticate the use
		router.Get("/", productHandler.ListProducts)
		router.Post("/", productHandler.CreateProduct)
		router.Get("/{id}", productHandler.GetProduct)
		router.Put("/{id}", productHandler.UpdateProduct)
		router.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.Create)
	r.Post("/users/jwt", userHandler.GetJwt)
	http.ListenAndServe(":8080", r)

}

func MyOwnMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("MyOwnMiddleware")
		next.ServeHTTP(w, r)
	})
}
