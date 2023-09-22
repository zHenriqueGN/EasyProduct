package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/zHenriqueGN/EasyProduct/configs"
	"github.com/zHenriqueGN/EasyProduct/internal/infra/database"
	"github.com/zHenriqueGN/EasyProduct/internal/infra/repository"
	"github.com/zHenriqueGN/EasyProduct/internal/infra/webserver/handlers"
)

func main() {
	cfg := configs.LoadConfig()
	DB := database.ConnectToDatabase(cfg.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName))
	productRepository := repository.NewProductRepository(DB)
	productHandler := handlers.NewProductHandler(productRepository)
	userRepository := repository.NewUserRepository(DB)
	userHandler := handlers.NewUserHandler(userRepository)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", cfg.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", cfg.JWTExpiresIn))
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cfg.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})
	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/getjwt", userHandler.GetJWT)
	r.Handle("/docs/*", http.StripPrefix("/docs/", http.FileServer(http.Dir("./api/swaggerui"))))
	fmt.Println("listening on http://localhost:8000")
	http.ListenAndServe(":8000", r)
}
