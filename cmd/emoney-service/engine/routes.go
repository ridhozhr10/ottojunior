package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/ridhozhr10/ottojunior/internal/controller"
	"github.com/ridhozhr10/ottojunior/internal/repository/httpapi"
	"github.com/ridhozhr10/ottojunior/internal/repository/psql"
	"github.com/ridhozhr10/ottojunior/internal/service/auth"
	"github.com/ridhozhr10/ottojunior/internal/service/balance"
	"github.com/ridhozhr10/ottojunior/internal/service/product"
	"github.com/ridhozhr10/ottojunior/internal/service/transaction"
	"github.com/ridhozhr10/ottojunior/pkg/database"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Config structure
type Config struct {
	DBHost               string
	DBUser               string
	DBPass               string
	DBPort               string
	DBName               string
	Port                 string
	BillerServiceBaseURL string
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
	balanceRepository := psql.NewBalancePsqlRepository(db)
	transactionRepository := psql.NewTransactionPsqlRepository(db)
	productRepository := httpapi.NewProductHttpapiRepository(config.BillerServiceBaseURL)

	// Service
	authService := auth.NewService(userRepository, balanceRepository)
	balanceService := balance.NewService(balanceRepository)
	productService := product.NewService(productRepository)
	transactionService := transaction.NewService(
		balanceRepository,
		transactionRepository,
		userRepository,
		productRepository,
	)

	// Controller
	authController := controller.NewAuthController(authService)
	balanceController := controller.NewBalanceController(balanceService)
	productController := controller.NewProductController(productService)
	transactionController := controller.NewTransactionController(transactionService)

	// Route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) // docs
	r.POST("/register", authController.HandleRegister)
	r.POST("/login", authController.HandleLogin)

	protectedRoute := r.Group("", authController.AuthMiddleware)
	protectedRoute.GET("/account-info", authController.HandleGetAccountInfo)
	protectedRoute.GET("/balance", balanceController.HandleGetBalance)
	protectedRoute.GET("/product", productController.HandleGetProduct)
	protectedRoute.GET("/transaction", transactionController.HandleGetTransaction)
	protectedRoute.POST("/confirm-transaction", transactionController.HandleConfirmTransaction)

	return r.Run(config.Port)
}
