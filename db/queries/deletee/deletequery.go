package delete

import (
	"errors"
	"finalproject/db"
	"gorm.io/gorm"
	"strings"
)

func DeleteRecord[T any](id uint) error {
	db := db.GetDB() 
	var record T
	result := db.Delete(&record, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrCheckConstraintViolated) {
			return errors.New("ErrCheckConstraintViolated error")
		}
		if errors.Is(result.Error, gorm.ErrEmptySlice) {
			return errors.New("ErrEmptySlice error")
		}
		if errors.Is(result.Error, gorm.ErrRecordNotFound)  {
			return errors.New("NOT found  record error")
		}
		if result.Error== gorm.ErrForeignKeyViolated || strings.Contains(result.Error.Error(), "violates foreign key constraint"){
			return errors.New("item you want to delete,has been refrensed by voucher items1")
		}
		
		return errors.New("cant to delete dl from database")
	} else {
		if result.RowsAffected == 0 {
			return errors.New("no records found with the specified ID")
		}
	}
	return nil
}
