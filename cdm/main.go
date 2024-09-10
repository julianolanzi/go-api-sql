package main

import (
	"go-api-sql/controller"
	"go-api-sql/db"
	"go-api-sql/repository" // Add the missing import statement
	"go-api-sql/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	// Camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	// Camada de usecase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	// Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)

	server.Run(":1111")

}
