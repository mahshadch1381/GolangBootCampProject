package count

import (
	"finalproject/db"
	"finalproject/internal/models"
	"fmt"
)
func CountBySlrefrences(sl_id *uint) (*int64, error) {
	db.ConnectWithGORM()
	var count int64

	if err := db.DB.Model(&models.VoucherItem{}).Where("sl_id = ?", sl_id).Count(&count).Error; err != nil {
		return nil,fmt.Errorf("failed to check references for SL: %w", err)
	}
	
	return &count,nil
}