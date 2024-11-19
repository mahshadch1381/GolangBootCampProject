package services

import (
	"errors"
	delete "finalproject/db/queries/deletee"
	"finalproject/db/queries/get"
	"finalproject/db/queries/insert"
	"finalproject/db/queries/update"
	"finalproject/internal/mapper"
	"finalproject/internal/models"
	"finalproject/internal/request/slrequest"
	"finalproject/internal/validation"
)

func Insertsl(req slrequest.SlInsertRequest) (string, error) {
	err := validation.InsertValidatin(&req.Code, &req.Title, &models.SL{})
	if err != nil {
		return "", err
	}
	sl := mapper.SlMapperInsert(&req)
	if err := insert.InsertRecord(&sl); err != nil {
		return "", err
	}
	return "save sl with id:", nil

}
func Updatesl(req slrequest.SlUpdateRequest) (string, error) {
	if err := req.Validate(); err != nil {
		return "", err
	}
	err := validation.UpdateValidatin(&req.Code, &req.Title, &req.ID, &models.SL{})
	if err != nil {
		return "", err
	}
	var existingsl models.SL
	if err := get.GetRecordByID(req.ID, &existingsl); err != nil {
		return "", err
	}
	existingsl = *mapper.SlMapperUpdate(&req, &existingsl)
	if err := update.Update(existingsl); err != nil {
		return "cant to update sl  bcz :", err
	}
	return "update sl with id:", nil

}
func GetSLByID(id uint) (*models.SL, error) {
	var sl models.SL
	if err := get.GetRecordByID(id, &sl); err != nil {
		return nil, err
	}
	return &sl, nil
}
func DeleteSLWithVersion(req slrequest.SlDeleteRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	var existingsl models.SL
	if err := get.GetRecordByID(req.ID, &existingsl); err != nil {
		return err
	}
	if existingsl.Version != req.Version {
		return errors.New("version mismatch: cannot delete, record has been modified by another transaction")
	}
	if err := delete.DeleteRecord[models.SL](existingsl.ID); err != nil {
		return err
	}
	println("Record deleted successfully from sl table")
	return nil
}
