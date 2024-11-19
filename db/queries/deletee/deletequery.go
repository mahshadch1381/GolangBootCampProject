package delete

import (
	"errors"
	"finalproject/db"
	"gorm.io/gorm"
)

func DeleteRecord[T any](id uint) error {
	db.ConnectWithGORM()
	var record T
	println(id)
	result := db.DB.Delete(&record, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrCheckConstraintViolated) {
			return errors.New("ErrCheckConstraintViolated error")
		}
		if errors.Is(result.Error, gorm.ErrEmptySlice) {
			return errors.New("ErrEmptySlice error")
		}
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("NOT found  record error")
		}
		return errors.New("cant to delete dl to database ")
	} else {
		if result.RowsAffected == 0 {
			return errors.New("no records found with the specified ID")
		}
	}
	return nil
}
