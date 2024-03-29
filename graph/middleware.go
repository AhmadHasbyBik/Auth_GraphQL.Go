package graph

import (
	"context"
	"net"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type UserAuth struct {
	UserID    int
	Roles     []string
	IPAddress string
	Token     string
}

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}
type UserClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := TokenFromHttpRequest(r)

			userId := UserIDFromToken(token)

			ip, _, _ := net.SplitHostPort(r.RemoteAddr)
			userAuth := UserAuth{
				UserID:    userId,
				IPAddress: ip,
				Token:     token,
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &userAuth)
			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
func TokenFromHttpRequest(r *http.Request) string {
	reqToken := r.Header.Get("Authorization")
	var tokenString string
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) > 1 {
		tokenString = splitToken[1]
	}
	return tokenString
}
func UserIDFromToken(tokenString string) int {

	token, err := JwtDecode(tokenString)
	if err != nil {
		return 0
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		if claims == nil {
			return 0
		}
		return claims.UserId
	} else {
		return 0
	}
}
func ForContext(ctx context.Context) *UserAuth {
	raw := ctx.Value(userCtxKey)
	if raw == nil {
		return nil
	}
	return raw.(*UserAuth)
}
func GetAuthFromContext(ctx context.Context) *UserAuth {
	return ForContext(ctx)
}
