package mapper

import (
	"finalproject/internal/models"
	"finalproject/internal/request/slrequest"
)

func SlMapperInsert(req *slrequest.SlInsertRequest) *models.SL {
	item := &models.SL{Code: req.Code,
		Title:        req.Title,
		Version:      1,
		IsDetailable: req.IsDetailable,
	}
	return item

}
func SlMapperUpdate(req *slrequest.SlUpdateRequest, existingsl *models.SL) *models.SL {

	existingsl.Code = req.Code
	existingsl.ID = req.ID
	existingsl.Title = req.Title
	existingsl.IsDetailable = req.IsDetailable
	existingsl.Version=req.Version+1
	return existingsl
}
