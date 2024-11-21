package get

import (
	"errors"
	"finalproject/db"
	"finalproject/internal/models"

	"gorm.io/gorm"
)

func GetRecordByID[T any](id uint, result *T) error {
	db := db.GetDB() 
	err := db.First(result, id).Error
	if err != nil {
		if errors.Is(err,gorm.ErrCheckConstraintViolated){
			return errors.New("ErrCheckConstraintViolated error")
		}
		if errors.Is(err,gorm.ErrEmptySlice){
			return errors.New("ErrEmptySlice error")
		}
		if errors.Is(err,gorm.ErrRecordNotFound){
			return errors.New("NOT found  record error")
		}
		return errors.New("cant to get dl to database ")
	}
	return nil
}
func GetLastID(tableName string) (uint, error) {
	db := db.GetDB() 
	var lastItem uint
	result := db.Table(tableName).Order("id DESC").Limit(1).Select("id").Scan(&lastItem)
	if result.Error != nil {
		return 0, result.Error
	}
	return lastItem, nil
}
func GetIDByCode(tableName string, code string) (uint, error) {
	db := db.GetDB() 
	var result uint 
	err := db.Table(tableName).Select("id").Where("code = ?", code).Scan(&result).Error
	if err != nil {
		return 0, err
	}
	return result, nil
}
func GetIDByNumber(tableName string, number string) (uint, error) {
	db := db.GetDB() 
	var result uint 
	err := db.Table(tableName).Select("id").Where("number = ?", number).Scan(&result).Error
	if err != nil {
		return 0, err
	}
	return result, nil
}

func GetLastVoucherItemIndex(voucherID uint) (uint, error) {
	var lastItem models.VoucherItem
	db := db.GetDB() 
	
	result := db.Where("voucher_id = ?", voucherID).Order("id DESC").First(&lastItem)
	if result.Error != nil {
		return 0, result.Error
	}

	return lastItem.ID, nil
}