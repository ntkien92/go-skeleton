package cmd

import (
	"blog-api/errors"
	"blog-api/response"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (server *ApiServer) setMiddleware() {
	server.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodDelete,
		},
	}))

	server.echo.Use(middleware.Logger())
	server.echo.Use(middleware.Recover())
	server.echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			os.Setenv("PROCESS_ID", uuid.New().String())
			return next(c)
		}
	})

	server.echo.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if httpErr, ok := err.(*echo.HTTPError); ok {
			code = httpErr.Code
		}
		resp := response.NewApiResponse(c.Path())
		resp.Status.Code = code
		resp.Status.Type = http.StatusText(code)

		c.JSON(code, resp)
	}
}

func (server *ApiServer) VerifyJwt() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			resp := response.NewApiResponse(c.Path())
			resp.Status.Code = http.StatusUnauthorized
			resp.Status.Type = http.StatusText(http.StatusUnauthorized)

			auth_header := c.Request().Header.Get("Authorization")
			if auth_header == "" {
				resp.Errors = append(resp.Errors, errors.New(
					-1,
					"Authorization header is empty",
					nil,
				))
				return c.JSON(http.StatusUnauthorized, resp)
			}

			jwt := ""
			if strings.HasPrefix(auth_header, "Bearer ") {
				jwt = auth_header[len("Bearer "):]
			}

			if jwt == "" {
				resp.Errors = append(resp.Errors, errors.New(
					-1,
					"JWT is empty",
					nil,
				))
				return c.JSON(http.StatusUnauthorized, resp)
			}

			err := verifyJWT(jwt)
			if err != nil {
				resp.Errors = append(resp.Errors, errors.New(
					-1,
					err.Error(),
					nil,
				))
				return c.JSON(http.StatusUnauthorized, resp)
			}

			return next(c)
		}
	}
}

var hmacSampleSecret = []byte("")

func verifyJWT(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		return fmt.Errorf("Invalid Token")
	}

	return nil
}
