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

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", r)
}
