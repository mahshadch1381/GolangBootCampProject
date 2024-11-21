package models

import (
	//"errors"
	//"gorm.io/gorm"
	//"fmt"
)

type SL struct {
    ID          uint   `gorm:"primaryKey"`
    Code        string `gorm:"size:64;not null;unique"`
    Title       string `gorm:"size:64;not null;unique"`
    IsDetailable bool   `gorm:"not null"`
    Version     int    `gorm:"default:1"`
}

func (SL) TableName() string {
	return "sl" 
}
// func (sl *SL) BeforeUpdate(tx *gorm.DB) (err error) {
// 	var existing SL
// 	if err := tx.First(&existing, sl.ID).Error; err != nil {
// 		return err 
// 	}
// 	if existing.Version != sl.Version {
// 		return errors.New("version mismatch: record has been updated by another transaction")
// 	}
// 	sl.Version++
// 	return nil
// }

