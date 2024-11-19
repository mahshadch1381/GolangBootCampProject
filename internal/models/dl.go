package models

import (
	"errors"
    "fmt"
	"gorm.io/gorm"
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
func (dl *DL) BeforeUpdate(tx *gorm.DB) (err error) {
	var existing DL
	if err := tx.First(&existing, dl.ID).Error; err != nil {
		return err 
	}
	if existing.Version != dl.Version {
		return errors.New("version mismatch: record has been updated by another transaction")
	}
	dl.Version++
	return nil
}
func (dl *DL) BeforeDelete(tx *gorm.DB) (err error) {
	var count int64

	// بررسی استفاده از DL در VoucherItem
	if err := tx.Model(&VoucherItem{}).Where("dl_id = ?", dl.ID).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to check references for DL: %w", err)
	}

	if count > 0 {
		return fmt.Errorf("DL with ID %d cannot be deleted because it is referenced in %d voucher items", dl.ID, count)
	}

	return nil
}
