package test

import (
	"finalproject/internal/test/helper"
	"finalproject/internal/request/slrequest"
	"finalproject/internal/services"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"finalproject/db/queries/get"
)

func TestCreateSl(t *testing.T) {
	t.Run("SL Creation", func(t *testing.T) {
		bigsizestr := helper.GenerateStringWithBigSize()
		t.Run("should not allow duplicate code", func(t *testing.T) {
            code1:=helper.GenerateNumericString()
			title1:=helper.GenerateRandomString()
			title2:=helper.GenerateRandomString()
			str1, err1 := services.Insertsl(slrequest.SlInsertRequest{Code: code1, Title: title1, IsDetailable: false})
			str2, err2 := services.Insertsl(slrequest.SlInsertRequest{Code: code1, Title: title2, IsDetailable: false})

			assert.Nil(t, err1, "expected first creation to succeed")
			assert.NotEmpty(t, str1)
			assert.NotNil(t, err2, "expected error for duplicate code")
			assert.Empty(t, str2)
		})

		t.Run("should not allow duplicate title", func(t *testing.T) {
			code2:=helper.GenerateNumericString()
			code3:=helper.GenerateNumericString()+"0"
			title3:=helper.GenerateRandomString()
			str1, err1 := services.Insertsl(slrequest.SlInsertRequest{Code: code2, Title: title3, IsDetailable: true})
			_, err2 := services.Insertsl(slrequest.SlInsertRequest{Code: code3, Title: title3, IsDetailable: false})

			assert.Nil(t, err1, "expected first creation to succeed")
			assert.Equal(t, str1, "successful sl insertion request")
			assert.NotNil(t, err2, "expected error for duplicate title")
			fmt.Println("t12 : "+err2.Error())

		})

		t.Run("should not allow empty code", func(t *testing.T) {
			title3:=helper.GenerateRandomString()
			_, err1 := services.Insertsl(slrequest.SlInsertRequest{Code: "", Title: title3, IsDetailable: true})
			assert.NotNil(t, err1, "expected error for empty code")
			fmt.Println("t13 : "+err1.Error())
		})

		t.Run("should not allow empty title", func(t *testing.T) {
			code4:=helper.GenerateNumericString()
			_, err1 := services.Insertsl(slrequest.SlInsertRequest{Code: code4, Title: "", IsDetailable: true})
			assert.NotNil(t, err1, "expected error for empty code")
			fmt.Println("t14 : "+err1.Error())
		})
		t.Run("should not allow big size(more than 64 char) code", func(t *testing.T) {

			_, err1 := services.Insertsl(slrequest.SlInsertRequest{Code: bigsizestr, Title: ""})
			assert.NotNil(t, err1, "expected error for big size code")
			fmt.Println("t15 : "+err1.Error())
		})
		t.Run("should not allow big size(more than 64 char) title", func(t *testing.T) {
			code := helper.GenerateNumericString()
			_, err1 := services.Insertsl(slrequest.SlInsertRequest{Code: code, Title: bigsizestr})
			assert.NotNil(t, err1, "expected error for big size title")
		})
	})
}
func TestGetSl(t *testing.T) {
	t.Run("SL Getting", func(t *testing.T) {
		t.Run("should not allow get sl that doesnt exist", func(t *testing.T) {
			sl, err := services.GetSLByID(20000)
			fmt.Println("t16 : "+err.Error())
			assert.NotNil(t, err, "expected error for duplicate code")
			assert.Empty(t, sl)
		})

		t.Run("get sl successfully", func(t *testing.T) {
			sl, err := services.GetSLByID(1)
			fmt.Printf("t17 : "+"this is sl you want: sl.code=%v, sl.title=%v, sl.version=%v, sl.is_detailed=%v \n", sl.Code, sl.Title, sl.Version,sl.IsDetailable)
			assert.Nil(t, err, "expected error for duplicate code")
			assert.NotEmpty(t, sl)
		})

	})
}
func TestDeleteSl(t *testing.T) {
	t.Run("DL deleting", func(t *testing.T) {
		t.Run("should not allow delete sl that doesnt exist", func(t *testing.T) {
            req:=slrequest.SlDeleteRequest{ID: 2000,Version :1}
			str, err := services.DeleteSLWithVersion(req)
			fmt.Println("t18 : "+err.Error())
			assert.NotNil(t, err, "expected error for not found")
			assert.Empty(t, str)
		})
		t.Run("delete sl successfully", func(t *testing.T) {
			code1:=helper.GenerateNumericString()+"2"
			title1:=helper.GenerateRandomString()+"j"
			_, err1 := services.Insertsl(slrequest.SlInsertRequest{Code: code1, Title: title1, IsDetailable: false})
            assert.Nil(t,err1)
			id,_:=get.GetLastID("sl")
            req:=slrequest.SlDeleteRequest{ID: id,Version :1}
			str, err := services.DeleteSLWithVersion(req)
			fmt.Println("t19 : "+str)
			assert.Nil(t, err)
			assert.NotEmpty(t, str)
		})

	})

}
func TestUpdateSl(t *testing.T) {

	t.Run("SL Update", func(t *testing.T) {
		    code1:=helper.GenerateNumericString()
			code2:=helper.GenerateNumericString()+"0"
			title1:=helper.GenerateRandomString()
			title2:=helper.GenerateRandomString()+"L"

			str1, err1 := services.Insertsl(slrequest.SlInsertRequest{Code: code1, Title: title1,IsDetailable: true})
			str2, err2 := services.Insertsl(slrequest.SlInsertRequest{Code: code2, Title: title2,IsDetailable: false})

			assert.Nil(t, err1, "expected first creation to succeed")
			assert.NotEmpty(t, str1)
			assert.Nil(t, err2,  "expected second creation to succeed")
			assert.NotEmpty(t, str2)
			id1,err:=get.GetIDByCode("sl",code1)
			assert.Nil(t,err)
		t.Run("should not allow Update with  duplicate code", func(t *testing.T) {
			req:=slrequest.SlUpdateRequest{ID: id1, Code: code2, Title: title1, Version: 1,IsDetailable: true}
            str,err:=services.Updatesl(req)
			fmt.Println("t20 : "+err.Error())
			assert.NotNil(t,err)
			assert.Empty(t,str)

		})
		t.Run("should not allow Update with  empty code", func(t *testing.T) {
			req:=slrequest.SlUpdateRequest{ID: id1, Code: "", Title: title1, Version: 1,IsDetailable: true}
            str,err:=services.Updatesl(req)
			fmt.Println("t21 : "+err.Error())
			assert.NotNil(t,err)
			assert.Empty(t,str)

		})
		t.Run("should not allow Update with  empty title", func(t *testing.T) {
			req:=slrequest.SlUpdateRequest{ID: id1, Code: code2, Title: "", Version: 1,IsDetailable: true}
            str,err:=services.Updatesl(req)
			fmt.Println("t22 : "+err.Error())
			assert.NotNil(t,err)
			assert.Empty(t,str)

		})
		t.Run("should not  allow Update with duplicate title", func(t *testing.T) {
			req:=slrequest.SlUpdateRequest{ID: id1, Code: code1, Title: title2, Version: 1,IsDetailable: true}
            str,err:=services.Updatesl(req)
			fmt.Println("t23 : "+err.Error())
			assert.NotNil(t,err)
			assert.Empty(t,str)

		})

		t.Run("should not allow update when mismatch version", func(t *testing.T) {
			code3:=helper.GenerateNumericString()
			req:=slrequest.SlUpdateRequest{ID: id1, Code: code3, Title: title1, Version: 1}
            str,err:=services.Updatesl(req)
			fmt.Println("t24 : "+str)
			assert.Nil(t,err)
			assert.NotEmpty(t,str)

			title3:=helper.GenerateRandomString()
			req=slrequest.SlUpdateRequest{ID: id1, Code: code1, Title: title3, Version: 1}
            str,err=services.Updatesl(req)
			fmt.Println("t25 : "+err.Error())
			assert.NotNil(t,err)
			assert.Empty(t,str)

		})

	})
	

}