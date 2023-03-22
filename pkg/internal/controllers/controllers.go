package controllers

import (
	"internal/schemas"

	"github.com/gorilla/securecookie"
)

var (
	sc         *securecookie.SecureCookie
	cookieName string
)

func Init(config *schemas.Controller) {
	secret := []byte(config.SecretKey)
	block := []byte(config.BlockKey)
	cookieName = config.CookieName

	sc = securecookie.New(secret, block)
}
