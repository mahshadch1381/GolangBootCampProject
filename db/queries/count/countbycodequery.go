package count

import (
	
	"gorm.io/gorm"
)
func CountByCode(db *gorm.DB, code string,table  interface{}) (int64, error) {
	var count int64
	if err := db.Model(table).Where("code = ?", code).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
func CountByCodeExcludingID(db *gorm.DB, code string, id uint, table interface{}) (int64, error) {
	var count int64
	if err := db.Model(table).
		Where("code = ? AND id != ?", code, id).
		Count(&count).Error; err != nil {
		return 0, err
	}
	
	return count, nil
}

