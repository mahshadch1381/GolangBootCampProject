package sum

import (
	"finalproject/db"
	"finalproject/internal/models"
)

func SumCreditAndDebitsForUpdate(voucherID uint, updatedIDs *[]uint, deletedIDs *[]uint) (*int, *int, error) {
	var totalDebit, totalCredit int
	db := db.GetDB() 
	if len(*updatedIDs) == 0 && len(*deletedIDs) == 0 {
		err := db.Model(&models.VoucherItem{}).
			Select("SUM(debit), SUM(credit)").
			Where("voucher_id = ?", voucherID).
			Row().
			Scan(&totalDebit, &totalCredit)
		if err != nil {
			return nil, nil, err
		}
		return &totalDebit, &totalCredit, nil
	} else {
		err := db.Model(&models.VoucherItem{}).
			Select("SUM(debit), SUM(credit)").
			Where("voucher_id = ?", voucherID).
			Where("id NOT IN (?)", append(*updatedIDs, *deletedIDs...)).
			Row().
			Scan(&totalDebit, &totalCredit)
		if err != nil {
			return nil, nil, err
		}
		return &totalDebit, &totalCredit, nil
	}
}
