package update

import (
	"errors"
    "finalproject/db"
	"gorm.io/gorm"
)

func Update[T any](existing T) error {
	db.ConnectWithGORM()
	result := db.DB.Save(&existing)
	if result.Error != nil {
		if errors.Is(result.Error,gorm.ErrCheckConstraintViolated){
			return errors.New("ErrCheckConstraintViolated error")
		}
		if errors.Is(result.Error,gorm.ErrEmptySlice){
			return errors.New("ErrEmptySlice error")
		}
		return errors.New("cant to update dl to database ")
	}
	return nil
}