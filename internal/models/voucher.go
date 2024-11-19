package models

import (
	"errors"
	"gorm.io/gorm"
)
type Voucher struct {
    ID      uint   `gorm:"primaryKey"`
    Number  string `gorm:"size:64;not null;unique"`
    Version int    `gorm:"default:1"`
}

func (Voucher) TableName() string {
	return "voucher" 
}

func (v *Voucher) BeforeUpdate(tx *gorm.DB) (err error) {
	var existing Voucher
	if err := tx.First(&existing, v.ID).Error; err != nil {
		return err 
	}
	if existing.Version != v.Version {
		return errors.New("version mismatch: record has been updated by another transaction")
	}
	v.Version++
	return nil
}