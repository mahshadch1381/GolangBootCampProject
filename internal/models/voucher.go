package models


type Voucher struct {
    ID      uint   `gorm:"primaryKey"`
    Number  string `gorm:"size:64;not null;unique"`
    Version int    `gorm:"default:1"`
}

func (Voucher) TableName() string {
	return "voucher" 
}
