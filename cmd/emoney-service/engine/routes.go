package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/ridhozhr10/ottojunior/internal/controller"
	"github.com/ridhozhr10/ottojunior/internal/repository/psql"
	"github.com/ridhozhr10/ottojunior/internal/service/auth"
	"github.com/ridhozhr10/ottojunior/pkg/database"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Config structure
type Config struct {
	DBHost string
	DBUser string
	DBPass string
	DBPort string
	DBName string
	Port   string
}

// New http api engine
func New(config Config) error {
	r := gin.Default()
	// engine
	dbConfig := database.Config{
		DBHost: config.DBHost,
		DBUser: config.DBUser,
		DBPass: config.DBPass,
		DBPort: config.DBPort,
		DBName: config.DBName,
	}
	db, err := database.ConnectGORM(dbConfig)
	if err != nil {
		return err
	}
	// Repository
	userRepository := psql.NewUserPsqlRepository(db)

	// Service
	authService := auth.NewService(userRepository)

	// Controller
	authController := controller.NewAuthController(authService)

	// Route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.POST("/register", authController.HandleRegister)

	return r.Run(config.Port)
}