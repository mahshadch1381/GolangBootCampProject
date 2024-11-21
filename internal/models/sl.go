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

