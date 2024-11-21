package services

import (
	"errors"
	delete "finalproject/db/queries/deletee"
	"finalproject/db/queries/get"
	"finalproject/db/queries/insert"
	"finalproject/db/queries/update"
	"finalproject/internal/mapper"
	"finalproject/internal/models"
	"finalproject/internal/request/dlrequest"
	"finalproject/internal/validation"
)

func Insertdl(req dlrequest.DlInsertRequest) (string, error) {
	if err := validation.SizeValidatin(&req.Code, &req.Title); err != nil {
		return "", err
	}
	dl := mapper.DlMapperInsert(&req)
	if err := insert.InsertRecord(*dl); err != nil {
		return "", err
	}
	return "successful dl insertion request", nil

}
func Updatedl(req dlrequest.DlUpdateRequest) (string, error) {
	if err := validation.SizeValidatin(&req.Code, &req.Title); err != nil {
		return "", err
	}
	var existingdl models.DL
	if err := get.GetRecordByID(req.ID, &existingdl); err != nil {
		return "", err
	}
	if existingdl.Version != req.Version {
		return "", errors.New("version mismatch: record has been updated by another transaction")
	}
	existingdl = *mapper.DlMapperUpdate(&req, &existingdl)

	if err := update.Update(existingdl); err != nil {
		return "", err
	}
	return "successful dl update request", nil

}
func GetDLByID(id uint) (*models.DL, error) {
	var dl models.DL
	if err := get.GetRecordByID(id, &dl); err != nil {
		return nil, err
	}
	return &dl, nil
}
func DeleteDLWithVersion(req dlrequest.DlDeleteRequest) (string, error) {
	var existingdl models.DL
	if err := get.GetRecordByID(req.ID, &existingdl); err != nil {
		return "", err
	}
	if existingdl.Version != req.Version {
		return "", errors.New("version mismatch: cannot delete, record has been modified by another transaction")
	}
	if err := delete.DeleteRecord[models.DL](existingdl.ID); err != nil {
		return "", err
	}

	return "successful dl record delete request", nil
}
