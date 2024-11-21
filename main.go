package main

import (
	// "finalproject/db"
	"finalproject/db"
	"finalproject/internal/models"
	// "log"
	//"finalproject/internal/request/slrequest"

	//"finalproject/internal/models"
	//"finalproject/internal/request/dlrequest"
	//"finalproject/db/queries/update"
	//"finalproject/internal/models"
	//"finalproject/internal/request/slrequest"

	"finalproject/internal/request/vrequest"
	"finalproject/internal/services"
	"fmt"
)

func main() {
	// insertreqdl := `{
	// 	"code":"1252",
	// 	 "title":""
	//   }`
	// var request dlrequest.DlInsertRequest
	// err := json.Unmarshal([]byte(insertreqdl), &request)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// if str, err := services.Insertdl(request); err != nil {
	// 	fmt.Printf(err.Error())
	// } else {
	// 	fmt.Println(str)
	// }

	//done

	// if str, err := services.Updatedl(dlrequest.DlUpdateRequest{ID: 23, Code: "122", Title: "قرارداد بانکی", Version: 1}); err != nil {
	// 	fmt.Println("lll")
	// 	fmt.Println(err.Error())
	// 	fmt.Println(str)
	// } else {
	// 	print(str)
	// }
	// if str, err := services.GetDLByID(2); err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Printf("hi%v\n", *str)
	// }
	// if err:=services.DeleteDLWithVersion(dlrequest.DlDeleteRequest{ID:10,Version: 1});err!=nil{
	// 	fmt.Println("Error:", err)
	// }
	// if str,err:=services.Insertsl(slrequest.SlInsertRequest{Code: "15636",Title: "وام اشتغال زایی",IsDetailable: true});err!=nil{
	// 	fmt.Printf(err.Error())
	//  }else{
	// 	fmt.Println(str)
	//  }
	// if str,err:=services.Updatesl(slrequest.SlUpdateRequest{ID: 1,Code: "109",Title: "کمک هزینه خرید مسکن" ,Version:2,IsDetailable: true });err!=nil{
	// 	fmt.Println("Error:", err)
	//  }else{
	// 	fmt.Println(str)
	//  }
	// if err:=services.DeleteSLWithVersion(slrequest.SlDeleteRequest{ID:7,Version: 1});err!=nil{
	//    fmt.Printf(err.Error())
	// }
	// if str,err:= services.GetSLByID(4);err!=nil{
	// 		println("cant get")
	// 	 }else{
	// 	fmt.Printf("hi%v\n",*str)
	// 	    }
	// var a uint
	// a=1
	// request := vrequest.VoucherInsertRequest{
	// 	Voucher: vrequest.VoucherInsertion{
	// 		Number:  "55645",
	// 		Version: "1",
	// 	},
	// 	Items: []vrequest.VoucherItemInsertion{
	// 		{

	// 			SLID:      8,
	// 			DLID:      &a, // No DLID for this item
	// 			Debit:     200,
	// 			Credit:    0,
	// 		},
	// 		{

	// 			SLID:      1,
	// 			DLID:      nil, // A DLID pointer
	// 			Debit:     10,
	// 			Credit:    200,
	// 		},
	// 	},
	// }
	// if str,err:=services.Insertvoucher(request);err!=nil{
	// 	println(err.Error())
	// }else{
	// 	println(str)
	// }

	// if str,err:= services.GetVoucherByID(11);err!=nil{
	// 		println("cant get")
	// 	 }else{
	// 	fmt.Printf("hi%v\n",*str)
	// 	    }

	// if err:=services.DeleteVoucher(vrequest.VoucherDeleteRequest{ID:13,Version: 1});err!=nil{
	//    fmt.Printf(err.Error())
	// }
db.InitDB()		
var x uint
x=11
updateRequest := vrequest.VoucherUpdateRequest{
    Voucher: models.Voucher{
        ID:      13,
        Number:  "12300",
        Version: 18,
    },
    Items: vrequest.VoucherItemUpdateList{
        Inserted: []vrequest.VoucherItemInsertion{
            {
                VoucherID: 13,
                SLID:      84,
                DLID:      &x,
                Debit:     0,
                Credit:    600,
            },
            {
                VoucherID: 13,
                SLID:      82,
                DLID:      nil,
                Debit:     600,
                Credit:    0,
            },
        },
        Updated: []vrequest.VoucherItemUpdate{
           
        },
        Deleted: []uint{}, // IDs of items to delete
    },
}
	// var request vrequest.VoucherUpdateRequest

	// // Unmarshal JSON into the struct
	// err := json.Unmarshal([]byte(jsond), &request)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	
	if _,err := services.UpdateVoucher(&updateRequest); err != nil {
		fmt.Println(err.Error())
	}

}
