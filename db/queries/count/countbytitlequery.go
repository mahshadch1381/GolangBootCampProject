package count

import (
	
	"gorm.io/gorm"
)

func CountByTitle(db *gorm.DB, title string,table  interface{}) (int64, error) {
	var count int64
	if err := db.Model(table).Where("title = ?", title).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
func CountByTitleExcludingID(db *gorm.DB, title string, id uint,table  interface{}) (int64, error) {
	var count int64
	if err := db.Model(table).Where("title = ? AND id != ?", title, id).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}