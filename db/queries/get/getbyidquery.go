package get

import (
	"errors"
    "finalproject/db"
	"gorm.io/gorm"
)

func GetRecordByID[T any](id uint, result *T) error {
	db.ConnectWithGORM()
	err := db.DB.First(result, id).Error
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