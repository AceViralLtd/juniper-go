package juniper

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var ErrInvalidJWT = errors.New("Invalid jwt")

// GenerateJWT will generate a new JWT for the given account model
func GenerateJWT(claims jwt.MapClaims, appSecret string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(appSecret))
}

// ExtractJwtFromAuthHeader will verify that the Authorization header both exists and is in the
// Bearer format, if so it will extract the token (hopefully this should be a valid JWT)
func ExtractJwtFromAuthHeader(ctx echo.Context) string {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		return ""
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 {
		return ""
	}

	return parts[1]
}

// VerifyJwt will check that the JWT is both a jwt and valid
func VerifyJwt(jwtString, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != "HS256" {
			return nil, ErrInvalidJWT
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return nil, ErrInvalidJWT
	} else {
		return claims, nil
	}
}
