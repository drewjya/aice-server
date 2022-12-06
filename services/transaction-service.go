package services

import (
	"aice/dto"
	"aice/entity"
	"aice/helper"
	"aice/repository"
	"errors"
	"log"
	"time"

	"github.com/mashingan/smapping"
)

type TransaksiService interface {
	CreateTransaksiNormal(auth entity.TokenValue, transaksiDTO dto.TambahTransaksiDTO, fotoSuperHyperDTO dto.FotoNormalDTO) (*entity.Transaksi, error)
	CreateTransaksiSuperHyper(auth entity.TokenValue, transaksiDTO dto.TambahTransaksiDTO, fotoSuperHyperDTO dto.FotoSuperHyperDTO) (*entity.Transaksi, error)
	GetTransaksiHistoryToday(auth entity.TokenValue) (interface{}, error)
	GetTransaksiHistoryTHisWeek(auth entity.TokenValue) (interface{}, error)
	GetTransactionHistoryDetail(auth entity.TokenValue, transactionId string) (interface{}, error)
}

type transaksiService struct {
	transaksiRepository repository.TransaksiRepository
}

func NewTransaksiService(transaksiRepository repository.TransaksiRepository) TransaksiService {
	return &transaksiService{transaksiRepository: transaksiRepository}
}

func (services *transaksiService) GetTransactionHistoryDetail(auth entity.TokenValue, transactionId string) (interface{}, error) {
	transaksiHistory, err := services.transaksiRepository.GetTransactionHistoryDetail(auth, uint64(helper.IntCoverter(transactionId)))
	if err != nil {
		return nil, err
	}
	return transaksiHistory, err
}
func (services *transaksiService) GetTransaksiHistoryTHisWeek(auth entity.TokenValue) (interface{}, error) {
	transaksiHistory, err := services.transaksiRepository.GetTransactionHistoryThisWeek(auth)
	if err != nil {
		return nil, err
	}
	return transaksiHistory, err
}
func (services *transaksiService) GetTransaksiHistoryToday(auth entity.TokenValue) (interface{}, error) {
	transaksiHistory, err := services.transaksiRepository.GetTransactionHistoryToday(auth)
	if err != nil {
		return nil, err
	}
	return transaksiHistory, err
}
func (services *transaksiService) CreateTransaksiNormal(auth entity.TokenValue, transaksiDTO dto.TambahTransaksiDTO, fotoNormalDTO dto.FotoNormalDTO) (*entity.Transaksi, error) {
	transaksi := entity.Transaksi{}
	err := smapping.FillStruct(&transaksi, smapping.MapFields(transaksiDTO))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	transaksi.CreatedAt = time.Now().UTC().Local()
	transaksi.UserID = 1

	fotoFreezerBawahExt := helper.ImageExtension(fotoNormalDTO.FotoFreezerBawah.Filename)
	fotoKulkasDariJauhExt := helper.ImageExtension(fotoNormalDTO.FotoKulkasDariJauh.Filename)
	fotoKulkasTerbukaExt := helper.ImageExtension(fotoNormalDTO.FotoKulkasTerbuka.Filename)
	fotoKulkasTertutupExt := helper.ImageExtension(fotoNormalDTO.FotoKulkasTertutup.Filename)
	fotoPOExt := helper.ImageExtension(fotoNormalDTO.FotoPO.Filename)
	fotoSelfieExt := helper.ImageExtension(fotoNormalDTO.FotoSelfie.Filename)
	var errD error
	fotoFreezerBawah := helper.RetrieveImage(fotoNormalDTO.FotoFreezerBawah, "FotoFreezerBawah/"+time.Now().Local().String()+fotoFreezerBawahExt)
	if fotoFreezerBawah == nil {
		errD = errors.New("gagal mengupload gambar")
		return nil, errD
	}
	fotoKulkasDariJauh := helper.RetrieveImage(fotoNormalDTO.FotoKulkasDariJauh, "FotoKulkasDariJauh/"+time.Now().Local().String()+fotoKulkasDariJauhExt)
	if fotoKulkasDariJauh == nil {
		errD = errors.New("gagal mengupload gambar")
		return nil, errD
	}
	fotoKulkasTerbuka := helper.RetrieveImage(fotoNormalDTO.FotoKulkasTerbuka, "FotoKulkasTerbuka/"+time.Now().Local().String()+fotoKulkasTerbukaExt)
	if fotoKulkasTerbuka == nil {
		errD = errors.New("gagal mengupload gambar")
		return nil, errD
	}
	fotoKulkasTertutup := helper.RetrieveImage(fotoNormalDTO.FotoKulkasTertutup, "FotoKulkasTertutup/"+time.Now().Local().String()+fotoKulkasTertutupExt)
	if fotoKulkasTertutup == nil {
		errD = errors.New("gagal mengupload gambar")
		return nil, errD
	}
	fotoPO := helper.RetrieveImage(fotoNormalDTO.FotoPO, "FotoPO/"+time.Now().Local().String()+fotoPOExt)
	if fotoPO == nil {
		errD = errors.New("gagal mengupload gambar")
		return nil, errD
	}
	fotoSelfie := helper.RetrieveImage(fotoNormalDTO.FotoSelfie, "FotoSelfie/"+time.Now().Local().String()+fotoSelfieExt)
	if fotoSelfie == nil {
		errD = errors.New("gagal mengupload gambar")
		return nil, errD
	}
	transaksi.FotoFreezerBawah = fotoFreezerBawah.(string)
	transaksi.FotoKulkasDariJauh = fotoKulkasDariJauh.(string)
	transaksi.FotoKulkasTerbuka = fotoKulkasTerbuka.(string)
	transaksi.FotoKulkasTertutup = fotoKulkasTertutup.(string)
	transaksi.FotoPO = fotoPO.(string)
	transaksi.FotoSelfie = fotoSelfie.(string)
	transaksi.UserID = uint64(helper.IntCoverter(auth.UserId))
	res, err := services.transaksiRepository.CreateTransaction(&transaksi)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (services *transaksiService) CreateTransaksiSuperHyper(auth entity.TokenValue, transaksiDTO dto.TambahTransaksiDTO, fotoSuperHyperDTO dto.FotoSuperHyperDTO) (*entity.Transaksi, error) {
	transaksi := entity.Transaksi{}
	err := smapping.FillStruct(&transaksi, smapping.MapFields(transaksiDTO))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	transaksi.CreatedAt = time.Now().UTC().Local()
	transaksi.UserID = 1
	var errD error
	fotoSelfieExt := helper.ImageExtension(fotoSuperHyperDTO.FotoSelfie.Filename)
	fotoKulkasDariJauhExt := helper.ImageExtension(fotoSuperHyperDTO.FotoKulkasDariJauh.Filename)
	fotoKulkasTertutupExt := helper.ImageExtension(fotoSuperHyperDTO.FotoKulkasTertutup.Filename)
	fotoPOExt := helper.ImageExtension(fotoSuperHyperDTO.FotoPO.Filename)
	fotoFreezerBawahExt := helper.ImageExtension(fotoSuperHyperDTO.FotoFreezerBawah.Filename)
	fotoFreezerOneExt := helper.ImageExtension(fotoSuperHyperDTO.FotoFreezerOne.Filename)
	fotoFreezerTwoExt := helper.ImageExtension(fotoSuperHyperDTO.FotoFreezerTwo.Filename)
	fotoFreezerThreeExt := helper.ImageExtension(fotoSuperHyperDTO.FotoFreezerThree.Filename)
	fotoFreezerIsland1Ext := helper.ImageExtension(fotoSuperHyperDTO.FotoFreezerIsland1.Filename)
	fotoFreezerIsland2Ext := helper.ImageExtension(fotoSuperHyperDTO.FotoFreezerIsland2.Filename)
	fotoFreezerIsland3Ext := helper.ImageExtension(fotoSuperHyperDTO.FotoFreezerIsland3.Filename)
	fotoPopExt := helper.ImageExtension(fotoSuperHyperDTO.FotoPop.Filename)
	fotoPeralatanExt := helper.ImageExtension(fotoSuperHyperDTO.FotoPeralatan.Filename)
	fotoSelfie := helper.RetrieveImage(fotoSuperHyperDTO.FotoSelfie, "FotoSelfie/"+time.Now().String()+fotoSelfieExt)
	if fotoSelfie == nil {
		errD = errors.New("gagal mengupload file")
		return nil, errD
	}
	fotoKulkasDariJauh := helper.RetrieveImage(fotoSuperHyperDTO.FotoKulkasDariJauh, "FotoKulkasDariJauh/"+time.Now().String()+fotoKulkasDariJauhExt)
	if fotoKulkasDariJauh == nil {
		errD = errors.New("gagal mengupload file")
		return nil, errD
	}
	fotoKulkasTertutup := helper.RetrieveImage(fotoSuperHyperDTO.FotoKulkasTertutup, "FotoKulkasTertutup/"+time.Now().String()+fotoKulkasTertutupExt)
	if fotoKulkasTertutup == nil {
		errD = errors.New("gagal mengupload file")
		return nil, errD
	}
	fotoPO := helper.RetrieveImage(fotoSuperHyperDTO.FotoPO, "FotoPO/"+time.Now().String()+fotoPOExt)
	if fotoPO == nil {
		errD = errors.New("gagal mengupload file")
		return nil, errD
	}
	fotoFreezerBawah := helper.RetrieveImage(fotoSuperHyperDTO.FotoFreezerBawah, "FotoFreezerBawah/"+time.Now().String()+fotoFreezerBawahExt)
	if fotoFreezerBawah == nil {
		errD = errors.New("gagal mengupload file")
		return nil, errD
	}
	fotoFreezerOne := helper.RetrieveImage(fotoSuperHyperDTO.FotoFreezerOne, "FotoFreezerOne/"+time.Now().String()+fotoFreezerOneExt)
	if fotoFreezerOne == nil {
		errD = errors.New("gagal mengupload file")
		return nil, errD
	}
	fotoFreezerTwo := helper.RetrieveImage(fotoSuperHyperDTO.FotoFreezerTwo, "FotoFreezerTwo/"+time.Now().String()+fotoFreezerTwoExt)
	if fotoFreezerTwo == nil {
		errD = errors.New("gagal mengupload file")
		return nil, errD
	}
	fotoFreezerThree := helper.RetrieveImage(fotoSuperHyperDTO.FotoFreezerThree, "FotoFreezerThree/"+time.Now().String()+fotoFreezerThreeExt)
	if fotoFreezerThree == nil {
		errD = errors.New("gagal mengupload file")
		return nil, errD
	}
	fotoFreezerIsland1 := helper.RetrieveImage(fotoSuperHyperDTO.FotoFreezerIsland1, "FotoFreezerIsland1/"+time.Now().String()+fotoFreezerIsland1Ext)
	if fotoFreezerIsland1 == nil {
		errD = errors.New("gagal mengupload file")
		return nil, errD
	}
	fotoFreezerIsland2 := helper.RetrieveImage(fotoSuperHyperDTO.FotoFreezerIsland2, "FotoFreezerIsland2/"+time.Now().String()+fotoFreezerIsland2Ext)
	if fotoFreezerIsland2 == nil {
		errD = errors.New("gagal mengupload file")
		return nil, errD
	}
	fotoFreezerIsland3 := helper.RetrieveImage(fotoSuperHyperDTO.FotoFreezerIsland3, "FotoFreezerIsland3/"+time.Now().String()+fotoFreezerIsland3Ext)
	if fotoFreezerIsland3 == nil {
		errD = errors.New("gagal mengupload file")
		return nil, errD
	}
	fotoPop := helper.RetrieveImage(fotoSuperHyperDTO.FotoPop, "FotoPop/"+time.Now().String()+fotoPopExt)
	if fotoPop == nil {
		errD = errors.New("gagal mengupload file")
		return nil, errD
	}
	fotoPeralatan := helper.RetrieveImage(fotoSuperHyperDTO.FotoPeralatan, "FotoPeralatan/"+time.Now().String()+fotoPeralatanExt)
	if fotoPeralatan == nil {
		errD = errors.New("gagal mengupload file")
		return nil, errD
	}

	transaksi.FotoSelfie = fotoSelfie.(string)
	transaksi.FotoKulkasDariJauh = fotoKulkasDariJauh.(string)
	transaksi.FotoKulkasTertutup = fotoKulkasTertutup.(string)
	transaksi.FotoPO = fotoPO.(string)
	transaksi.FotoFreezerBawah = fotoFreezerBawah.(string)
	transaksi.FotoFreezerOne = fotoFreezerOne.(string)
	transaksi.FotoFreezerTwo = fotoFreezerTwo.(string)
	transaksi.FotoFreezerThree = fotoFreezerThree.(string)
	transaksi.FotoFreezerIsland1 = fotoFreezerIsland1.(string)
	transaksi.FotoFreezerIsland2 = fotoFreezerIsland2.(string)
	transaksi.FotoFreezerIsland3 = fotoFreezerIsland3.(string)
	transaksi.FotoPop = fotoPop.(string)
	transaksi.UserID = uint64(helper.IntCoverter(auth.UserId))
	transaksi.FotoPeralatan = fotoPeralatan.(string)

	res, err := services.transaksiRepository.CreateTransaction(&transaksi)
	if err != nil {
		return nil, err
	}
	return res, nil

}
