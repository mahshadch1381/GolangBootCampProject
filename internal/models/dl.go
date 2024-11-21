package models

import (
	
   
	
)
type DL struct {
    ID      uint   `gorm:"primaryKey"`
    Code    string `gorm:"size:64;not null;unique"`
    Title   string `gorm:"size:64;not null;unique"`
    Version int    `gorm:"default:1"`
}

func (DL) TableName() string {
	return "dl" 
}

