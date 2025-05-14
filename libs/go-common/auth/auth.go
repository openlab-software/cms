package auth

import (
	"context"

	"github.com/coreos/go-oidc"
)

const (
	issuer   = ""
	cliendID = ""
)

func NewOIDCProvider() (*oidc.Provider, error) {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		return nil, err
	}

	return provider, nil
}
