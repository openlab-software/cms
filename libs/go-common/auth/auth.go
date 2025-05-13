package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
)

type User struct {
	Email string `json:"email"`
}

func NewOIDCProvider() (*oidc.Provider, error) {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "")
	if err != nil {
		return nil, err
	}

	return provider, nil
}

func NewAuthMiddleware(oidcProvider *oidc.Provider) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader("Authorization")

		if authorizationHeader == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authorizationHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid auth header"})
			return
		}
		config := &oidc.Config{
			SkipClientIDCheck: true,
		}

		verifier := oidcProvider.Verifier(config)
		idToken, err := verifier.Verify(ctx.Request.Context(), parts[1])

		if err != nil {
			log.Fatal(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		claims := &User{}
		idToken.Claims(claims)
		jsonByte, _ := json.Marshal(claims)

		fmt.Println(string(jsonByte))
		ctx.Set("user", claims.Email)
		ctx.Next()
	}
}
