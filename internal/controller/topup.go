package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridhozhr10/ottojunior/internal/model"
	"github.com/ridhozhr10/ottojunior/internal/service/topup"
)

// TopupController gin wrapper
type TopupController struct {
	TopupService topup.Service
}

// NewTopupController helper function
func NewTopupController(TopupService topup.Service) *TopupController {
	return &TopupController{TopupService}
}

// HandleTopupBalance godoc
// @Summary handle new topup balance
// @Schemes
// @Description handle new topup balance
// @Tags topup-service
// @Accept json
// @Produce json
// @Param request body model.TopupBalanceRequest true "query params"
// @Success 200 {string} Helloworld
// @Router /topup [post]
func (pc *TopupController) HandleTopupBalance(c *gin.Context) {
	payload := model.TopupBalanceRequest{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}
	response, err := pc.TopupService.TopupBalance(payload)
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
