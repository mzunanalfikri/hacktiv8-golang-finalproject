package controller

import (
	"net/http"
	"project-4/model"
	"project-4/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var (
		product model.Product
	)

	err := c.ShouldBind(&product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	_, err = service.GetCategoryDetail(product.CategoryID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Category ID Not Found")
		return
	}

	result, err := service.CreateProduct(product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetAllProducts(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	products, err := service.GetAllProducts()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}

func UpdateProduct(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var product model.Product

	err := c.ShouldBind(&product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	id, _ := strconv.Atoi(c.Param("productId"))

	newProduct, err := service.GetProductById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	newProduct.Title = product.Title
	newProduct.Price = product.Price
	newProduct.Stock = product.Stock
	newProduct.CategoryID = product.CategoryID

	result, err := service.UpdateProduct(*newProduct)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

func DeleteProducts(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	id, _ := strconv.Atoi(c.Param("productId"))

	err := service.DeleteProduct(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Product has been successfully deleted",
	})
}
