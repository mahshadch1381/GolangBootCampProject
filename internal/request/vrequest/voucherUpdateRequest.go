package vrequest

import (
	"errors"
	"finalproject/db/queries/count"
	"finalproject/db/queries/get"
	"finalproject/db/queries/sum"
	"finalproject/internal/models"
	"finalproject/internal/validation"
	
)

type VoucherItemUpdate struct {
	ID        uint
	VoucherID uint
	SLID      uint  `json:"sl_id"`
	DLID      *uint `json:"dl_id"`
	Debit     int
	Credit    int
}

type VoucherItemUpdateList struct {
	Inserted []VoucherItemInsertion
	Updated  []VoucherItemUpdate
	Deleted  []uint
}

type VoucherUpdateRequest struct {
	Voucher models.Voucher
	Items   VoucherItemUpdateList
}



func (r *VoucherUpdateRequest) Validate() ( *models.Voucher,error) {
	voucher ,err1:=r.ExistnceOfVoucher()
	if err1 != nil {
		return nil,err1
	}
	if err := validation.LenghtValidation(&r.Voucher.Number); err != nil {
		return nil,err
	}
	if err := r.ValidateCountOfItems(); err != nil {
		return nil,err
	}
	if err := r.ValidateDebitsAndCredits(); err != nil {
		return nil,err
	}
	return voucher,nil
}

func (r *VoucherUpdateRequest) ExistnceOfVoucher() ( *models.Voucher,error) {
	var voucher models.Voucher
	if err := get.GetRecordByID(r.Voucher.ID, &voucher); err != nil {
		str:=err.Error()
		str = "cant find voucher that you want to update :" +str
		return nil,errors.New(str)
	}
	return &voucher,nil
}
func (r *VoucherUpdateRequest) ValidateCountOfItems() error {
	if len(r.Items.Inserted) == 0 && len(r.Items.Deleted) == 0 {
		return nil
	}
	var counter int64
	counter, err := count.CountByVoucher(r.Voucher.ID, &models.VoucherItem{})
	if err != nil {
		return err
	}
	if (counter + int64(len(r.Items.Inserted))) - int64(len(r.Items.Deleted))> 500 {
		return errors.New("we cant insert your items.bcz we should have at most 500 items")
	}
	if (counter + int64(len(r.Items.Inserted)))- int64(len(r.Items.Deleted)) < 2 {
		return errors.New("we cant delete your items.bcz we should have at leat 2 items")
	}
	return nil
}
func (r *VoucherUpdateRequest) ValidateItemsOfInserted() (*int , *int ,error) {
	var debits, credits int
	for _, item := range r.Items.Inserted {
		if !((item.Debit > 0 && item.Credit == 0) || (item.Debit == 0 && item.Credit > 0)) {
			return nil,nil,errors.New("each item must have either debit or credit greater than 0, and the other 0 ")
		}
		
		if err := CheckRefrences(item.DLID, item.SLID); err != nil {
			return  nil,nil,err
		}
		debits += item.Debit
		credits += item.Credit

	}
   return &debits,&credits,nil
}
func (r *VoucherUpdateRequest) ValidateItemsOfUpdated(debits *int,credits *int) (*[]uint ,error) {
	listupdatedId := []uint{}
	for _, item := range r.Items.Updated {
		if !((item.Debit > 0 && item.Credit == 0) || (item.Debit == 0 && item.Credit > 0)) {
			return nil,errors.New("each item must have either debit or credit greater than 0, and the other 0 ")
		}
		if err := CheckRefrences(item.DLID, item.SLID); err != nil {
			return nil,err
		}
		*debits += item.Debit
		*credits += item.Credit
		listupdatedId = append(listupdatedId, item.ID)
	}
   return &listupdatedId,nil
}

func CheckRefrences(dl_id *uint, sl_id uint) error {
	if sl_id == 0 {
		return errors.New("sL field in voucher items should not be empty")
	}
	var sl models.SL
	if err := get.GetRecordByID(sl_id, &sl); err != nil {
		str:=err.Error()
		str = "cant find refrenced Sl :" +str
		return errors.New(str)
		
	}
	if !sl.IsDetailable {
		if dl_id != nil {
			return errors.New("the sl in voucheritem does not have dl so dl_id should be empty")
		}
	} else {
		if dl_id == nil {
			return errors.New("the sl in voucheritem  has dl so dl_id should not be empty")
		}
		var dl models.DL
		if err := get.GetRecordByID(*dl_id, &dl); err != nil {
			str:=err.Error()
			str = "cant find refrenced dl :" +str
			return errors.New(str)
		}
	}
	return nil
}

func (r *VoucherUpdateRequest) ValidateDebitsAndCredits() error {
	var debits, credits *int
	debits,credits,err:=r.ValidateItemsOfInserted()
	if err!= nil{
         return err
	}
	listupdatedId ,err:= r.ValidateItemsOfUpdated(debits,credits)
	if err!= nil{
		return err
   }
	lastdebits, lastcredits, err := sum.SumCreditAndDebitsForUpdate(r.Voucher.ID, &r.Items.Deleted, listupdatedId)
	if err != nil {
		return err
	}
	if *lastcredits+*credits != *lastdebits+*debits {
		return errors.New("sum of credits and debits are not 0 ")
	}
	return nil
}
