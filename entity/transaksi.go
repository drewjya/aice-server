package entity

import "time"

type Transaksi struct {
	ID                    uint64    `gorm:"primary_key:auto_increment" json:"id"`
	UserID                uint64    `json:"userId"`
	CreatedAt             time.Time `json:"createdAt"`
	PilihanTokoID         uint64    `json:"pilihanTokoId"`
	NamaToko              string    `gorm:"type:varchar(255);" json:"namaToko"`
	KodeToko              string    `gorm:"type:varchar(50);" json:"kodeToko"`
	KualitasProduk        string    `json:"kualitasProduk"`
	StickerFreezer        string    `json:"stickerFreezer"`
	PapanHarga            string    `json:"papanHarga"`
	DividerKulkas         string    `json:"dividerKulkas"`
	LabelHarga            string    `json:"labelHarga"`
	WoblerPromo           string    `json:"woblerPromo"`
	Spanduk               string    `json:"spanduk"`
	BrandLain             string    `json:"brandLain"`
	StockBrandLain        string    `json:"stockBrandLain"`
	StockDibawahFreezer   string    `json:"stockDibawahFreezer"`
	ProdukPromosi         string    `json:"produkPromosi"`
	KebersihanBungaEs     string    `json:"kebersihanBungaEs"`
	KepenuhanFreezerAtas  string    `json:"kepenuhanFreezerAtas"`
	KebersihanLemFreezer  string    `json:"kebersihanLemFreezer"`
	KebersihanDebuFreezer string    `json:"kebersihanDebuFreezer"`
	PosisiFreezer         string    `json:"posisiFreezer"`
	JumlahPO              int       `json:"jumlahPO"`
	JumlahItemTerdisplay  int       `json:"jumlahItemTerdisplay"`
	SaranDanKendala       string    `json:"saranDanKendala"`
	ProdukRetur           string    `json:"produkRetur"`
	KategoriFreezer       string    `json:"kategoriFreezer"`
	NamaDistributor       string    `json:"namaDistributor"`
	FotoSelfie            string    `json:"fotoSelfie,omitempty"`
	FotoKulkasDariJauh    string    `json:"fotoKulkasDariJauh,omitempty"`
	FotoKulkasTertutup    string    `json:"fotoKulkasTertutup,omitempty"`
	FotoPO                string    `json:"fotoPo,omitempty"`
	FotoFreezerBawah      string    `json:"fotoFreezerBawah,omitempty"`
	FotoFreezerOne        string    `json:"fotoFreezerOne,omitempty"`
	FotoFreezerTwo        string    `json:"fotoFreezerTwo,omitempty"`
	FotoFreezerThree      string    `json:"fotoFreezerThree,omitempty"`
	FotoFreezerIsland1    string    `json:"fotoFreezerIsland1,omitempty"`
	FotoFreezerIsland2    string    `json:"fotoFreezerIsland2,omitempty"`
	FotoFreezerIsland3    string    `json:"fotoFreezerIsland3,omitempty"`
	FotoPop               string    `json:"fotoPop,omitempty"`
	FotoPeralatan         string    `json:"fotoPeralatan,omitempty"`
	FotoKulkasTerbuka     string    `json:"fotoKulkasTerbuka,omitempty"`
}

//	FotoID                uint64    `json:"fotoID"`
