package dto

type TambahTransaksiDTO struct {
	PilihanTokoID         uint64 `json:"pilihanTokoId" form:"pilihanTokoId" binding:"required"`
	NamaToko              string `json:"namaToko" form:"namaToko" binding:"required"`
	KodeToko              string `json:"kodeToko" form:"kodeToko" binding:"required"`
	KualitasProduk        string `json:"kualitasProduk" form:"kualitasProduk" binding:"required"`
	StickerFreezer        string `json:"stickerFreezer" form:"stickerFreezer" binding:"required"`
	PapanHarga            string `json:"papanHarga" form:"papanHarga" binding:"required"`
	DividerKulkas         string `json:"dividerKulkas" form:"dividerKulkas" binding:"required"`
	LabelHarga            string `json:"labelHarga" form:"labelHarga" binding:"required"`
	WoblerPromo           string `json:"woblerPromo" form:"woblerPromo" binding:"required"`
	Spanduk               string `json:"spanduk" form:"spanduk" binding:"required"`
	BrandLain             string `json:"brandLain" form:"brandLain" binding:"required"`
	StockBrandLain        string `json:"stockBrandLain" form:"stockBrandLain" binding:"required"`
	StockDibawahFreezer   string `json:"stockDibawahFreezer" form:"stockDibawahFreezer" binding:"required"`
	ProdukPromosi         string `json:"produkPromosi" form:"produkPromosi" binding:"required"`
	KebersihanBungaEs     string `json:"kebersihanBungaEs" form:"kebersihanBungaEs" binding:"required"`
	KepenuhanFreezerAtas  string `json:"kepenuhanFreezerAtas" form:"kepenuhanFreezerAtas" binding:"required"`
	KebersihanLemFreezer  string `json:"kebersihanLemFreezer" form:"kebersihanLemFreezer" binding:"required"`
	KebersihanDebuFreezer string `json:"kebersihanDebuFreezer" form:"kebersihanDebuFreezer" binding:"required"`
	PosisiFreezer         string `json:"posisiFreezer" form:"posisiFreezer" binding:"required"`
	JumlahPO              int    `json:"jumlahPO" form:"jumlahPO" binding:"required"`
	JumlahItemTerdisplay  int    `json:"jumlahItemTerdisplay" form:"jumlahItemTerdisplay" binding:"required"`
	SaranDanKendala       string `json:"saranDanKendala" form:"saranDanKendala" binding:"required"`
	ProdukRetur           string `json:"produkRetur" form:"produkRetur" binding:"required"`
	KategoriFreezer       string `json:"kategoriFreezer" form:"kategoriFreezer" binding:"required"`
	NamaDistributor       string `json:"namaDistributor" form:"namaDistributor" binding:"required"`
}
