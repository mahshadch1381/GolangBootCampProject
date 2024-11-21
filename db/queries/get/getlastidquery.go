package get

import (
	
	"finalproject/db"
)
func GetLastID(tableName string) (uint, error) {
	db := db.GetDB() 
	var lastItem uint
	result := db.Table(tableName).Order("id DESC").Limit(1).Select("id").Scan(&lastItem)
	if result.Error != nil {
		return 0, result.Error
	}
	return lastItem, nil
}