package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
)

func (middlwares *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Please add authorization header", http.StatusUnauthorized)
			return
		}

		authHeaderArr := strings.Split(authHeader, " ")
		if len(authHeaderArr) != 2 {
			http.Error(w, "Unauthorize, authHeader", http.StatusUnauthorized)
			return
		}

		accessToken := authHeaderArr[1]

		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorize, token parts", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		signature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload
		byteArrMessage := []byte(message)

		byteArrSecret := []byte(middlwares.cnf.JWTSecret)
		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)

		hash := h.Sum(nil)
		newSignature := base64UrlEncode(hash)

		if signature != newSignature {
			http.Error(w, "Unauthorize, Tui Hacker", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
