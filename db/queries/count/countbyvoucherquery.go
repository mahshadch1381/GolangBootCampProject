package count

import (
	 "finalproject/db"
	
)
func CountByVoucher(vouchrid uint,table  interface{}) (int64, error) {
	db.ConnectWithGORM()
	var count int64
	if err := db.DB.Model(table).Where("voucher_id = ?", vouchrid).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}