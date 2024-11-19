package mapper

import (
	"finalproject/internal/models"
	"finalproject/internal/request/dlrequest"
)
func DlMapperInsert(req *dlrequest.DlInsertRequest) *models.DL {
	item := &models.DL{
		                      Code: req.Code,
							  Title: req.Title, 
							   Version: 1,
							  
							}
	return item

}
func DlMapperUpdate(req *dlrequest.DlUpdateRequest,existingdl *models.DL) *models.DL {
	existingdl.Code=req.Code
	existingdl.ID=req.ID
	existingdl.Title=req.Title
	
	return existingdl
}