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
	err := validation.SizeValidatin(&req.Code, &req.Title)
	if err != nil {
		return "", err
	}
	sl := mapper.SlMapperInsert(&req)
	if err := insert.InsertRecord(&sl); err != nil {
		return "", err
	}
	return "successful sl insertion request", nil

}
func Updatesl(req slrequest.SlUpdateRequest) (string, error) {
	if err := req.Validate(); err != nil {
		return "", err
	}
	err := validation.SizeValidatin(&req.Code, &req.Title)
	if err != nil {
		return "", err
	}
	var existingsl models.SL
	if err := get.GetRecordByID(req.ID, &existingsl); err != nil {
		return "", err
	}
	if existingsl.Version != req.Version {
		return "",errors.New("version mismatch: record has been updated by another transaction")
	}
	existingsl = *mapper.SlMapperUpdate(&req, &existingsl)
	if err := update.Update(existingsl); err != nil {
		return "", err
	}
	return "successful sl update request", nil

}
func GetSLByID(id uint) (*models.SL, error) {
	var sl models.SL
	if err := get.GetRecordByID(id, &sl); err != nil {
		return nil, err
	}
	return &sl, nil
}
func DeleteSLWithVersion(req slrequest.SlDeleteRequest) (string,error) {

	var existingsl models.SL
	if err := get.GetRecordByID(req.ID, &existingsl); err != nil {
		return "",err
	}
	if existingsl.Version != req.Version {
		return "",errors.New("version mismatch: cannot delete, record has been modified by another transaction")
	}
	if err := delete.DeleteRecord[models.SL](existingsl.ID); err != nil {
		return "",err
	}
	
	return "successful sl recored delete request",nil
}
