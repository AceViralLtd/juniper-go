package juniper

import "github.com/golang-jwt/jwt"

// GenerateJWT will generate a new JWT for the given account model
func GenerateJWT(claims jwt.MapClaims, appSecret string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(appSecret))
}
