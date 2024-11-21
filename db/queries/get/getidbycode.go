package get

import (

	"finalproject/db"
)


func GetIDByCode(tableName string, code string) (uint, error) {
	db := db.GetDB() 
	var result uint 
	err := db.Table(tableName).Select("id").Where("code = ?", code).Scan(&result).Error
	if err != nil {
		return 0, err
	}
	return result, nil
}