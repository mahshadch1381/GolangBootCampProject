package mapper

import (
	"finalproject/internal/models"
	"finalproject/internal/request/vrequest"
)

func VoucherMapper(req *vrequest.VoucherInsertRequest) *models.Voucher {
	voucher := &models.Voucher{
		Number:  req.Voucher.Number,
		Version: 1,
	}
	return voucher
}
func VoucherMapperUpdate(req *vrequest.VoucherUpdateRequest, existingv *models.Voucher) *models.Voucher {

	existingv.Number = req.Voucher.Number
	existingv.Version=req.Voucher.Version+1
	return existingv
}
