package middlewares

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/golang-basic-auth-api/configs"
	"github.com/sherwin-77/golang-basic-auth-api/db"
)

func ValidateUUID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if _, err := uuid.Parse(id); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid UUID")
		}
		return next(c)
	}
}

func Authenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		config := configs.GetConfig()

		// Extract the "Authorization" header.
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid token")
		}

		// Split the token from the header.
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized, "Malformed token")
		}

		tokenString := strings.TrimSpace(splitToken[1])

		// Parse the JWT token.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte(config.Key), nil
		})

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		c.Set("user_id", claims["user_id"])
		return next(c)
	}
}

func AuthLevel(level int) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			DB := db.DB
			if DB == nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Database connection not available")
			}

			userID := c.Get("user_id").(string)
			var userLevel int

			DB.Table("user_roles").
				Where("user_id = ?", userID).
				Joins("JOIN roles ON user_roles.role_id = roles.id").
				Select("MAX(roles.auth_level)").
				Scan(&userLevel)

			if userLevel < level {
				return echo.NewHTTPError(http.StatusForbidden, "Insufficient permission")
			}
			return next(c)
		}
	}
}
