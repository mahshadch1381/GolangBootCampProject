package dto

import (
	"finalproject/internal/models"
)

type  Voucherdto struct{
	Voucher  models.Voucher
	Items     []models.VoucherItem
}