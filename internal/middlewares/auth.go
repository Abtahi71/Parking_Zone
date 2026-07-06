package middleware

import (
	"gotickets/internal/auth"
	"net/http"
	"strings"
	"github.com/labstack/echo/v5"
)

func AuthMiddleware(jwtService auth.JWTService)echo.MiddlewareFunc{
     return func(next echo.HandlerFunc) echo.HandlerFunc{
		return func(c *echo.Context)error{
			authHeader:=c.Request().Header.Get("Authorization")
			if authHeader==""{
                return c.JSON(http.StatusUnauthorized,map[string]string{"error":"You do not have access"})
			}
			parts:=strings.Split(authHeader, " ")
			if len(parts)!=2 || parts[0]!="Bearer"{
				return c.JSON(http.StatusUnauthorized,map[string]string{"error":"Invalid authorization header format"})
			}
			token:=parts[1]

			claims,err:=jwtService.ValidateToken(token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "invalid or expired token",
				})
			}

			c.Set("userId",claims.UserId)
			c.Set("name",claims.Name)
			c.Set("email",claims.Email)
			c.Set("role",claims.Role)

			return next(c)
			
		}
	 }
}

func RequireRole(UserRole string)echo.MiddlewareFunc{
	return  func(next echo.HandlerFunc)echo.HandlerFunc{
		return func(c *echo.Context)error{
			role,ok:=c.Get("role").(string)
			if !ok{
				return c.JSON(http.StatusUnauthorized,map[string]string{"error":"You do not have access"})
			}
			if role!=UserRole{
				return c.JSON(http.StatusForbidden,map[string]string{"error":"You do not have access"})
			}
			return next(c)
		}
	}
}
