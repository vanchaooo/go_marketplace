package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vanchaooo/go-marketplace/internal/config"
	"github.com/vanchaooo/go-marketplace/internal/handler"
	"github.com/vanchaooo/go-marketplace/internal/repository"
	"github.com/vanchaooo/go-marketplace/internal/service"
)

func main() {
	// Config
	config, err := config.Load()
	if err != nil {
		log.Fatalf("Config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, config.DatabaseURL)
	if err != nil {
		log.Fatalf("database: не удалось создать пул подключений: %v", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("database: не удалось подключиться к PostgreSQL: %v", err)
	}



	// Маршрутизатор
	productRepository := repository.NewProductRepository(pool)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.Health)


	mux.HandleFunc("POST /api/v1/products", productHandler.CreateProduct)
	mux.HandleFunc("GET /api/v1/products", productHandler.GetAllProducts)
	mux.HandleFunc("GET /api/v1/products/{id}", productHandler.GetProduct)
	mux.HandleFunc("PUT /api/v1/products/{id}", productHandler.UpdateProduct)
	mux.HandleFunc("DELETE /api/v1/products/{id}", productHandler.DeleteProduct)

	mux.HandleFunc("POST /products", productHandler.CreateProduct)
	mux.HandleFunc("GET /products", productHandler.GetAllProducts)
	mux.HandleFunc("GET /products/{id}", productHandler.GetProduct)
	mux.HandleFunc("PUT /products/{id}", productHandler.UpdateProduct)
	mux.HandleFunc("DELETE /products/{id}", productHandler.DeleteProduct)
}