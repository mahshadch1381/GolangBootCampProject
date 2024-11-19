package count

import (
	 "finalproject/db"
	
)
func CountByNumber(number string,table  interface{}) (int64, error) {
	db.ConnectWithGORM()
	var count int64
	if err := db.DB.Model(table).Where("number = ?", number).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
func CountByNumberExcludingID( number string, id uint, table interface{}) (int64, error) {
	var count int64
	if err := db.DB.Model(table).Where("number = ? AND id != ?", number, id).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
