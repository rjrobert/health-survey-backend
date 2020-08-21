package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rjrobert/health-survey-backend/cmd/health-survey-backend/config"
	"github.com/rjrobert/health-survey-backend/cmd/health-survey-backend/httputil"
)

// APIKeyAuth validates request with api_key
func APIKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			httputil.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			c.Abort()
		} else if authHeader != config.Config.APIKey {
			httputil.NewError(c, http.StatusUnauthorized, fmt.Errorf("Invalid API key: api_key=%s", authHeader))
			c.Abort()
		}
		c.Next()
	}
}
