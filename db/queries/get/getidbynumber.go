package get

import (

	"finalproject/db"
)

func GetIDByNumber(tableName string, number string) (uint, error) {
	db := db.GetDB() 
	var result uint 
	err := db.Table(tableName).Select("id").Where("number = ?", number).Scan(&result).Error
	if err != nil {
		return 0, err
	}
	return result, nil
}
