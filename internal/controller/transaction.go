package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridhozhr10/ottojunior/internal/model"
	"github.com/ridhozhr10/ottojunior/internal/service/transaction"
)

// TransactionController gin wrapper
type TransactionController struct {
	TransactionService transaction.Service
}

// NewTransactionController helper function
func NewTransactionController(
	TransactionService transaction.Service,
) *TransactionController {
	return &TransactionController{TransactionService}
}

// HandleGetTransaction godoc
// @Summary handle get request transaction
// @Schemes
// @Description handle get request transaction
// @Tags emoney-service
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {string} Helloworld
// @Router /transaction [get]
func (tc *TransactionController) HandleGetTransaction(c *gin.Context) {
	userID := c.GetInt("user_id")
	response, err := tc.TransactionService.GetListTransaction(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"ok":   true,
		"data": response,
	})
}

// HandleConfirmTransaction godoc
// @Summary handle get request transaction
// @Schemes
// @Description handle get request transaction
// @Tags emoney-service
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {string} Helloworld
// @Param request body []model.ConfirmTransactionRequest true "query params"
// @Router /confirm-transaction [post]
func (tc *TransactionController) HandleConfirmTransaction(c *gin.Context) {
	payload := []model.ConfirmTransactionRequest{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}
	userID := c.GetInt("user_id")
	response, err := tc.TransactionService.ConfirmTransaction(userID, payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"ok":   true,
		"data": response,
	})
}
