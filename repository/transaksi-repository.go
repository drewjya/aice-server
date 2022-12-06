package repository

import (
	"github.com/drewjya/aice-server/entity"
	"github.com/drewjya/aice-server/helper"

	"gorm.io/gorm"
)

type TransaksiRepository interface {
	CreateTransaction(transaksiData *entity.Transaksi) (*entity.Transaksi, error)
	GetTransactionHistoryToday(auth entity.TokenValue) (interface{}, error)
	GetTransactionHistoryThisWeek(auth entity.TokenValue) (interface{}, error)
	GetTransactionHistoryDetail(auth entity.TokenValue, transactionId uint64) (interface{}, error)
}

type transaksiConnection struct {
	connection *gorm.DB
}

func NewTransaksiRepository(db *gorm.DB) TransaksiRepository {
	return &transaksiConnection{
		connection: db,
	}
}

func (db *transaksiConnection) CreateTransaction(transaksi *entity.Transaksi) (*entity.Transaksi, error) {

	errTrx := db.connection.Save(&transaksi)
	if errTrx.Error != nil {
		return nil, errTrx.Error
	}

	return transaksi, nil
}

func (db *transaksiConnection) GetTransactionHistoryDetail(auth entity.TokenValue, transactionId uint64) (interface{}, error) {
	var transaksiHistory entity.Transaksi

	errTrx := db.connection.
		// .Where("id = ? AND userId = ?", transactionId, uint64(helper.IntCoverter(auth.UserId))).
		Where(&entity.Transaksi{ID: transactionId, UserID: uint64(helper.IntCoverter(auth.UserId))}).
		Take(&transaksiHistory)

	if errTrx.Error != nil {
		return nil, errTrx.Error
	}
	return transaksiHistory, nil
}
func (db *transaksiConnection) GetTransactionHistoryToday(auth entity.TokenValue) (interface{}, error) {
	var transaksiHistory []entity.TransaksiHistory
	errTrx := db.connection.Model(&entity.Transaksi{}).Where("TO_CHAR( \"createdAt\",'YYYY-mm-DD' ) = TO_CHAR( NOW(),'YYYY-mm-DD' ) AND \"userId\" = ?", auth.UserId).Find(&transaksiHistory)
	if errTrx.Error != nil {
		return nil, errTrx.Error
	}
	return transaksiHistory, nil
}
func (db *transaksiConnection) GetTransactionHistoryThisWeek(auth entity.TokenValue) (interface{}, error) {
	var transaksiHistory []entity.TransaksiHistory
	errTrx := db.connection.Model(&entity.Transaksi{}).Where("\"createdAt\"::date > CURRENT_DATE - 7AND \"userId\" = ?", auth.UserId).Find(&transaksiHistory)
	if errTrx.Error != nil {
		return nil, errTrx.Error
	}
	return transaksiHistory, nil
}
