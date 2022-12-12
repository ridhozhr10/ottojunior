package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridhozhr10/ottojunior/internal/model"
	"github.com/ridhozhr10/ottojunior/internal/service/auth"
)

// AuthController gin http handler wrapper
type AuthController struct {
	AuthService auth.Service
}

// NewAuthController helper function
func NewAuthController(
	AuthService auth.Service,
) *AuthController {
	return &AuthController{AuthService}
}

// HandleRegister godoc
// @Summary register new user
// @Schemes
// @Description register new user
// @Tags example
// @Accept json
// @Produce json
// @Param request body model.UserRegisterRequest true "query params"
// @Success 200 {string} Helloworld
// @Router /register [post]
// HandleRegister request
func (ac *AuthController) HandleRegister(c *gin.Context) {
	user := model.UserRegisterRequest{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}
	newUser := model.User{
		Name:        user.Name,
		Username:    user.Username,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
	}
	res, err := ac.AuthService.Register(newUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": res,
	})
}
