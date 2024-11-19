package insert

import (
	"errors"
    "finalproject/db"
	"gorm.io/gorm"
)

func InsertRecord[T any](record T) error {
	db.ConnectWithGORM()
	result := db.DB.Create(&record)
	if result.Error != nil {
		if errors.Is(result.Error,gorm.ErrCheckConstraintViolated){
			return errors.New("ErrCheckConstraintViolated error")
		}
		if errors.Is(result.Error,gorm.ErrEmptySlice){
			return errors.New("ErrEmptySlice error")
		}
		return errors.New("cant to insert to database ")
	}
	return nil
}