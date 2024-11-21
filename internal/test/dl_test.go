package test

import (
	"finalproject/db"
	"finalproject/db/queries/get"
	"finalproject/internal/request/dlrequest"
	"finalproject/internal/services"
	"finalproject/internal/test/helper"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDl(t *testing.T) {
	db.InitDB()
	helper.Init()
	bigsizestr := helper.GenerateStringWithBigSize()
	t.Run("DL Creation", func(t *testing.T) {

		t.Run("should not allow duplicate code", func(t *testing.T) {
			code1 := helper.GenerateNumericString()
			title1 := helper.GenerateRandomString()
			title2 := helper.GenerateRandomString()

			str1, err1 := services.Insertdl(dlrequest.DlInsertRequest{Code: code1, Title: title1})
			str2, err2 := services.Insertdl(dlrequest.DlInsertRequest{Code: code1, Title: title2})

			assert.Nil(t, err1, "expected first creation to succeed")
			assert.NotEmpty(t, str1)
			assert.NotNil(t, err2, "expected error for duplicate code")
			assert.Empty(t, str2)
		})

	})
	t.Run("should not allow duplicate title", func(t *testing.T) {
		code2 := helper.GenerateNumericString()
		code3 := helper.GenerateNumericString()
		title3 := helper.GenerateRandomString()
		str1, err1 := services.Insertdl(dlrequest.DlInsertRequest{Code: code2, Title: title3})
		_, err2 := services.Insertdl(dlrequest.DlInsertRequest{Code: code3, Title: title3})

		assert.Nil(t, err1, "expected first creation to succeed")
		assert.Equal(t, str1, "successful dl insertion request")
		assert.NotNil(t, err2, "expected error for duplicate title")

	})

	t.Run("should not allow empty code", func(t *testing.T) {
		title3 := helper.GenerateRandomString()
		_, err1 := services.Insertdl(dlrequest.DlInsertRequest{Code: "", Title: title3})
		assert.NotNil(t, err1, "expected error for empty code")
	})

	t.Run("should not allow empty title", func(t *testing.T) {
		code := helper.GenerateNumericString()
		_, err1 := services.Insertdl(dlrequest.DlInsertRequest{Code: code, Title: ""})
		assert.NotNil(t, err1, "expected error for empty title")
	})
	t.Run("should not allow big size(more than 64 char) code", func(t *testing.T) {

		_, err1 := services.Insertdl(dlrequest.DlInsertRequest{Code: bigsizestr, Title: ""})
		assert.NotNil(t, err1, "expected error for big size code")
	})
	t.Run("should not allow big size(more than 64 char) title", func(t *testing.T) {
		code := helper.GenerateNumericString()
		_, err1 := services.Insertdl(dlrequest.DlInsertRequest{Code: code, Title: bigsizestr})
		assert.NotNil(t, err1, "expected error for big size title")
	})

}
func TestGetDl(t *testing.T) {
	t.Run("DL Getting", func(t *testing.T) {
		t.Run("should not allow get dl that doesnt exist", func(t *testing.T) {
			str, err := services.GetDLByID(2000)
			fmt.Println(err.Error())
			assert.NotNil(t, err, "expected error for duplicate code")
			assert.Empty(t, str)
		})

		t.Run("get dl successfully", func(t *testing.T) {
			dl, err := services.GetDLByID(1)
			fmt.Printf("this is dl you want: dl.code=%v, dl.title=%v, dl.version=%v \n", dl.Code, dl.Title, dl.Version)
			assert.Nil(t, err, "expected error for duplicate code")
			assert.NotEmpty(t, dl)
		})

	})

}
func TestDeleteDl(t *testing.T) {
	t.Run("DL deleting", func(t *testing.T) {
		t.Run("should not allow delete dl that doesnt exist", func(t *testing.T) {
			req := dlrequest.DlDeleteRequest{ID: 20000, Version: 1}
			str, err := services.DeleteDLWithVersion(req)
			fmt.Println(err.Error())
			assert.NotNil(t, err, "expected error for not found")
			assert.Empty(t, str)
		})
		t.Run("delete dl successfully", func(t *testing.T) {
			code1 := helper.GenerateNumericString()
			title1 := helper.GenerateRandomString()
			_, err1 := services.Insertdl(dlrequest.DlInsertRequest{Code: code1, Title: title1})
			assert.Nil(t, err1)
			id, _ := get.GetLastID("dl")
			req := dlrequest.DlDeleteRequest{ID: id, Version: 1}
			str, err := services.DeleteDLWithVersion(req)
			fmt.Println(str)
			assert.Nil(t, err)
			assert.NotEmpty(t, str)
		})

	})

}
func TestUpdateDl(t *testing.T) {
	db.InitDB()
	t.Run("DL Update", func(t *testing.T) {
		code1 := helper.GenerateNumericString()
		code2 := helper.GenerateNumericString() + "0"
		title1 := helper.GenerateRandomString()
		title2 := helper.GenerateRandomString() + "L"

		str1, err1 := services.Insertdl(dlrequest.DlInsertRequest{Code: code1, Title: title1})
		str2, err2 := services.Insertdl(dlrequest.DlInsertRequest{Code: code2, Title: title2})

		assert.Nil(t, err1, "expected first creation to succeed")
		assert.NotEmpty(t, str1)
		assert.Nil(t, err2, "expected error for duplicate code")
		assert.NotEmpty(t, str2)
		id1, err := get.GetIDByCode("dl", code1)
		assert.Nil(t, err)
		t.Run("should not allow Update with  duplicate code", func(t *testing.T) {
			req := dlrequest.DlUpdateRequest{ID: id1, Code: code2, Title: title1, Version: 1}
			str, err := services.Updatedl(req)
			fmt.Println(err.Error())
			assert.NotNil(t, err)
			assert.Empty(t, str)

		})
		t.Run("should not allow Update with  empty code", func(t *testing.T) {
			req := dlrequest.DlUpdateRequest{ID: id1, Code: "", Title: title1, Version: 1}
			str, err := services.Updatedl(req)
			fmt.Println(err.Error())
			assert.NotNil(t, err)
			assert.Empty(t, str)

		})
		t.Run("should not allow Update with  empty title", func(t *testing.T) {
			req := dlrequest.DlUpdateRequest{ID: id1, Code: code2, Title: "", Version: 1}
			str, err := services.Updatedl(req)
			fmt.Println(err.Error())
			assert.NotNil(t, err)
			assert.Empty(t, str)

		})
		t.Run("should not  allow Update with duplicate title", func(t *testing.T) {
			req := dlrequest.DlUpdateRequest{ID: id1, Code: code1, Title: title2, Version: 1}
			str, err := services.Updatedl(req)
			fmt.Println(err.Error())
			assert.NotNil(t, err)
			assert.Empty(t, str)

		})

		t.Run("should not allow update when mismatch version", func(t *testing.T) {
			code3 := helper.GenerateNumericString()
			req := dlrequest.DlUpdateRequest{ID: id1, Code: code3, Title: title1, Version: 1}
			str, err := services.Updatedl(req)
			fmt.Println(str)
			assert.Nil(t, err)
			assert.NotEmpty(t, str)

			title3 := helper.GenerateRandomString()
			req = dlrequest.DlUpdateRequest{ID: id1, Code: code1, Title: title3, Version: 1}
			str, err = services.Updatedl(req)
			fmt.Println(err.Error())
			assert.NotNil(t, err)
			assert.Empty(t, str)

		})

	})

}

func TestDeleteDlWithDetail(t *testing.T) {

	t.Run("DL deleting versions", func(t *testing.T) {
		t.Run("should not allow delete dl that version isnt match", func(t *testing.T) {
			fmt.Println("kkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk")
			code1 := helper.GenerateNumericString() + "3"
			title1 := helper.GenerateRandomString()
			str1, err1 := services.Insertdl(dlrequest.DlInsertRequest{Code: code1, Title: title1})
			assert.Nil(t, err1, "expected first creation to succeed")
			assert.NotEmpty(t, str1)
			id, _ := get.GetLastID("dl")
			req := dlrequest.DlUpdateRequest{ID: id, Code: code1, Title: title1, Version: 1}
			str, err := services.Updatedl(req)
			assert.Nil(t, err, "expected first creation to succeed")
			assert.NotEmpty(t, str)
			println(id)
			req2 := dlrequest.DlDeleteRequest{ID: id, Version: 1}
			str, err = services.DeleteDLWithVersion(req2)
			assert.NotNil(t, err, "expected error for mismatche version")
			assert.Empty(t, str)
		})

	})

}
