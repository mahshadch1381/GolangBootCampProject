package update

import (
	"errors"
    "finalproject/db"
	"gorm.io/gorm"
)

func Update[T any](existing T) error {
	db := db.GetDB() 
	result := db.Save(&existing)
	if result.Error != nil {
		if errors.Is(result.Error,gorm.ErrCheckConstraintViolated){
			return errors.New("ErrCheckConstraintViolated error")
		}
		if errors.Is(result.Error,gorm.ErrEmptySlice){
			return errors.New("ErrEmptySlice error")
		}
		return result.Error
	}
	return nil
}