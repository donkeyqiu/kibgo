package middleware

import (
	"net/http"
	"fmt"
	"strings"
			. "../utils"

	"github.com/dgrijalva/jwt-go"
)

func parseToken(tokenString string, key string) (interface{}, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println(err)
		return "", false
	}
}

func Auth(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if name != "Login" {
			var authString = ""
			if len(r.Header["Authorization"]) > 0 {
				authString = r.Header["Authorization"][0]
				kv := strings.Split(authString, " ")
				if len(kv) != 2 || kv[0] != "Bearer" {
					RespondWithError(w, http.StatusBadRequest, "Token错误")
				}
				tokenString := kv[1]
				_, ok := parseToken(tokenString, "Welcome to PurePearl")
				if ok {
					inner.ServeHTTP(w, r)
				} else {
					RespondWithError(w, http.StatusUnauthorized, "Token错误或过期")
				}
			} else {
				RespondWithError(w, http.StatusUnauthorized, "您没有授权")
			}
		} else {
			inner.ServeHTTP(w, r)
		}
	})
}
