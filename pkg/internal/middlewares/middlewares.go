package middlewares

import "internal/schemas"

var middleware *schemas.Middleware

func Init(mid *schemas.Middleware) {
	middleware = mid
}
