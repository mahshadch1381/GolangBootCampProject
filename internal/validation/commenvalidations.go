package validation

import (
	"errors"
	
)
func LenghtValidation(inp *string) error{
	if len(*inp) == 0 || len(*inp) > 64 {
		return  errors.New("format of code/title/number is not valid,it should not be empty and more than 64 characters")
	}
	return nil

}
func SizeValidatin(code *string,title *string)  error {

	if err:=LenghtValidation(code);err!=nil{
		return err
	}
	if err:=LenghtValidation(title);err!=nil{
		return err
	}
	return nil

}
