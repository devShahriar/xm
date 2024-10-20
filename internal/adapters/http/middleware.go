package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devShahriar/xm/internal/common"
	"github.com/devShahriar/xm/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
)

type MiddlewareNamedHandler struct {
	Perm string
	Func echo.HandlerFunc
}

func Middleware(permission string, handler echo.HandlerFunc) MiddlewareNamedHandler {
	return MiddlewareNamedHandler{
		Perm: permission,
		Func: handler,
	}
}

func (s *Server) Auth(next MiddlewareNamedHandler) echo.HandlerFunc {

	return func(c echo.Context) error {
		token := ParseToken(c.Request(), "Authorization")
		jwtPayload, err := ParseJWTToken(token, []byte(config.GetAppConfig().JWTSecretKey))
		if err != nil || jwtPayload == nil {
			log.Printf("failed to parse jwt token %v", err)
			return c.JSON(http.StatusForbidden, common.ErrorMsg{Message: "access denied"})
		}

		if !ValidatePermission(jwtPayload, next.Perm) {
			log.Println("invalid permission")
			return c.JSON(http.StatusForbidden, common.ErrorMsg{Message: "access denied"})
		}
		return next.Func(c)
	}

}

func ValidatePermission(jwtPayload *common.JwtPayload, permission string) bool {
	return jwtPayload.PermissionMap[permission]
}

func ParseToken(r *http.Request, headerKey string) string {

	token := r.Header.Get(headerKey)
	return token
}

func ParseJWTToken(tokenString string, secretKey []byte) (*common.JwtPayload, error) {
	claims := &common.JwtPayload{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*common.JwtPayload); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
