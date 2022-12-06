package controller

import (
	"net/http"

	"github.com/drewjya/aice-server/dto"
	"github.com/drewjya/aice-server/entity"
	"github.com/drewjya/aice-server/helper"
	"github.com/drewjya/aice-server/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
)

type TransaksiController interface {
	TambahTransaksiNormal(ctx *gin.Context, validate *validator.Validate)
	TambahTransaksiSuperHyper(ctx *gin.Context)
	GetTransaksiHistoryToday(ctx *gin.Context)
	GetTransaksiHistoryTHisWeek(ctx *gin.Context)
	GetTransactionHistoryDetail(ctx *gin.Context)
}

type transaksiController struct {
	transaksiServices services.TransaksiService
}

func NewTransaksiController(transaksiServices services.TransaksiService) TransaksiController {
	return &transaksiController{
		transaksiServices: transaksiServices,
	}
}

func (c *transaksiController) GetTransactionHistoryDetail(ctx *gin.Context) {
	auth, _ := ctx.Get("Token")
	transactionId := ctx.Param("id")
	history, err := c.transaksiServices.GetTransactionHistoryDetail(auth.(entity.TokenValue), transactionId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusAccepted, helper.BuildResponse(http.StatusAccepted, "Login Success", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(http.StatusOK, "OK", history))
}
func (c *transaksiController) GetTransaksiHistoryTHisWeek(ctx *gin.Context) {
	auth, _ := ctx.Get("Token")
	history, err := c.transaksiServices.GetTransaksiHistoryTHisWeek(auth.(entity.TokenValue))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusAccepted, helper.BuildResponse(http.StatusAccepted, "Login Success", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(http.StatusOK, "OK", history))
}
func (c *transaksiController) GetTransaksiHistoryToday(ctx *gin.Context) {
	auth, _ := ctx.Get("Token")
	history, err := c.transaksiServices.GetTransaksiHistoryToday(auth.(entity.TokenValue))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusAccepted, helper.BuildResponse(http.StatusAccepted, "Login Success", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(http.StatusOK, "OK", history))

}
func (c *transaksiController) TambahTransaksiNormal(ctx *gin.Context, validate *validator.Validate) {
	auth, _ := ctx.Get("Token")
	var fotoNormalDTO = dto.FotoNormalDTO{}
	errDTO := ctx.ShouldBindWith(&fotoNormalDTO, binding.FormMultipart)
	if errDTO != nil {
		ctx.AbortWithStatusJSON(http.StatusAccepted, helper.BuildResponse(http.StatusAccepted, "Login Success", errDTO.Error()))
		return
	}
	err := helper.ValidateStruct(validate, fotoNormalDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusAccepted, helper.BuildError(http.StatusAccepted, "Login Success", err.(string)))
		return
	}

	var transaksiDTO = dto.TambahTransaksiDTO{}
	errDTO = ctx.ShouldBind(&transaksiDTO)
	if errDTO != nil {
		ctx.AbortWithStatusJSON(http.StatusAccepted, helper.BuildResponse(http.StatusAccepted, "Login Success", errDTO.Error()))
		return
	}

	transaksiRes, eror := c.transaksiServices.CreateTransaksiNormal((auth.(entity.TokenValue)), transaksiDTO, fotoNormalDTO)
	if eror != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildError(http.StatusBadRequest, "Failed To Register", eror.Error()))
		return
	}

	response := helper.BuildResponse(http.StatusOK, "Success", transaksiRes)
	ctx.JSON(http.StatusOK, response)
}
func (c *transaksiController) TambahTransaksiSuperHyper(ctx *gin.Context) {
	auth, _ := ctx.Get("Token")
	var fotoSuperHyperDTO = dto.FotoSuperHyperDTO{}
	errDTO := ctx.ShouldBindWith(&fotoSuperHyperDTO, binding.FormMultipart)
	if errDTO != nil {
		ctx.AbortWithStatusJSON(http.StatusAccepted, helper.BuildResponse(http.StatusAccepted, "Login Success", errDTO.Error()))
		return
	}
	var transaksiDTO = dto.TambahTransaksiDTO{}
	errDTO = ctx.ShouldBind(&transaksiDTO)
	if errDTO != nil {
		ctx.AbortWithStatusJSON(http.StatusAccepted, helper.BuildResponse(http.StatusAccepted, "Login Success", errDTO.Error()))
		return
	}

	transaksiRes, eror := c.transaksiServices.CreateTransaksiSuperHyper((auth.(entity.TokenValue)), transaksiDTO, fotoSuperHyperDTO)
	if eror != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildError(http.StatusBadRequest, "Failed To Register", eror.Error()))
		return
	}

	response := helper.BuildResponse(http.StatusOK, "Success", transaksiRes)
	ctx.JSON(http.StatusOK, response)
}
