package controller

import (
	"net/http"
	"strconv"

	"aice-server/dto"
	"aice-server/entity"
	"aice-server/helper"
	"aice-server/services"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtService  services.JWTService
}

func NewAuthController(authService services.AuthService,
	jwtService services.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO = dto.LoginDTO{}
	err := ctx.ShouldBind(&loginDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildError(http.StatusBadRequest, "Failed Login", err.Error()))
		return
	}

	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)

	if v, ok := authResult.(entity.User); ok {

		generateToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10), v.Name)
		v.Token = generateToken
		response := helper.BuildResponse(http.StatusOK, "Success", v)
		ctx.JSON(http.StatusOK, response)
		return
	}

	ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildError(http.StatusUnauthorized, "Check Your Email/Password", "Invalid Credential"))
}
func (c *authController) Register(ctx *gin.Context) {
	var loginDTO = dto.RegisterDTO{}
	err := ctx.ShouldBind(&loginDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildError(http.StatusBadRequest, "Failed To Register", err.Error()))
		return
	}
	authResult, eror := c.authService.CreateUser(loginDTO)
	if eror != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildError(http.StatusBadRequest, "Failed To Register", eror.Error()))
		return
	}
	generateToken := c.jwtService.GenerateToken(strconv.FormatUint(authResult.ID, 10), authResult.Name)
	authResult.Token = generateToken
	response := helper.BuildResponse(http.StatusOK, "Success", authResult)
	ctx.JSON(http.StatusOK, response)

	// ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildError(http.StatusUnauthorized, "Check Your Email/Password", "Invalid Credential"))
}
