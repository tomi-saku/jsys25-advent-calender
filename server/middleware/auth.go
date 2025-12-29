package middlewares

import (
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tomi-saku/jsys25-advent-calender/models"
	"google.golang.org/api/idtoken"
)

func verifyTokenAndGetPayload(c *gin.Context, clientID string) *idtoken.Payload {
	const bearerPrefix = "Bearer "
	AuthHeader := c.GetHeader("Authorization")
	slog.Debug("Auth Middleware", "AuthHeader", AuthHeader)
	if AuthHeader == "" {
		errRes := models.Error{
			Message: "Authorization header is required",
		}
		c.AbortWithStatusJSON(401, errRes)
		return nil
	}
	if !strings.HasPrefix(AuthHeader, bearerPrefix) {
		errRes := models.Error{
			Message: "Authorization header must be Bearer token",
		}
		c.AbortWithStatusJSON(401, errRes)
		return nil
	}
	token := strings.TrimPrefix(AuthHeader, bearerPrefix)
	slog.Debug("Auth Utility", "Token", token)
	payload, err := idtoken.Validate(c.Request.Context(), token, clientID)
	if err != nil {
		slog.Debug("failed to Authorize", "error", err)
		errRes := models.Error{
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(401, errRes)
		return nil
	}
	return payload
}

func AuthMiddleware(clientID string) gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := verifyTokenAndGetPayload(c, clientID)
		if payload == nil {
			return
		}
		c.Set("userID", payload.Subject)
		slog.Debug("Auth Middleware", "userID", payload.Subject)
		c.Next()
	}
}
