package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/sangianpatrick/devoria-article-service/domain/account/entity"
	"github.com/sangianpatrick/devoria-article-service/jwt"
)

// BearerAuth is a concrete struct of bearer auth verifier.
type BearerAuth struct {
	jsonWebToken jwt.JSONWebToken
}

// NewBearer is a constructor.
func NewBearerAuth(jsonWebToken jwt.JSONWebToken) RouteMiddlewareBearer {
	return &BearerAuth{jsonWebToken}
}

// Verify will verify the request to ensure it comes with an authorized bearer auth token.
func (b *BearerAuth) VerifyBearer(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		auth := r.Header.Get("Authorization")
		token := strings.Split(auth, " ")[1]
		res, err := b.jsonWebToken.Parse(ctx, token, &entity.AccountStandardJWTClaims{})
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		claims := (res.Claims).(*entity.AccountStandardJWTClaims)
		ctx = context.WithValue(ctx, entity.EmailCtx, claims.Email)

		next(w, r.WithContext(ctx))
	})
}
