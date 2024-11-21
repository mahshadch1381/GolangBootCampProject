package models

type VoucherItem struct {
	ID        uint  `gorm:"primaryKey"`
	SLID      uint  `gorm:"not null;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	DLID      *uint `gorm:"constraint:OnDelete:RESTRICT;"`
	VoucherId uint  `gorm:"not null"` //foregin key to voucher must be set, so we can decide to not be a pointer
	Debit     int   `gorm:"default:0"`
	Credit    int   `gorm:"default:0"`
}

func (VoucherItem) TableName() string {
	return "voucheritem"
}
