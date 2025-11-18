package main
import "demob/src/application"
import "demob/src/infraestructure"
import "demob/src/domain"

import "github.com/gin-gonic/gin"
import "net/http"
import "strconv"
import "os"

func main () {
	mysql := infraestructure.NewMysql()
	myGin := gin.Default()

	createProduct := application.NewCreateUseCase(mysql)
	getAll := application.NewViewAllUseCase(mysql)
	getById := application.NewViewByIdProductUseCase(mysql)

	myGin.GET("/products", func(c *gin.Context) {
		products, err := getAll.Run()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	})

	myGin.GET("/products/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		product, err := getById.Run(int32(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		c.JSON(http.StatusOK, product)
	})

	myGin.POST("/products", func(c *gin.Context) {
		var newProduct domain.Product
	
		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		if err := createProduct.Run(newProduct); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create product"})
			return
		}
	
		c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
	})
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	myGin.Run(":" + port)
}
