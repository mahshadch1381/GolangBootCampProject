package get

import (
	"finalproject/db"
	"finalproject/internal/models"
)

func GetLastVoucherItemIndex(voucherID uint) (uint, error) {
	var lastItem models.VoucherItem
	db := db.GetDB()

	result := db.Where("voucher_id = ?", voucherID).Order("id DESC").First(&lastItem)
	if result.Error != nil {
		return 0, result.Error
	}

	return lastItem.ID, nil
}
