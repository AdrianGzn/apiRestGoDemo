package main
import "demob/src/application"
import "demob/src/infraestructure"
import "demob/src/domain"

import "github.com/gin-gonic/gin"
import "net/http"
import "strconv"

func main () {
	mysql := infraestructure.NewMysql()
	myGin := gin.Default()

	createProduct := application.NewCreateUseCase(mysql)
	getAll := application.NewViewAllUseCase(mysql)
	getById := application.NewViewByIdProductUseCase(mysql)
	updateProduct := application.NewUpdateProductUseCase(mysql)
	deleteProduct := application.NewDeleteProductUseCase(mysql)

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
	

	myGin.PUT("/products/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var updatedProduct domain.Product
		if err := c.ShouldBindJSON(&updatedProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := updateProduct.Run(int32(id), updatedProduct); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
	})

	myGin.DELETE("/products/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := deleteProduct.Run(int32(id)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
	})


	myGin.Run()

	/*
	Esto es para crear un producto
	product1 := domain.NewProduct("PauPau", 15.50)
	createProduct.Run(*product1)

	esto para obtener todos los productos
	products, _ := getAll.Run()
	fmt.Println("Products:", products)

	Esto para obtener por id
	product, err := getById.Run(1)
	if err != nil {
		fmt.Println("Ha ocurrido un error")
	} else {
		fmt.Println("Product: ", product)
	}

	para actualizar un producto
	updatedProduct := domain.NewProduct("Gaming Laptop", 2000.75)
	updateProduct.Run(1, *updatedProduct)

	Obtener todos los productos otra vez
	products, _ = getAll.Run()
	fmt.Println("Products: ", products)

	Eliminar un producto por id
	deleteProduct.Run(1)

	Y obtener nuevamente
	products, _ = getAll.Run()
	fmt.Println("Products: ", products)
	*/
}
