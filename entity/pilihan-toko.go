package entity

type PilihanToko struct {
	ID   uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name string `gorm:"uniqueIndex;type:varchar(255);" json:"name"`
}
