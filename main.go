package main

import (
	// "finalproject/db"
	// "finalproject/internal/models"
	// "log"
	//"finalproject/internal/request/slrequest"
	"encoding/json"
	//"finalproject/internal/models"
	"finalproject/internal/request/dlrequest"

	//"finalproject/internal/request/vrequest"
	"finalproject/internal/services"
	"fmt"
)

func main() {

	// db.ConnectWithGORM()

	// // Insert into DL table

	// newDL := models.DL{
	// 	Code:  "1001",
	// 	Title: "Accounts Receivable",
	// }
	insertreqdl:=`{
	"code":"0796",
	 "title":"منابع باستانی"
  }`
  var request dlrequest.DlInsertRequest
  err := json.Unmarshal([]byte(insertreqdl), &request)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
if str,err:=services.Insertdl(request);err!=nil{
		fmt.Printf(err.Error())
	 }else{
		fmt.Println(str)
	 }

	// log.Printf("DL record created: %+v", newDL)
	//    if str,err:= services.Updatedl(request.DlUpdateRequest{ID:2,Code: "1232",Title: "قرارداد بانکی"});err!=nil{
	// 	println("cant update")
	//    }else{
	// 	print(str)
	//    }
	// if str,err:= services.GetDLByID(1);err!=nil{
	// 		println("cant get")
	// 	 }else{
	// 	fmt.Printf("hi%v\n",*str)
	// 	    }
	if err:=services.DeleteDLWithVersion(dlrequest.DlDeleteRequest{ID:1,Version: 1});err!=nil{
	   fmt.Printf(err.Error())
	}
	// if str,err:=services.Insertdl(request.DlInsertRequest{Code: "69",Title: ""});err!=nil{
	// 	fmt.Printf(err.Error())
	//  }else{
	// 	fmt.Println(str)
	//  }
	// if str,err:=services.Updatedl(request.DlUpdateRequest{ID: 5,Code: "236",Title: " وام مسکن مهر"});err!=nil{
	// 	fmt.Printf(err.Error())
	//  }else{
	// 	fmt.Println(str)
	//  }
	//  if str,err:=services.Updatedl(request.DlUpdateRequest{ID: 5,Code: "236",Title: " وام مسکن مهر"});err!=nil{
	// 	fmt.Printf(err.Error())
	//  }else{
	// 	fmt.Println(str)
	//  }

	// if str,err:=services.Insertsl(request.SlInsertRequest{Code: "136",Title: "ساخت مدرسه"});err!=nil{
	// 	fmt.Printf(err.Error())
	//  }else{
	// 	fmt.Println(str)
	//  }
	// if str,err:=services.Updatesl(request.SlUpdateRequest{ID: 2,Code: "102",Title: "وام کشاورزی"});err!=nil{
	// 	fmt.Printf(err.Error())
	//  }else{
	// 	fmt.Println(str)
	//  }

	// if err:=services.DeleteSLWithVersion(slrequest.SlDeleteRequest{ID:2,Version: 3});err!=nil{
	//    fmt.Printf(err.Error())
	// }

	// if str,err:= services.GetSLByID(3);err!=nil{
	// 		println("cant get")
	// 	 }else{
	// 	fmt.Printf("hi%v\n",*str)
	// 	    }
	// request := request.VoucherInsertionRequest{
	// 	Voucher: request.VoucherInsert{
	// 		Number:  "123325",
	// 		Version: "1",
	// 	},
	// 	Items: []request.VoucherItemrequest{
	// 		{

	// 			SLID:      2,
	// 			DLID:      nil, // No DLID for this item
	// 			Debit:     5200,
	// 			Credit:    0,
	// 		},
	// 		{

	// 			SLID:      1,
	// 			DLID:      nil, // A DLID pointer
	// 			Debit:     0,
	// 			Credit:    200,
	// 		},
	// 	},
	// }
	// if str,err:=services.Insertvoucher(request);err!=nil{
	// 	println(err.Error())
	// }else{
	// 	println(str)
	// }

	// if str,err:= services.GetvoucherByID(1);err!=nil{
	// 		println("cant get")
	// 	 }else{
	// 	fmt.Printf("hi%v\n",*str)
	// 	    }
	// }

	// if err:=services.DeleteVoucherWithVersion(request.VoucherDeleteRequest{ID:16,Version: 1});err!=nil{
	//    fmt.Printf(err.Error())
	// }
// 	jsond := `{
// 	"voucher": {
// 	  "id": 13,
// 	  "number": "900",
// 	  "version": 1
// 	},
// 	"items": {
// 	  "inserted":[
// 	  ],
// 	  "updated": [
// 		{
// 	      "id":5,
// 		  "sl_id": 2,
// 		  "debit": 0,
// 		  "credit": 600
// 		}
// 	  ],
// 	  "deleted": [6]
// 	}
//   }`
	// var request vrequest.VoucherUpdateRequest

	// // Unmarshal JSON into the struct
	// err := json.Unmarshal([]byte(jsond), &request)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Printf("correct %v\n", request)
	// if err := services.UpdateVoucher(&request); err != nil {
	// 	fmt.Println(err.Error())
	// }

}
