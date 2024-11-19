package validation

import (
	"errors"
	"finalproject/db"
	"finalproject/db/queries/count"
)

func LenghtValidation(inp *string) error{
	if len(*inp) == 0 || len(*inp) > 64 {
		return  errors.New("format of code/title/number is not valid,it should not be empty and more than 64 characters")
	}
	return nil

}
//&models.DL{}
func CodeDuclicateValidation(code string,table  interface{})error{
	db.ConnectWithGORM()
	var counter int64
	counter, err := count.CountByCode(db.DB, code,table)
	if err != nil {
		return err
	}
	if counter > 0 {
		return  errors.New("duplicated code")
	}
	return nil

}
func CodeDuclicateValidationUpdate(code string,id uint,table  interface{})error{
	db.ConnectWithGORM()
	var counter int64
	counter, err := count.CountByCodeExcludingID(db.DB, code,id,table)
	if err != nil {
		return err
	}
	if counter > 0 {
		return  errors.New("duplicated code")
	}
	return nil

}
func TitleDuclicateValidationUpdate(title string,id uint,table  interface{})error{
	db.ConnectWithGORM()
	var counter int64
	counter, err := count.CountByTitleExcludingID(db.DB, title,id,table)
	if err != nil {
		return err
	}
	if counter > 0 {
		return  errors.New("duplicated title")
	}
	return nil

}
func TitleDuclicateValidation(title string,table  interface{})error{
	db.ConnectWithGORM()
	var counter int64
	counter, err := count.CountByTitle(db.DB, title,table)
	if err != nil {
		return err
	}
	if counter > 0 {
		return  errors.New("duplicated title")
	}
	return nil

}
func InsertValidatin(code *string,title *string,table  interface{})  error {

	if err:=LenghtValidation(code);err!=nil{
		return err
	}
	if err:=LenghtValidation(title);err!=nil{
		return err
	}
	if err:=CodeDuclicateValidation(*code,table);err!=nil{
		return err
	}
	if err:=TitleDuclicateValidation(*title,table);err!=nil{
		return err
	}
	return nil

}
func UpdateValidatin(code *string,title *string,id *uint,table  interface{}) ( error) {

	if err:=LenghtValidation(code);err!=nil{
		return err
	}
	if err:=LenghtValidation(title);err!=nil{
		return err
	}
	if err:=CodeDuclicateValidationUpdate(*code,*id,table);err!=nil{
		return err
	}
	if err:=TitleDuclicateValidationUpdate(*title,*id,table);err!=nil{
		return err
	}
	return nil

}
