package dto

import (
	"mime/multipart"

	"aice-server/helper"

	"github.com/go-playground/validator"
)

type FotoSuperHyperDTO struct {
	FotoSelfie         *multipart.FileHeader `json:"fotoSelfie" binding:"required"`
	FotoKulkasDariJauh *multipart.FileHeader `json:"fotoKulkasDariJauh" binding:"required"`
	FotoKulkasTertutup *multipart.FileHeader `json:"fotoKulkasTertutup" binding:"required"`
	FotoPO             *multipart.FileHeader `json:"fotoPo" binding:"required"`
	FotoFreezerBawah   *multipart.FileHeader `json:"fotoFreezerBawah" binding:"required"`
	FotoFreezerOne     *multipart.FileHeader `json:"fotoFreezerOne" binding:"required"`
	FotoFreezerTwo     *multipart.FileHeader `json:"fotoFreezerTwo" binding:"required"`
	FotoFreezerThree   *multipart.FileHeader `json:"fotoFreezerThree" binding:"required"`
	FotoFreezerIsland1 *multipart.FileHeader `json:"fotoFreezerIsland1" binding:"required"`
	FotoFreezerIsland2 *multipart.FileHeader `json:"fotoFreezerIsland2" binding:"required"`
	FotoFreezerIsland3 *multipart.FileHeader `json:"fotoFreezerIsland3" binding:"required"`
	FotoPop            *multipart.FileHeader `json:"fotoPop" binding:"required"`
	FotoPeralatan      *multipart.FileHeader `json:"fotoPeralatan" binding:"required"`
}

func FotoSuperhyperDTOValidation(val validator.StructLevel) {
	fotoSuperHyperDTO := val.Current().Interface().(FotoSuperHyperDTO)
	helper.MultipartDTOValidation(fotoSuperHyperDTO, val)

}
