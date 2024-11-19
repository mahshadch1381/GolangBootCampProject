package services

import (
	"errors"
	"finalproject/db/queries/get"
	"finalproject/db/queries/deletee"
	"finalproject/db/queries/insert"
	"finalproject/db/queries/update"
	"finalproject/internal/mapper"
	"finalproject/internal/models"
	"finalproject/internal/request/dlrequest"
	"finalproject/internal/validation"
)

func Insertdl(req dlrequest.DlInsertRequest) (string, error) {
	if err := validation.InsertValidatin(&req.Code, &req.Title, &models.DL{}); err != nil {
		return "", err
	}
	dl := mapper.DlMapperInsert(&req)
	if err := insert.InsertRecord(*dl); err != nil {
		return "", err
	}
	return "save dl with id:", nil

}
func Updatedl(req dlrequest.DlUpdateRequest) (string, error) {
	if err := validation.UpdateValidatin(&req.Code, &req.Title, &req.ID, &models.DL{}); err != nil {
		return "", err
	}
	var existingdl models.DL
	if err := get.GetRecordByID(req.ID, &existingdl); err != nil {
		return "", err
	}
	existingdl = *mapper.DlMapperUpdate(&req, &existingdl)

	if err := update.Update(existingdl); err != nil {
		return "cant to update dl t bcz :", err
	}
	return "update dl with id:", nil

}
func GetDLByID(id uint) (*models.DL, error) {
	var dl models.DL
	if err := get.GetRecordByID(id, &dl); err != nil {
		return nil, err
	}
	return &dl, nil
}
func DeleteDLWithVersion(req dlrequest.DlDeleteRequest) error {
	if err:=req.Validate();err!=nil{
		return err
	}
	var existingdl models.DL
	if err := get.GetRecordByID(req.ID, &existingdl); err != nil {
		return err
	}
	if existingdl.Version != req.Version {
		return errors.New("version mismatch: cannot delete, record has been modified by another transaction")
	}
	if err := delete.DeleteRecord[models.DL](existingdl.ID); err != nil {
		return err
	}
	println("Record deleted successfully")
	return nil
}
