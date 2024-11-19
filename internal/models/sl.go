package models

import (
	"errors"
	"gorm.io/gorm"
	"fmt"
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
func (sl *SL) BeforeUpdate(tx *gorm.DB) (err error) {
	var existing SL
	if err := tx.First(&existing, sl.ID).Error; err != nil {
		return err 
	}
	if existing.Version != sl.Version {
		return errors.New("version mismatch: record has been updated by another transaction")
	}
	sl.Version++
	return nil
}
func (sl *SL) BeforeDelete(tx *gorm.DB) (err error) {
	var count int64

	// بررسی استفاده از SL در VoucherItem
	if err := tx.Model(&VoucherItem{}).Where("sl_id = ?", sl.ID).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to check references for SL: %w", err)
	}

	if count > 0 {
		return fmt.Errorf("SL with ID %d cannot be deleted because it is referenced in %d voucher items", sl.ID, count)
	}

	return nil
}
