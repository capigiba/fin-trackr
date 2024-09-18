package middleware

import (
	"fintrack/internal/service/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthUserMiddleware handles user authentication
type AuthUserMiddleware struct {
	authService *auth.AuthService
}

// NewAuthUserMiddleware creates a new AuthUserMiddleware
func NewAuthUserMiddleware(authService *auth.AuthService) *AuthUserMiddleware {
	return &AuthUserMiddleware{
		authService: authService,
	}
}

// Auth is a middleware function that authenticates the user if a token is present
func (am *AuthUserMiddleware) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := extractToken(ctx)
		if len(token) == 0 {
			ctx.Next()
			return
		}

		// Get user info based on token
		userInfo, err := am.authService.GetUserByToken(token)
		if err != nil || userInfo == nil {
			ctx.Next()
			return
		}

		// Set the user info in the context
		ctx.Set("userInfo", userInfo)
		ctx.Next()
	}
}

// MustAuth ensures the user is authenticated; otherwise, returns an error
func (am *AuthUserMiddleware) MustAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := extractToken(ctx)
		if len(token) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Get user info based on token
		userInfo, err := am.authService.GetUserByToken(token)
		if err != nil || userInfo == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Example: Assuming status 1 means "active"
		if userInfo.Status != 1 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not active"})
			return
		}

		// Set the user info in the context
		ctx.Set("userInfo", userInfo)
		ctx.Next()
	}
}

// extractToken extracts the token from the Authorization header or query parameter
func extractToken(ctx *gin.Context) string {
	token := ctx.GetHeader("Authorization")
	if len(token) == 0 {
		token = ctx.Query("Authorization")
	}
	return strings.TrimSpace(strings.TrimPrefix(token, "Bearer "))
}
