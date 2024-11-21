package services

import (
	"errors"
	"finalproject/db/queries/get"
	"finalproject/db/queries/deletee"
	"finalproject/db/queries/insert"
	"finalproject/db/queries/update"
	"finalproject/internal/dto"
	"finalproject/internal/mapper"
	"finalproject/internal/models"
	"finalproject/internal/request/vrequest"
	
)

func Insertvoucher(req vrequest.VoucherInsertRequest) (string, error) {
	if err := req.Validate(); err != nil {
		return "", err
	}
	voucher := mapper.VoucherMapper(&req)

	if err := insert.InsertRecord(voucher); err != nil {
		return "", err
	}
	if err := insertvoucherItem(&req.Items, voucher.ID); err != nil {
		return "", err
	}
	return "successful voucher insertion request", nil

}
func insertvoucherItem(req *[]vrequest.VoucherItemInsertion, voucherid uint) error {
	for _, item := range *req {
		voucheritem := mapper.VoucherItemMapper(&item)
		voucheritem.VoucherId = voucherid
		if err := insert.InsertRecord(*voucheritem); err != nil {
			return err
		}
	}
	
	return nil

}

func UpdateVoucher(r *vrequest.VoucherUpdateRequest) (string,error) {
    existingvoucher,err:=r.Validate()
	if err!=nil{
		return "",err
	}
	if existingvoucher.Version != r.Voucher.Version {
		return "",errors.New("version mismatch: record has been updated by another transaction")
	}
	existingvoucher2 := *mapper.VoucherMapperUpdate(r,existingvoucher)

	if err := update.Update(existingvoucher2); err != nil {
		return "",err
	}
	if err := updateVoucherItem(r); err != nil {
		return "",err
	}
	return "successful voucher update request",nil
}
func updateVoucherItem(r *vrequest.VoucherUpdateRequest) error {
	if err := insertvoucherItem(&r.Items.Inserted, r.Voucher.ID); err != nil {
		return err
	}
	if err:=updateItem(r);err!=nil{
		return err
	}
	if err:=deleteVoucherItem(r);err!=nil{
		return err
	}
	return nil

}
func updateItem(r *vrequest.VoucherUpdateRequest) error {
	for _, item := range r.Items.Updated {
		var existingvoucheritem models.VoucherItem
		if err := get.GetRecordByID(item.ID, &existingvoucheritem); err != nil {
			return err
		}
		existingvoucheritem = *mapper.VoucheItemMapperUpdate(&item, &existingvoucheritem)
		if err := update.Update(existingvoucheritem); err != nil {
			return err
		}
	}
	return nil
}

func GetVoucherByID(id uint) (*dto.Voucherdto, error) {
	var voucher models.Voucher
	if err := get.GetRecordByID(id, &voucher); err != nil {
		return nil, err
	}
	voucherItems, err := get.GetVoucherItemsByVoucherID(id)
	if err != nil {
		return nil, err
	}
	result := dto.Voucherdto{Voucher: voucher, Items: *voucherItems}
	return &result, nil
}

func DeleteVoucherWithVersion(req vrequest.VoucherDeleteRequest) (string,error) {
	var existingvoucher models.Voucher
	if err := get.GetRecordByID(req.ID, &existingvoucher); err != nil {
		return "",err
	}
	if existingvoucher.Version != req.Version {
		return "",errors.New("version mismatch: cannot delete, record has been modified by another transaction")
	}
	if err := delete.DeleteVoucherItemsByVoucherID(req.ID); err != nil {
		return "",err
	}
	if err := delete.DeleteRecord[models.Voucher](req.ID); err != nil {
		return "",err
	}
	
	return "successful voucher record delete request",nil
}
func deleteVoucherItem(r *vrequest.VoucherUpdateRequest) error {
	for _, item := range r.Items.Deleted {
		if err := delete.DeleteVoucherItemsByIdAndVoucherID(r.Voucher.ID, item); err != nil {
			return err
		}
	}

	return nil
}
