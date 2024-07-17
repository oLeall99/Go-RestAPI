package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
) //Import

func main() {
	// Inicia o Servidor
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada Repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	// Camada UseCase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	// Camada de Controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Rota de buscar produtos
	server.GET("/products", ProductController.GetProducts)

	// Rota de Criar Produto
	server.POST("/product", ProductController.CreateProducts)

	// Rota de Buscar produto por Id
	server.GET("/product/:productId", ProductController.GetProductById)

	server.Run(":8000")
}
