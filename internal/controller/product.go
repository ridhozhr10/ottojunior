package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridhozhr10/ottojunior/internal/service/product"
)

// ProductController gin wrapper
type ProductController struct {
	ProductService product.Service
}

// NewProductController helper function
func NewProductController(ProductService product.Service) *ProductController {
	return &ProductController{ProductService}
}

// HandleGetProduct godoc
// @Summary get all available product to buy
// @Schemes
// @Description get all available product to buy
// @Tags example
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {string} Helloworld
// @Router /product [get]
func (pc *ProductController) HandleGetProduct(c *gin.Context) {
	product, err := pc.ProductService.GetProductList()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"ok":   true,
		"data": product,
	})
}
