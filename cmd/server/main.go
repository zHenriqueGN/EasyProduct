package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	userHandler := handlers.NewUserHandler(userRepository, cfg.TokenAuth, cfg.JWTExperesIn)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products", productHandler.GetProducts)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)
	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/getjwt", userHandler.GetJWT)
	http.ListenAndServe(":8000", r)
}
