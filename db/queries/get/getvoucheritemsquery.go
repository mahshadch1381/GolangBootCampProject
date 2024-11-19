package get

import (
	"finalproject/internal/models"
    "finalproject/db"
	
)

func GetVoucherItemsByVoucherID(voucherID uint) (*[]models.VoucherItem, error) {
	db.ConnectWithGORM()
	var items []models.VoucherItem
	err := db.DB.Where("voucher_id = ?", voucherID).Find(&items).Error
	if err != nil {
		return nil, err
	}
	return &items, nil
}
