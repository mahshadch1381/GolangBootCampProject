package get

import (
	"finalproject/internal/models"
    "finalproject/db"
	
)

func GetVoucherItemsByVoucherID(voucherID uint) (*[]models.VoucherItem, error) {
	db := db.GetDB() 
	var items []models.VoucherItem
	err := db.Where("voucher_id = ?", voucherID).Find(&items).Error
	if err != nil {
		return nil, err
	}
	return &items, nil
}
