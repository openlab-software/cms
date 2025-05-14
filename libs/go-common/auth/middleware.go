package auth

import (
	"log"
	"net/http"

	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
)

func NewMiddleware(oidcProvider *oidc.Provider) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader("Authorization")

		if authorizationHeader == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := extractToken(authorizationHeader)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid auth header"})
			return
		}

		claims, err := checkToken(ctx.Request.Context(), oidcProvider, token)
		if err != nil {
			log.Fatal(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		ctx.Set("user", claims.Email)
		ctx.Next()
	}
}
