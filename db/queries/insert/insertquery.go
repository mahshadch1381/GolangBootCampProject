package insert

import (
	"errors"
	"finalproject/db"
	"strings"

	"gorm.io/gorm"
)

func InsertRecord[T any](record T) error {
	db := db.GetDB()
	result := db.Create(&record)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return errors.New("cant to insert to database : we have duplicate value in table ")
		}
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) || strings.Contains(result.Error.Error(), "duplicate key") {
			return errors.New("cant to insert to database: have duplicate code value ")
		}
		if result.Error == gorm.ErrForeignKeyViolated || strings.Contains(result.Error.Error(), "violates foreign key constraint") {
			return errors.New("item you want to insert, as refrensed , not exist")
		}
		return result.Error
	}
	return nil
}
