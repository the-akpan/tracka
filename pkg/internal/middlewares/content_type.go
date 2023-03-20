package middlewares

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func AcceptedContentTypeMiddleware(next http.Handler) http.Handler {
	return handlers.ContentTypeHandler(next, middleware.ContentTypes...)
}

func ResponseContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(res, req)
	})
}
