package test

import (
	"finalproject/db/queries/get"
	"finalproject/internal/request/dlrequest"
	"finalproject/internal/request/slrequest"
	"finalproject/internal/request/vrequest"
	"finalproject/internal/services"
	"fmt"
	"testing"
    "finalproject/internal/models"
	"github.com/stretchr/testify/assert"
	"finalproject/internal/test/helper"
)

func TestCreateVoucher(t *testing.T) {
     var id_dl uint
	 var id_sl_withoutdl uint
	 var id_sl_withdl uint
	t.Run("Voucer Creation", func(t *testing.T) {
		t.Run("should create needed objects", func(t *testing.T) {
			code1 := helper.GenerateNumericString()
			title1 := helper.GenerateRandomString()
			code2 := helper.GenerateNumericString()+"109"
			title2 := helper.GenerateRandomString()
			_, err1 := services.Insertdl(dlrequest.DlInsertRequest{Code: code1, Title: title1})
			assert.Nil(t, err1)
			_, err2 := services.Insertsl(slrequest.SlInsertRequest{Code: code2, Title: title2, IsDetailable: false})
			assert.Nil(t, err2)

			id_dl1, err := get.GetIDByCode("dl", code1)
			assert.Nil(t, err)
			id_dl=id_dl1
			id_sl_withoutdl1, err := get.GetIDByCode("sl", code2)
			assert.Nil(t, err)
			id_sl_withoutdl=id_sl_withoutdl1
		})
		t.Run("should not allow duplicate number", func(t *testing.T) {

			number := helper.GenerateNumericString()
			request1 := vrequest.VoucherInsertRequest{
				Voucher: vrequest.VoucherInsertion{Number: number},
				Items: []vrequest.VoucherItemInsertion{{SLID: id_sl_withoutdl, DLID: nil, Debit: 200, Credit: 0},
					{SLID: id_sl_withoutdl, DLID: nil, Debit: 0, Credit: 200}},
			}
			request2 := vrequest.VoucherInsertRequest{
				Voucher: vrequest.VoucherInsertion{Number: number},
				Items: []vrequest.VoucherItemInsertion{{SLID: id_sl_withoutdl, DLID: nil, Debit: 200, Credit: 0},
					{SLID: id_sl_withoutdl, DLID: nil, Debit: 0, Credit: 200}},
			}
			str1, err1 := services.Insertvoucher(request1)
			str2, err2 := services.Insertvoucher(request2)
			assert.Nil(t, err1, "expected first creation to succeed")
			assert.Equal(t, str1, "successful voucher insertion request")
			assert.NotNil(t, err2, "expected error for duplicate number")
			assert.Empty(t, str2)
		})
		t.Run("should not allow empty number", func(t *testing.T) {
			request2 := vrequest.VoucherInsertRequest{
				Voucher: vrequest.VoucherInsertion{Number: ""},
				Items: []vrequest.VoucherItemInsertion{{SLID: id_sl_withoutdl, DLID: nil, Debit: 200, Credit: 0},
					{SLID: id_sl_withoutdl, DLID: nil, Debit: 0, Credit: 200}},
			}
			str2, err2 := services.Insertvoucher(request2)
			assert.NotNil(t, err2, "expected error for empty number")
			assert.Empty(t, str2)
		})
		t.Run("should not allow empty number", func(t *testing.T) {
			number:=helper.GenerateStringWithBigSize()
			request2 := vrequest.VoucherInsertRequest{
				Voucher: vrequest.VoucherInsertion{Number: number},
				Items: []vrequest.VoucherItemInsertion{{SLID: id_sl_withoutdl, DLID: nil, Debit: 200, Credit: 0},
					{SLID: id_sl_withoutdl, DLID: nil, Debit: 0, Credit: 200}},
			}
			str2, err2 := services.Insertvoucher(request2)
			assert.NotNil(t, err2, "expected error for big size number")
			assert.Empty(t, str2)
		})
		t.Run("should not allow less than 2 items in voucheritem list", func(t *testing.T) {
			request1 := vrequest.VoucherInsertRequest{
				Voucher: vrequest.VoucherInsertion{Number: "10"},
				Items:   []vrequest.VoucherItemInsertion{{SLID: id_sl_withoutdl, DLID: nil, Debit: 200, Credit: 0}},
			}
			str1, err1 := services.Insertvoucher(request1)
			fmt.Println(err1.Error())
			assert.NotNil(t, err1, "expected error for duplicate number")
			assert.Empty(t, str1)
		})

		t.Run("should not allow both credit and debit field in voucher item more than 0", func(t *testing.T) {
			request1 := vrequest.VoucherInsertRequest{
				Voucher: vrequest.VoucherInsertion{Number: "1100"},
				Items: []vrequest.VoucherItemInsertion{{SLID: id_sl_withoutdl, DLID: nil, Debit: 200, Credit: 100},
					{SLID: id_sl_withoutdl, DLID: nil, Debit: 0, Credit: 200}},
			}
			str1, err1 := services.Insertvoucher(request1)
			fmt.Println(err1.Error())
			assert.NotNil(t, err1, "expected error for duplicate number")
			assert.Empty(t, str1)
		})

		t.Run("should not allow sum of debita and credits are not 0", func(t *testing.T) {
			request1 := vrequest.VoucherInsertRequest{
				Voucher: vrequest.VoucherInsertion{Number: "1101"},
				Items: []vrequest.VoucherItemInsertion{{SLID: id_sl_withoutdl, DLID: nil, Debit: 400, Credit: 0},
					{SLID: id_sl_withoutdl, DLID: nil, Debit: 0, Credit: 100},
					{SLID: id_sl_withoutdl, DLID: nil, Debit: 0, Credit: 200}},
			}
			str1, err1 := services.Insertvoucher(request1)
			fmt.Println(err1.Error())
			assert.NotNil(t, err1, "expected error for duplicate number")
			assert.Empty(t, str1)
		})

		t.Run("should not allow sl field in voucher item be empty", func(t *testing.T) {
			var a uint
			a = id_dl
			request1 := vrequest.VoucherInsertRequest{
				Voucher: vrequest.VoucherInsertion{Number: "1102"},
				Items: []vrequest.VoucherItemInsertion{{DLID: &a, Debit: 400, Credit: 0},
					{SLID: id_sl_withoutdl, DLID: nil, Debit: 0, Credit: 200},
					{SLID: id_sl_withoutdl, DLID: nil, Debit: 0, Credit: 200}},
			}
			str1, err1 := services.Insertvoucher(request1)
			fmt.Println(err1.Error())
			assert.NotNil(t, err1, "expected error for duplicate number")
			assert.Empty(t, str1)
		})

		t.Run("should not allow voucher item with sl that is detailable,has not dl field ", func(t *testing.T) {
			code3 := helper.GenerateNumericString()+"12"
			title3 := helper.GenerateRandomString()+"po"
			_, err3 := services.Insertsl(slrequest.SlInsertRequest{Code: code3, Title: title3, IsDetailable: true})
			assert.Nil(t, err3)
			id_sl_dl, err := get.GetIDByCode("sl", code3)
			//	id_sl_withdl = id_sl_dl
			assert.Nil(t, err)
			request1 := vrequest.VoucherInsertRequest{
				Voucher: vrequest.VoucherInsertion{Number: "1103"},
				Items: []vrequest.VoucherItemInsertion{{SLID: id_sl_dl, Debit: 400, Credit: 0},
					{SLID: 1, DLID: nil, Debit: 0, Credit: 200},
					{SLID: 1, DLID: nil, Debit: 0, Credit: 200}},
			}
			str1, err1 := services.Insertvoucher(request1)
			fmt.Println(err1.Error())
			assert.NotNil(t, err1, "expected error for duplicate number")
			assert.Empty(t, str1)
		})
		t.Run("should not allow voucher item that its sl is not detailable has dl field ", func(t *testing.T) {
			var a uint
			a = id_dl
			request1 := vrequest.VoucherInsertRequest{
				Voucher: vrequest.VoucherInsertion{Number: "1104"},
				Items: []vrequest.VoucherItemInsertion{{SLID: id_sl_withoutdl, DLID: &a, Debit: 400, Credit: 0},
					{SLID: id_sl_withoutdl, DLID: &a, Debit: 0, Credit: 200},
					{SLID: id_sl_withoutdl, DLID: nil, Debit: 0, Credit: 200}},
			}
			str1, err1 := services.Insertvoucher(request1)
			fmt.Println(err1.Error())
			assert.NotNil(t, err1, "expected error for duplicate number")
			assert.Empty(t, str1)
		})

		t.Run("should not allow voucher item referes to sl that doesn't exist ", func(t *testing.T) {
			code3 := helper.GenerateNumericString()
			title3 := helper.GenerateRandomString()
			_, err3 := services.Insertsl(slrequest.SlInsertRequest{Code: code3, Title: title3, IsDetailable: false})
			assert.Nil(t, err3)
			id_sl_nodl, err := get.GetIDByCode("sl", code3)
			assert.Nil(t, err)
			req := slrequest.SlDeleteRequest{ID: id_sl_nodl, Version: 1}
			_, err = services.DeleteSLWithVersion(req)
			assert.Nil(t, err)
			request1 := vrequest.VoucherInsertRequest{
				Voucher: vrequest.VoucherInsertion{Number: code3},
				Items: []vrequest.VoucherItemInsertion{
					{SLID: id_sl_nodl, DLID: nil, Debit: 200, Credit: 0},
					{SLID: id_sl_nodl, DLID: nil, Debit: 0, Credit: 200}},
			}
			str1, err1 := services.Insertvoucher(request1)
			fmt.Println(err1.Error())
			assert.NotNil(t, err1, "expected error for duplicate number")
			assert.Empty(t, str1)
		
			t.Run("should not allow voucher item referes to dl that doesn't exist ", func(t *testing.T) {
				req := dlrequest.DlDeleteRequest{ID: id_dl, Version: 1}
				_, err := services.DeleteDLWithVersion(req)
				assert.Nil(t, err)
				request1 := vrequest.VoucherInsertRequest{
					Voucher: vrequest.VoucherInsertion{Number: "1106"},
					Items: []vrequest.VoucherItemInsertion{{SLID: id_sl_withdl, DLID: &id_dl, Debit: 400, Credit: 0},
						{SLID: id_sl_withoutdl, DLID: nil, Debit: 0, Credit: 200},
						{SLID: id_sl_withoutdl, DLID: nil, Debit: 0, Credit: 200}},
				}
				str1, err1 := services.Insertvoucher(request1)
				fmt.Println(err1.Error())
				assert.NotNil(t, err1, "expected error for duplicate number")
				assert.Empty(t, str1)
			})

		})
	})
}
func TestGetVoucher(t *testing.T) {
	t.Run("Voucher Getting", func(t *testing.T) {
		t.Run("should not allow get Voucher that doesnt exist", func(t *testing.T) {
			v, err := services.GetVoucherByID(20000)
			fmt.Println(err.Error())
			assert.NotNil(t, err, "expected error for duplicate code")
			assert.Empty(t, v)
		})

		t.Run("get Voucher successfully", func(t *testing.T) {
			v, err := services.GetVoucherByID(1)
			fmt.Printf("this is Voucher you want: Voucher.Number=%v, Voucher.version=%v, Voucher.items=%v \n", v.Voucher.Number, v.Voucher.Version, v.Items)
			assert.Nil(t, err, "expected error for duplicate code")
			assert.NotEmpty(t, v)
		})

	})
}

func TestDeleteVoucher(t *testing.T) {
	t.Run("Voucher deleting", func(t *testing.T) {
		t.Run("should not allow delete voucher that doesnt exist", func(t *testing.T) {
			req := vrequest.VoucherDeleteRequest{ID: 20000, Version: 1}
			str, err := services.DeleteVoucherWithVersion(req)
			fmt.Println(err.Error())
			assert.NotNil(t, err, "expected error for not found")
			assert.Empty(t, str)
		})
		t.Run("delete sl successfully", func(t *testing.T) {
			code1 := helper.GenerateNumericString()+"0"
			title1 := helper.GenerateRandomString()+"L"
			_, err1 := services.Insertsl(slrequest.SlInsertRequest{Code: code1, Title: title1, IsDetailable: false})
			assert.Nil(t, err1)
			id_sl, err := get.GetIDByCode("sl", code1)
			assert.Nil(t, err)
			number := helper.GenerateNumericString()
			request1 := vrequest.VoucherInsertRequest{
				Voucher: vrequest.VoucherInsertion{Number: number},
				Items: []vrequest.VoucherItemInsertion{{SLID: id_sl, DLID: nil, Debit: 200, Credit: 0},
					{SLID: id_sl, DLID: nil, Debit: 0, Credit: 200}},
			}
			_, err1 = services.Insertvoucher(request1)
			assert.Nil(t, err1)
            id_v,err:=get.GetIDByNumber("voucher",number)
			assert.Nil(t, err)
		    req:=vrequest.VoucherDeleteRequest{ID: id_v,Version :1}
			str, err := services.DeleteVoucherWithVersion(req)
			fmt.Println(str)
			assert.Nil(t, err)
			assert.NotEmpty(t, str)
		})

	})

}
func TestUpdateVoucher(t *testing.T) {
	var id_dl uint
	var id_sl_withdl uint
	var id_voucher1 uint
	t.Run("SL Update", func(t *testing.T) {
		    number1:=helper.GenerateNumericString()
			number2:=helper.GenerateNumericString()+"0"

			code1 := helper.GenerateNumericString()
			title1 := helper.GenerateRandomString()

			code2 := helper.GenerateNumericString()
			title2 := helper.GenerateRandomString()
			_, err1 := services.Insertdl(dlrequest.DlInsertRequest{Code: code1, Title: title1})
			assert.Nil(t, err1)
			_, err2 := services.Insertsl(slrequest.SlInsertRequest{Code: code2, Title: title2, IsDetailable: true})
			assert.Nil(t, err2)

			id_dl1, err := get.GetIDByCode("dl", code1)
			assert.Nil(t, err)
			id_dl=id_dl1
			id_sl_withdl1, err := get.GetIDByCode("sl", code2)
			assert.Nil(t, err)
			id_sl_withdl=id_sl_withdl1
		
			request1 := vrequest.VoucherInsertRequest{
				Voucher: vrequest.VoucherInsertion{Number: number1},
				Items: []vrequest.VoucherItemInsertion{{SLID: id_sl_withdl, DLID: &id_dl, Debit: 300, Credit: 0},
					{SLID: id_sl_withdl, DLID: &id_dl, Debit: 0, Credit: 300}},
			}
			request2 := vrequest.VoucherInsertRequest{
				Voucher: vrequest.VoucherInsertion{Number: number2},
				Items: []vrequest.VoucherItemInsertion{{SLID: id_sl_withdl, DLID: &id_dl, Debit: 200, Credit: 0},
					{SLID: id_sl_withdl, DLID: &id_dl, Debit: 0, Credit: 200}},
			}
			str1, err1 := services.Insertvoucher(request1)
			str2, err2 := services.Insertvoucher(request2)
			assert.Nil(t, err1, "expected first creation to succeed")
			assert.Equal(t, str1, "successful voucher insertion request")
			assert.Nil(t, err2, "expected first creation to succeed")
			assert.Equal(t, str2, "successful voucher insertion request")
            id_v1,err:=get.GetIDByNumber("voucher",number1)
			assert.Nil(t, err)
			id_voucher1=id_v1
		t.Run("should not allow Update with  duplicate number", func(t *testing.T) {
			updateRequest := vrequest.VoucherUpdateRequest{
				Voucher: models.Voucher{
					ID:      id_voucher1,
					Number:  number2,
					Version: 1,
				},
				Items: vrequest.VoucherItemUpdateList{
					Inserted: []vrequest.VoucherItemInsertion{
					},
					Updated: []vrequest.VoucherItemUpdate{
					},
					Deleted: []uint{}, 
				},
			}
            str,err:=services.UpdateVoucher(&updateRequest)
			fmt.Println(err.Error())
			assert.NotNil(t,err)
			assert.Empty(t,str)

		 })
		t.Run("should not allow Update with  empty number", func(t *testing.T) {
			updateRequest := vrequest.VoucherUpdateRequest{
				Voucher: models.Voucher{
					ID:      id_voucher1,
					Number:  "",
					Version: 1,
				},
				Items: vrequest.VoucherItemUpdateList{
					Inserted: []vrequest.VoucherItemInsertion{
					},
					Updated: []vrequest.VoucherItemUpdate{
					},
					Deleted: []uint{}, 
				},
			}
            str,err:=services.UpdateVoucher(&updateRequest)
			fmt.Println(err.Error())
			assert.NotNil(t,err)
			assert.Empty(t,str)

		})
		t.Run("should not allow Update with  number that has more than 64 characters", func(t *testing.T) {
			number3:=helper.GenerateStringWithBigSize()
			updateRequest := vrequest.VoucherUpdateRequest{
				Voucher: models.Voucher{
					ID:      id_voucher1,
					Number:  number3,
					Version: 1,
				},
				Items: vrequest.VoucherItemUpdateList{
					Inserted: []vrequest.VoucherItemInsertion{
					},
					Updated: []vrequest.VoucherItemUpdate{
					},
					Deleted: []uint{}, 
				},
			}
            str,err:=services.UpdateVoucher(&updateRequest)
			fmt.Println(err.Error())
			assert.NotNil(t,err)
			assert.Empty(t,str)

		})
		t.Run("should not allow Update with inserted items that both of creits and debits are more than 0", func(t *testing.T) {
			number5:=helper.GenerateNumericString()
			updateRequest := vrequest.VoucherUpdateRequest{
				Voucher: models.Voucher{
					ID:      id_voucher1,
					Number:  number5,
					Version: 1,
				},
				Items: vrequest.VoucherItemUpdateList{
					Inserted: []vrequest.VoucherItemInsertion{{
						VoucherID: id_voucher1,
						SLID:      id_sl_withdl,
						DLID:      &id_dl,
						Debit:     200,
						Credit:    600,
					},
				
					},
					Updated: []vrequest.VoucherItemUpdate{
					},
					Deleted: []uint{}, 
				},
			}
            str,err:=services.UpdateVoucher(&updateRequest)
			fmt.Println("to  "+err.Error())
			assert.NotNil(t,err)
			assert.Empty(t,str)

		})
		t.Run("should not allow Update with inserted items that sum of creits and debits is not 0", func(t *testing.T) {
			number5:=helper.GenerateNumericString()
			updateRequest := vrequest.VoucherUpdateRequest{
				Voucher: models.Voucher{
					ID:      id_voucher1,
					Number:  number5,
					Version: 1,
				},
				Items: vrequest.VoucherItemUpdateList{
					Inserted: []vrequest.VoucherItemInsertion{{
						VoucherID: id_voucher1,
						SLID:      id_sl_withdl,
						DLID:      &id_dl,
						Debit:     0,
						Credit:    600,
					},
					{
						VoucherID: id_voucher1,
						SLID:     id_sl_withdl,
						DLID:     &id_dl,
						Debit:     600,
						Credit:    0,
					},
					{
						VoucherID: id_voucher1,
						SLID:     id_sl_withdl,
						DLID:     &id_dl,
						Debit:     200,
						Credit:    0,
					},
					},
					Updated: []vrequest.VoucherItemUpdate{
					},
					Deleted: []uint{}, 
				},
			}
            str,err:=services.UpdateVoucher(&updateRequest)
			fmt.Println(err.Error())
			assert.NotNil(t,err)
			assert.Empty(t,str)

		})
		t.Run("should not allow Update with  inserted  items that sl field has dl but  in request we dont have dl_id", func(t *testing.T) {
			number5:=helper.GenerateNumericString()
			updateRequest := vrequest.VoucherUpdateRequest{
				Voucher: models.Voucher{
					ID:      id_voucher1,
					Number:  number5,
					Version: 1,
				},
				Items: vrequest.VoucherItemUpdateList{
					Inserted: []vrequest.VoucherItemInsertion{{
						VoucherID: id_voucher1,
						SLID:      id_sl_withdl,
						DLID:      nil,
						Debit:     0,
						Credit:    600,
					},
					{
						VoucherID: id_voucher1,
						SLID:     id_sl_withdl,
						DLID:     &id_dl,
						Debit:     600,
						Credit:    0,
					},
					
					},
					Updated: []vrequest.VoucherItemUpdate{
					},
					Deleted: []uint{}, 
				},
			}
            str,err:=services.UpdateVoucher(&updateRequest)
			fmt.Println(err.Error())
			assert.NotNil(t,err)
			assert.Empty(t,str)

		})

		t.Run("should not allow Update with inserted  items that sl_id field of them is empty", func(t *testing.T) {
			number5:=helper.GenerateNumericString()
			updateRequest := vrequest.VoucherUpdateRequest{
				Voucher: models.Voucher{
					ID:      id_voucher1,
					Number:  number5,
					Version: 1,
				},
				Items: vrequest.VoucherItemUpdateList{
					Inserted: []vrequest.VoucherItemInsertion{{
						VoucherID: id_voucher1,
						DLID:      &id_dl,
						Debit:     0,
						Credit:    600,
					},
					{
						VoucherID: id_voucher1,
						SLID:     id_sl_withdl,
						DLID:     &id_dl,
						Debit:     600,
						Credit:    0,
					},
					},
					Updated: []vrequest.VoucherItemUpdate{
					},
					Deleted: []uint{}, 
				},
			}
            str,err:=services.UpdateVoucher(&updateRequest)
			fmt.Println(err.Error())
			assert.NotNil(t,err)
			assert.Empty(t,str)

		})
		t.Run("should not allow Update with update items that sl_id field of them is empty", func(t *testing.T) {
			id_v2,err:=get.GetIDByNumber("voucher",number2)
			assert.Nil(t,err)
			id_last_item,err:=get.GetLastVoucherItemIndex(id_v2)
			assert.Nil(t,err)
			number5:=helper.GenerateNumericString()
			updateRequest := vrequest.VoucherUpdateRequest{
				Voucher: models.Voucher{
					ID:      id_v2,
					Number:  number5,
					Version: 1,
				},
				Items: vrequest.VoucherItemUpdateList{
					Inserted: []vrequest.VoucherItemInsertion{
					},
					Updated: []vrequest.VoucherItemUpdate{{
                     VoucherID: id_v2,
					 ID: id_last_item,
					 SLID: 0,
					 DLID: nil,
					 Credit: 200,
					 Debit: 0,
					},
					},
					Deleted: []uint{}, 
				},
			}
            str,err:=services.UpdateVoucher(&updateRequest)
			fmt.Println("T2 "+err.Error())
			assert.NotNil(t,err)
			assert.Empty(t,str)

		})
		t.Run("should not allow Update with update items that both of credits and debits field has value more than 0", func(t *testing.T) {
			id_v2,err:=get.GetIDByNumber("voucher",number2)
			assert.Nil(t,err)
			id_last_item,err:=get.GetLastVoucherItemIndex(id_v2)
			assert.Nil(t,err)
			number5:=helper.GenerateNumericString()
			updateRequest := vrequest.VoucherUpdateRequest{
				Voucher: models.Voucher{
					ID:      id_v2,
					Number:  number5,
					Version: 1,
				},
				Items: vrequest.VoucherItemUpdateList{
					Inserted: []vrequest.VoucherItemInsertion{
					},
					Updated: []vrequest.VoucherItemUpdate{{
                     VoucherID: id_v2,
					 ID: id_last_item,
					 SLID: id_sl_withdl1,
					 DLID: &id_dl1,
					 Credit: 200,
					 Debit: 10,
					},
					},
					Deleted: []uint{}, 
				},
			}
            str,err:=services.UpdateVoucher(&updateRequest)
			fmt.Println("T5 "+err.Error())
			assert.NotNil(t,err)
			assert.Empty(t,str)

		})
		t.Run("should not allow Update with update items that makes the voucher unbalanced", func(t *testing.T) {
			id_v2,err:=get.GetIDByNumber("voucher",number2)
			assert.Nil(t,err)
			id_last_item,err:=get.GetLastVoucherItemIndex(id_v2)
			assert.Nil(t,err)
			number5:=helper.GenerateNumericString()
			updateRequest := vrequest.VoucherUpdateRequest{
				Voucher: models.Voucher{
					ID:      id_v2,
					Number:  number5,
					Version: 1,
				},
				Items: vrequest.VoucherItemUpdateList{
					Inserted: []vrequest.VoucherItemInsertion{
					},
					Updated: []vrequest.VoucherItemUpdate{{
                     VoucherID: id_v2,
					 ID: id_last_item,
					 SLID: id_sl_withdl1,
					 DLID: &id_dl,
					 Credit: 500,
					 Debit: 0,
					},
					},
					Deleted: []uint{}, 
				},
			}
            str,err:=services.UpdateVoucher(&updateRequest)
			fmt.Println("T3 "+err.Error())
			assert.NotNil(t,err)
			assert.Empty(t,str)

		})
		t.Run("should not allow Update with delete items that count of items of voucher become less than 2", func(t *testing.T) {
			id_v2,err:=get.GetIDByNumber("voucher",number2)
			assert.Nil(t,err)
			id_last_item,err:=get.GetLastVoucherItemIndex(id_v2)
			assert.Nil(t,err)
			number5:=helper.GenerateNumericString()
			updateRequest := vrequest.VoucherUpdateRequest{
				Voucher: models.Voucher{
					ID:      id_v2,
					Number:  number5,
					Version: 1,
				},
				Items: vrequest.VoucherItemUpdateList{
					Inserted: []vrequest.VoucherItemInsertion{
					},
					Updated: []vrequest.VoucherItemUpdate{
					},
					Deleted: []uint{id_last_item}, 
				},
			}
            str,err:=services.UpdateVoucher(&updateRequest)
			fmt.Println("T1"+err.Error())
			assert.NotNil(t,err)
			assert.Empty(t,str)

		})
        
		t.Run("should not allow update when mismatch version", func(t *testing.T) {
			number4:=helper.GenerateNumericString()
			updateRequest4 := vrequest.VoucherUpdateRequest{
				Voucher: models.Voucher{
					ID:      id_voucher1,
					Number: number4,
					Version: 1,
				},
				Items: vrequest.VoucherItemUpdateList{
					Inserted: []vrequest.VoucherItemInsertion{
					},
					Updated: []vrequest.VoucherItemUpdate{
					},
					Deleted: []uint{}, 
				},
			}
            str,err:=services.UpdateVoucher(&updateRequest4)
			fmt.Println(str)
			assert.Nil(t,err)
			assert.NotEmpty(t,str)
            number5:=helper.GenerateNumericString()+"9"
			updateRequest5 := vrequest.VoucherUpdateRequest{
				Voucher: models.Voucher{
					ID:      id_voucher1,
					Number: number5,
					Version: 1,
				},
				Items: vrequest.VoucherItemUpdateList{
					Inserted: []vrequest.VoucherItemInsertion{
					},
					Updated: []vrequest.VoucherItemUpdate{
					},
					Deleted: []uint{}, 
				},
			}
			str,err=services.UpdateVoucher(&updateRequest5)
			fmt.Println(err.Error())
			assert.NotNil(t,err)
			assert.Empty(t,str)

		})

	})
	

}
func TestDeleteGeneralDetailed(t *testing.T) {
	var id_dl uint
	var id_sl_withdl uint
	var id_voucher1 uint
	t.Run("deleting details", func(t *testing.T) {
		    number1:=helper.GenerateNumericString()

			code1 := helper.GenerateNumericString()
			title1 := helper.GenerateRandomString()

			code2 := helper.GenerateNumericString()
			title2 := helper.GenerateRandomString()
			_, err1 := services.Insertdl(dlrequest.DlInsertRequest{Code: code1, Title: title1})
			assert.Nil(t, err1)
			_, err2 := services.Insertsl(slrequest.SlInsertRequest{Code: code2, Title: title2, IsDetailable: true})
			assert.Nil(t, err2)

			id_dl1, err := get.GetIDByCode("dl", code1)
			assert.Nil(t, err)
			id_dl=id_dl1
			id_sl_withdl1, err := get.GetIDByCode("sl", code2)
			assert.Nil(t, err)
			id_sl_withdl=id_sl_withdl1
		
			request1 := vrequest.VoucherInsertRequest{
				Voucher: vrequest.VoucherInsertion{Number: number1},
				Items: []vrequest.VoucherItemInsertion{{SLID: id_sl_withdl, DLID: &id_dl, Debit: 300, Credit: 0},
					{SLID: id_sl_withdl, DLID: &id_dl, Debit: 0, Credit: 300}},
			}
			
			str1, err1 := services.Insertvoucher(request1)
			assert.Nil(t, err1, "expected first creation to succeed")
			assert.Equal(t, str1, "successful voucher insertion request")
            id_v1,err:=get.GetIDByNumber("voucher",number1)
			assert.Nil(t, err)
			id_voucher1=id_v1
		t.Run("should not allow Delete Voucher with mismatch version", func(t *testing.T) {
			number2:=helper.GenerateNumericString()
			updateRequest := vrequest.VoucherUpdateRequest{
				Voucher: models.Voucher{
					ID:      id_voucher1,
					Number:  number2,
					Version: 1,
				},
				Items: vrequest.VoucherItemUpdateList{
					Inserted: []vrequest.VoucherItemInsertion{
					},
					Updated: []vrequest.VoucherItemUpdate{
					},
					Deleted: []uint{}, 
				},
			}
            str,err:=services.UpdateVoucher(&updateRequest)
			fmt.Println(str)
			assert.Nil(t,err)
			assert.NotEmpty(t,str)
            req:=vrequest.VoucherDeleteRequest{ID:  id_voucher1,Version :1}
			str, err = services.DeleteVoucherWithVersion(req)
			fmt.Println(err.Error())
			assert.NotNil(t, err)
			assert.Empty(t, str)

		 })

		 t.Run("should not allow Delete sl that is refrenced", func(t *testing.T) {
			req:=slrequest.SlDeleteRequest{ID: id_sl_withdl,Version :1}
			str, err := services.DeleteSLWithVersion(req)
			fmt.Println(err.Error())
			assert.NotNil(t, err)
			assert.Empty(t, str)
		 })
		 t.Run("should not allow update sl that is refrenced", func(t *testing.T) {
			req:=slrequest.SlUpdateRequest{ID: id_sl_withdl, Code: code2, Title: title2, Version: 1,IsDetailable: true}
            str,err:=services.Updatesl(req)
			fmt.Println(err.Error())
			assert.NotNil(t,err)
			assert.Empty(t,str)
		 })
		 t.Run("should not allow Delete dl that is refrenced", func(t *testing.T) {
			req:=dlrequest.DlDeleteRequest{ID: id_dl,Version :1}
			str, err := services.DeleteDLWithVersion(req)
			fmt.Println(err.Error())
			assert.NotNil(t, err)
			assert.Empty(t, str)
		 })

	})


	

}