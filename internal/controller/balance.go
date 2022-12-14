package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridhozhr10/ottojunior/internal/service/balance"
)

// BalanceController gin wrapper
type BalanceController struct {
	BalanceService balance.Service
}

// NewBalanceController helper function
func NewBalanceController(BalanceService balance.Service) *BalanceController {
	return &BalanceController{BalanceService}
}

// HandleGetBalance godoc
// @Summary get latest balance from logged in user
// @Schemes
// @Description get latest balance from logged in user
// @Tags emoney-service
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {string} Helloworld
// @Router /balance [get]
func (bc *BalanceController) HandleGetBalance(c *gin.Context) {
	userID := c.GetInt("user_id")
	balance, err := bc.BalanceService.GetBalance(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"ok":   true,
		"data": balance,
	})
}
