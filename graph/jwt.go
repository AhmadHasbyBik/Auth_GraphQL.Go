package graph

import "github.com/dgrijalva/jwt-go"

var mySigningKey = []byte("graph.io")

func JwtDecode(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
}
func JwtCreate(userID int, expiredAt int64) string {
	claims := UserClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: expiredAt,
			Issuer:    "graph",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, _ := token.SignedString(mySigningKey)
	return jwt
}
