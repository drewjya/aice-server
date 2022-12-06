package dto

import (
	"mime/multipart"

	"aice-server/helper"

	"github.com/go-playground/validator"
)

type FotoNormalDTO struct {
	FotoFreezerBawah   *multipart.FileHeader `form:"fotoFreezerBawah" binding:"required"`
	FotoKulkasDariJauh *multipart.FileHeader `form:"fotoKulkasDariJauh" binding:"required"`
	FotoKulkasTerbuka  *multipart.FileHeader `form:"fotoKulkasTerbuka" binding:"required"`
	FotoKulkasTertutup *multipart.FileHeader `form:"fotoKulkasTertutup" binding:"required"`
	FotoPO             *multipart.FileHeader `form:"fotoPO" binding:"required"`
	FotoSelfie         *multipart.FileHeader `form:"fotoSelfie" binding:"required"`
}

func FotoNormalDTOValidation(val validator.StructLevel) {
	fotoNormalDTO := val.Current().Interface().(FotoNormalDTO)
	helper.MultipartDTOValidation(fotoNormalDTO, val)
}
