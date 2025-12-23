package middleware

import (
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/phanhotboy/nien-su-viet/apps/cms/global"
)

type User struct {
	ID    string `json:"id" doc:"The unique identifier of the user" example:"123e4567-e89b-12d3-a456-426614174000"`
	Email string `json:"email" doc:"The email address of the user" example:"example@gmail.com"`
	Role  string `json:"role" doc:"The role of the user" example:"admin"`
	Name  string `json:"name" doc:"The full name of the user" example:"John Doe"`
	Image string `json:"image" doc:"The profile image URL of the user" example:"https://example.com/profile.jpg"`
}

type JwtPayload struct {
	User     User        `json:"user" doc:"The user information extracted from the JWT token"`
	Version  string      `json:"version" doc:"The version of the JWT token" example:"1.0"`
	UpdateAt int         `json:"updateAt" doc:"The update timestamp of the JWT token" example:"1627847261"`
	Session  interface{} `json:"session" doc:"The session information associated with the JWT token"`
}

func Authentication(api huma.API) func(huma.Context, func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		config := global.Config.Security
		cookieStr := ctx.Header("Cookie")
		cookies := parseCookies(cookieStr)

		token := cookies[config.AuthCookiePrefix+".session_data"]

		if token == "" {
			huma.WriteErr(api, ctx, http.StatusUnauthorized,
				"Please login to access this resource",
			)
			return
		}

		secret := config.BetterAuthSecret
		if secret == "" {
			huma.WriteErr(api, ctx, http.StatusBadRequest, "JWT secret not configured")
		}

		// Parse and verify JWT token
		parsedToken, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Verify signing method is HS256
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, fmt.Sprintf("Invalid token: %v", err))
			return
		}

		if !parsedToken.Valid {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid token")
			return
		}

		// Extract claims
		claims, ok := parsedToken.Claims.(*jwt.MapClaims)
		if !ok {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid token claims")
			return
		}

		// Attach user info and session to context
		ctx = huma.WithValue(ctx, "user", (*claims)["user"])
		ctx = huma.WithValue(ctx, "session", (*claims)["session"])

		next(ctx)
	}
}

func parseCookies(cookieHeader string) map[string]string {
	res := make(map[string]string)
	header := http.Header{}
	header.Set("Cookie", cookieHeader)
	request := http.Request{Header: header}
	cookies := request.Cookies()
	for _, cookie := range cookies {
		res[cookie.Name] = cookie.Value
	}
	return res
}
