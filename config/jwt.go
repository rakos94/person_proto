package config

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"

	"person_proto/helper"
	pb "person_proto/models"

	"github.com/labstack/echo"
)

func SetJwt(e *echo.Echo) *echo.Group {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	jwtGroup := e.Group("/api")

	jwtGroup.Use(middlewareCredential)

	return jwtGroup
}

func CreateJwtToken(email string) (string, error) {
	// Create token
	var signingMethod *jwt.SigningMethodHMAC
	if SigningMethodExample == "HS512" {
		signingMethod = jwt.SigningMethodHS512
	}
	token := jwt.New(signingMethod)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Duration(JwtExpiredHour) * time.Hour).Unix()

	t, err := token.SignedString([]byte(SecretKeyExample))
	if err != nil {
		return "", err
	}

	return t, nil
}

func middlewareCredential(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqToken := c.Request().Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) == 1 {
			return c.JSON(http.StatusUnauthorized, "No token")
		}
		reqToken = splitToken[1]
		err := CheckCredentialToken(reqToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		return next(c)
	}
}

// CheckCredentialToken ...
func CheckCredentialToken(token string) error {
	res, err := Client.ValidateToken(Ctx,
		&pb.Token{Data: token})

	if err != nil {
		desc := strings.Split(err.Error(), "desc = ")
		err = errors.New(desc[1])
		log.Println("Error validate =>", helper.RPCErrDesc(err))
		return err
	}

	log.Println("Success validate =>", res)
	return nil
}
