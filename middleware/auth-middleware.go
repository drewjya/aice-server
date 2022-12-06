package middleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/drewjya/aice-server/entity"
	"github.com/drewjya/aice-server/helper"
	"github.com/drewjya/aice-server/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authorization(jwtService services.JWTService) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		log.Println("========================================")
		log.Println(ctx.Get("Token"))
		log.Println("========================================")
		if authHeader == "" {

			response := helper.BuildError(http.StatusBadRequest, "No Token Found", errors.New("failed to process request").Error())
			ctx.AbortWithStatusJSON(int(response.Status), response)
			return
		}
		tken, err := jwtService.ValidateToken(authHeader)
		if err != nil {

			response := helper.BuildError(http.StatusUnauthorized, "Token is not valid", err.Error())
			ctx.AbortWithStatusJSON(int(response.Status), response)
			return
		}
		if tken.Valid {
			claims := tken.Claims.(jwt.MapClaims)
			ctx.Set("Token", entity.TokenValue{
				UserId: claims["userId"].(string),
				Name:   claims["name"].(string),
			})
			log.Println("========================================")
			log.Println(ctx.Get("Token"))
			log.Println("========================================")

		} else {

			response := helper.BuildError(http.StatusUnauthorized, "Token is not valid", err.Error())
			ctx.AbortWithStatusJSON(int(response.Status), response)
			return
		}
	}
}
