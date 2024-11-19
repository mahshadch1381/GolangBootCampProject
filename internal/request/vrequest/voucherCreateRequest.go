package vrequest

import (
	"errors"
	"finalproject/db/queries/count"
	"finalproject/internal/models"
	"finalproject/internal/validation"
)

type VoucherInsertion struct {
	Number  string
	Version string
}

type VoucherItemInsertion struct {
	VoucherID uint
	SLID      uint  `json:"sl_id"`
	DLID      *uint `json:"dl_id"`
	Debit     int
	Credit    int
}

type VoucherInsertRequest struct {
	Voucher VoucherInsertion
	Items   []VoucherItemInsertion
}

// validations

func (r *VoucherInsertRequest) Validate() error {
	if err := validation.LenghtValidation(&r.Voucher.Number); err != nil {
		return err
	}
	if err := r.NumberDuclicateValidation(); err != nil {
		return err
	}
	if err := r.ValidateCountOfItems(); err != nil {
		return err
	}
	if err := r.ValidateDebitsAndCredits(); err != nil {
		return err
	}
	if err := r.ValidateSumOfDebitsAndCredits(); err != nil {
		return err
	}
	return nil
}

func (r *VoucherInsertRequest) NumberDuclicateValidation() error {
	var counter int64
	counter, err := count.CountByNumber(r.Voucher.Number, &models.Voucher{})
	if err != nil {
		return err
	}
	if counter > 0 {
		return errors.New("duplicated code")
	}
	return nil

}

func (r *VoucherInsertRequest) ValidateCountOfItems() error {
	itemCount := len(r.Items)

	if itemCount < 2 {
		return errors.New("the voucher must have at least 2 items")
	}
	if itemCount > 500 {
		return errors.New("the voucher cannot have more than 500 items")
	}
	return nil
}

func (r *VoucherInsertRequest) ValidateDebitsAndCredits() error {
	for _, item := range r.Items {
		if !((item.Debit > 0 && item.Credit == 0) || (item.Debit == 0 && item.Credit > 0)) {
			return errors.New("each item must have either debit or credit greater than 0, and the other 0 ")
		}
	}
	return nil
}

func (r *VoucherInsertRequest) ValidateSumOfDebitsAndCredits() error {
	var debits, credits int
	for _, item := range r.Items {
		debits += item.Debit
		credits += item.Credit
	}
	if credits != debits {
		return errors.New("sum of credits and debits are not 0 ")
	}
	return nil
}
