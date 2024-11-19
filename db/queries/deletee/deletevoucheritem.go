package delete

import (
	"errors"
	"finalproject/db"
	"finalproject/internal/models"
)


func DeleteVoucherItemsByVoucherID(voucherID uint) error {
	db.ConnectWithGORM()
	// Delete all VoucherItems with the given voucher_id
	if result := db.DB.Where("voucher_id = ?", voucherID).Delete(&models.VoucherItem{}); result.Error != nil {
		return errors.New("failed to delete voucher items")
	} else {
		if result.RowsAffected == 0 {
			return errors.New("no records found with the specified ID")
		}
	}
	return nil
}
func DeleteVoucherItemsByIdAndVoucherID(voucherID uint,id uint) error {
	db.ConnectWithGORM()
	// Delete all VoucherItems with the given voucher_id
	if result := db.DB.Where("voucher_id = ? AND id = ?", voucherID,id).Delete(&models.VoucherItem{}); result.Error != nil {
		return errors.New("failed to delete voucher items")
	} else {
		if result.RowsAffected == 0 {
			return errors.New("no records found with the specified ID and VoucherId")
		}
	}
	return nil
}