package router

import (
	"fintrack/internal/handler/rest/v1/transaction"
	"fintrack/internal/handler/rest/v1/user"
	"fintrack/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// Approuter holds all controllers for routing
type AppRouter struct {
	userHandler        *user.UserHandler
	transactionHandler *transaction.TransactionHandler
	authMiddleware     *middleware.AuthUserMiddleware
	swaggerRouter      *SwaggerRouter
}

// NewAppRouter creates a new AppRouter instance
func NewAppRouter(userHandler *user.UserHandler, transactionHandler *transaction.TransactionHandler, authMiddleware *middleware.AuthUserMiddleware, swaggerRouter *SwaggerRouter) *AppRouter {
	return &AppRouter{
		userHandler:        userHandler,
		transactionHandler: transactionHandler,
		authMiddleware:     authMiddleware,
		swaggerRouter:      swaggerRouter,
	}
}

// RegisterUserRoutes sets up the routes for user-related operation
func (a *AppRouter) RegisterUserRoutes(r *gin.RouterGroup) {
	public := r.Group("/users")
	{
		public.POST("/register", a.userHandler.RegisterUser)
		public.POST("/login", a.userHandler.Login)
	}

	protected := r.Group("/users")
	protected.Use(a.authMiddleware.MustAuth())
	{
		protected.GET("/:id", a.userHandler.GetUser)
		protected.PUT("/:id", a.userHandler.UpdatedUser)
		protected.DELETE("/:id", a.userHandler.DeleteUser)
	}
}

// RegisterTransactionRoutes sets up the routes for transaction-related operation
func (a *AppRouter) RegisterTransactionRoutes(r *gin.RouterGroup) {
	protected := r.Group("/transaction")
	protected.Use(a.authMiddleware.MustAuth())
	{
		protected.POST("/", a.transactionHandler.PostTransaction)
		protected.GET("/:year/:month", a.transactionHandler.GetTransactions)
	}
}

// RegisterSwaggerRoutes sets up the route for Swagger API documentation
func (a *AppRouter) RegisterSwaggerRoutes(r *gin.RouterGroup) {
	// Check if SwaggerRouter is initialized before registering
	if a.swaggerRouter != nil {
		a.swaggerRouter.Register(r)
	}
}
