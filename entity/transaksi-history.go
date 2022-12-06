package entity

import "time"

type TransaksiHistory struct {
	ID                    uint64    `gorm:"column:id,primary_key:auto_increment" json:"id"`
	UserID                uint64    `gorm:"column:userId" json:"userId"`
	CreatedAt             time.Time `gorm:"column:createdAt" json:"createdAt"`
	PilihanTokoID         uint64    `gorm:"column:pilihanTokoId" json:"pilihanTokoId"`
	NamaToko              string    `gorm:"column:namaToko,type:varchar(255);" json:"namaToko"`
	KodeToko              string    `gorm:"column:kodeToko,type:varchar(50);" json:"kodeToko"`
}