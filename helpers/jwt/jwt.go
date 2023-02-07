package jwt

import (
	jwt "github.com/golang-jwt/jwt"
	"github.com/segmentio/ksuid"
)

// jwtClaims ...
type jwtClaims struct {
	jwt.StandardClaims
}

// GetToken ...
func GetToken(secret string) (string, error) {
	claims := &jwtClaims{
		jwt.StandardClaims{
			Id: ksuid.New().String(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte(secret))

	return token, err
}
