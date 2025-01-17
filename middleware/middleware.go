package middleware

import "net/http"

// RouteMiddleware is a abstraction of route middleware.
type RouteMiddleware interface {
	Verify(next http.HandlerFunc) http.HandlerFunc
}

type RouteMiddlewareBearer interface {
	VerifyBearer(next http.HandlerFunc) http.HandlerFunc
}