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

// HandleLogin godoc
// @Summary logged in existing user
// @Schemes
// @Description login
// @Tags example
// @Accept json
// @Produce json
// @Param request body model.UserLoginRequest true "query params"
// @Success 200 {string} Helloworld
// @Router /login [post]
func (ac *AuthController) HandleLogin(c *gin.Context) {
	payload := model.UserLoginRequest{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}
	res, err := ac.AuthService.Login(payload)
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

// HandleGetAccountInfo godoc
// @Summary get accountt information from logged in user
// @Schemes
// @Description get account information from logged in user
// @Tags example
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {string} Helloworld
// @Router /account-info [get]
func (ac *AuthController) HandleGetAccountInfo(c *gin.Context) {
	userID := c.GetInt("user_id")
	user, err := ac.AuthService.GetAccountInfo(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"ok":   true,
		"data": user,
	})
}

// AuthMiddleware for protecting us
func (ac *AuthController) AuthMiddleware(c *gin.Context) {
	token := c.Request.Header["Authorization"]
	if len(token) < 1 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"ok":  false,
			"msg": "unauthorized",
		})
		return
	}
	userID, err := ac.AuthService.DecodeToken(token[0])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}
	c.Set("user_id", userID)
	c.Next()
}
