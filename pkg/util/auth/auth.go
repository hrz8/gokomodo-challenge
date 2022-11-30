package auth

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	res "github.com/hrz8/gokomodo-challenge/pkg/util/response"
	"github.com/labstack/echo/v4"
)

type AuthPayload struct {
	Sub string
	Aud string
}

func VerifyJWT(c *echo.Context) (*AuthPayload, error) {
	ctx := *c

	authorization := ctx.Request().Header["Authorization"]
	jwtKey := os.Getenv("JWT_SECRET")
	if jwtKey == "" {
		jwtKey = "secret"
	}

	if len(authorization) == 0 {
		return nil, res.ErrorBuilder(
			&res.ErrorConstant.Unauthenticated,
			errors.New("you are forbidden"),
		)
	}

	bearer := strings.Split(authorization[0], "Bearer ")
	if len(bearer) < 2 {
		return nil, res.ErrorBuilder(
			&res.ErrorConstant.Unauthenticated,
			errors.New("you are forbidden"),
		)
	}

	token := bearer[1]
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
		}

		return []byte(jwtKey), nil
	})

	if !jwtToken.Valid || err != nil {
		return nil, res.ErrorBuilder(
			&res.ErrorConstant.Unauthorized,
			err,
		)
	}

	tokenMap := jwtToken.Claims.(jwt.MapClaims)

	var sub string
	destructedSub := tokenMap["sub"]
	if destructedSub != nil {
		sub = destructedSub.(string)
	}

	var aud string
	destructedAud := tokenMap["aud"]
	if destructedAud != nil {
		aud = destructedAud.(string)
	}

	return &AuthPayload{sub, aud}, nil
}
