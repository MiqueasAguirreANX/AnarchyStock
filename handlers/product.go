package handlers

import (
	"AnarchyStock/database"
	"AnarchyStock/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context) {
	var prodSer models.ProductSerializer
	if err := ctx.BindJSON(&prodSer); err != nil {
		ErrorMessage(ctx, err.Error())
		return
	}
	models.CreateProduct(prodSer)
	SuccessMessage(ctx, "Product Created")
}

func SearchProduct(ctx *gin.Context) {
	type SearchInput struct {
		Type  string
		Value string
	}
	var search SearchInput
	if err := ctx.BindJSON(&search); err != nil {
		ErrorMessage(ctx, err.Error())
		return
	}
	fmt.Println(search)
	var products []models.Product
	switch search.Type {
	case "id":
		pkUint, err := strconv.ParseUint(search.Value, 10, 32)
		if err != nil {
			ErrorMessage(ctx, err.Error())
			return
		}
		prod := models.GetProductByID(uint(pkUint))
		products = append(products, prod)
	}
	DataMessage(ctx, products)
}

func GetAllProducts(ctx *gin.Context) {
	products := models.GetAllProducts()
	DataMessage(ctx, products)
}

func GetProductByID(ctx *gin.Context) {
	pk := ctx.Param("pk")
	pkUint, err := strconv.ParseUint(pk, 10, 32)
	if err != nil {
		ErrorMessage(ctx, err.Error())
		return
	}
	product := models.GetProductByID(uint(pkUint))
	DataMessage(ctx, product)
}

func GetProductsPaginated(ctx *gin.Context) {
	page := ctx.Param("page")
	pageUint64, err := strconv.ParseUint(page, 10, 64)
	if err != nil {
		ErrorMessage(ctx, err.Error())
		return
	}
	pageUint := uint(pageUint64)
	if pageUint < 1 {
		pageUint = 1
	}

	var products []map[string]interface{}
	products, count, pages := database.Paginator(&models.Product{}, pageUint)
	DataMessage(ctx, gin.H{
		"products":     products,
		"page":         pageUint,
		"previousPage": pageUint - 1,
		"nextPage":     pageUint + 1,
		"count":        count,
		"pages":        pages,
	})
}
