package mapper

import (
	"finalproject/internal/models"
	"finalproject/internal/request/vrequest"
)

func VoucherItemMapper(req *vrequest.VoucherItemInsertion) *models.VoucherItem {
	item := &models.VoucherItem{
		                       SLID: req.SLID,
		                       DLID: req.DLID, 
							   Credit: req.Credit,
							   Debit: req.Debit,
							}
	return item

}
func VoucheItemMapperUpdate(item *vrequest.VoucherItemUpdate, existingvoucheritem *models.VoucherItem) *models.VoucherItem {
	existingvoucheritem.SLID=item.SLID
	existingvoucheritem.DLID=item.DLID
	existingvoucheritem.Credit=item.Credit
	existingvoucheritem.Debit=item.Debit

	return existingvoucheritem
}

