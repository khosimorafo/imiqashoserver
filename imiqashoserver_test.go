package imiqashoserver_test

import (
	"testing"
	"os"
	"github.com/khosimorafo/imiqashoserver"
	"time"
	"fmt"
	"encoding/json"
	"github.com/antonholmquist/jason"
	"github.com/jinzhu/now"
)

var a imiqashoserver.App

func TestMain(m *testing.M) {

	a = imiqashoserver.App{}

	a.Initialize()

	//ensureTableExists()

	code := m.Run()

	//clearTable()

	os.Exit(code)
}

/*
func TestCreateFinancialPeriodRange(t *testing.T) {

	err := imiqashoserver.RemoveFinancialPeriodRange()
	if err != nil{

		t.Errorf("Failed to remove records. > %v", err)
	}

	err1 := imiqashoserver.CreateFinancialPeriodRange("2017-01-14", 12)

	if err1 !=nil {

		t.Errorf("Failed to create period. > %v", err1)
		return
	}
}
*/
func TestReadFinancialPeriodRange(t *testing.T) {

	ps, err := imiqashoserver.ReadFinancialPeriodRange("open")

	if err != nil{

		t.Errorf("Failed to read records.  > %v", err)
	}

	t.Log(len(ps))
}

func TestGetDaysLeft(t *testing.T) {

	start_date, err1 := now.Parse("2017-04-14")

	if err1 != nil{

		t.Error("Failed to parse date : %v",err1)
		return
	}
	period := imiqashoserver.P{start_date}

	days, all, err := period.GetDaysLeft()

	if err != nil{

		t.Error("Failed to get number of days : %v",err)
		return
	}

	t.Log("No of days left is ", days)
	t.Log("Total days in month is ", all)
}

func TestGetProRataDays(t *testing.T) {

	start_date, err1 := now.Parse("2017-04-14")

	if err1 != nil{

		t.Error("Failed to parse date : %v",err1)
		return
	}
	period := imiqashoserver.P{start_date}

	pr, err := period.GetProRataDays()

	if err1 != nil{

		t.Error("Failed to evaluate prorata days : %v",err)
		return
	}

	t.Log("Pro rata is : ", pr)


}

func TestGetPeriod(t *testing.T) {

	period := imiqashoserver.P{time.Now()}

	p, err := period.GetPeriod()

	if err != nil{

		t.Error("Failed to get a period for the date given : ")
		return
	}

	t.Log("Index is ", p.Index)
	t.Log("Name is ", p.Name)
}

func TestGetPeriodByPeriodName(t *testing.T) {

	p, err := imiqashoserver.GetPeriodByName("May-2017")

	if err != nil{

		t.Error("Failed to get a period for the date given : ")
		return
	}

	t.Log("Index is ", p.Index)
	t.Log("Name is ", p.Name)
}

func TestGetPeriodByPeriodIndex(t *testing.T) {

	p, err := imiqashoserver.GetPeriodByIndex(5)

	if err != nil{

		t.Error("Failed to get a period for the date given : ")
		return
	}

	t.Log("Index is ", p.Index)
	t.Log("Name is ", p.Name)
}

func TestGetNextPeriodByPeriodName(t *testing.T) {

	p, err := imiqashoserver.GetNextPeriodByName("May-2017")

	if err != nil{

		t.Error("Failed to get a period for the date given : ")
		return
	}

	t.Log("Index is ", p.Index)
	t.Log("Name is ", p.Name)
}

func TestCreateandDeleteLatePaymentRequest(t *testing.T) {

	var pay imiqashoserver.LatePayment
	layout := "2006-01-02"

	t1 := time.Now()
	fmt.Println(t1.Format(layout))
	pay.CustomerID = "2343253532"
	pay.InvoiceID = "2343253532"
	pay.Date = t1.Format(layout)
	pay.Period = "May-2017"
	pay.Status = "approved"

	_, err := pay.Create()
	if err != nil{

		t.Error("Failed to make payment extention request!")
		return
	}


	//Delete
	_, err_del := pay.Delete()

	if err_del != nil{

		t.Error(err_del)
		return
	}
}

func TestChangePaymentRequestStatus(t *testing.T) {

	var pay imiqashoserver.LatePayment
	layout := "2006-01-02"

	t1 := time.Now()
	fmt.Println(t1.Format(layout))
	pay.CustomerID = "2343253532"
	pay.InvoiceID = "2343253532"
	pay.Date = t1.Format(layout)
	pay.Period = "May-2017"
	pay.Status = "approved"

	_, err := pay.Create()
	if err != nil{

		t.Error("Failed to make payment extention request!")
		return
	}

	//Read and set status to rejected
	pay.RequestStatusAsRejected()

	_, payment, err_rej := pay.Read()

	if err_rej != nil {

		t.Errorf("Failed to read!")
		return
	}

	b, _ := json.Marshal(payment)
	v, _ := jason.NewObjectFromBytes(b)
	status, _ := v.GetString("status")

	if status != "rejected" {
		t.Errorf("Expected status REJECTED. Got %v", status)
	}

	t.Log(v.String())

	//Read and set status to expired
	pay.RequestStatusAsExpired()

	_, payment, err_exp := pay.Read()

	if err_exp != nil {

		t.Errorf("Failed to read!")
		return
	}

	b_exp, _ := json.Marshal(payment)
	v_exp, _ := jason.NewObjectFromBytes(b_exp)
	status_exp, _ := v_exp.GetString("status")

	if status_exp != "expired" {
		t.Errorf("Expected status EXPIRED. Got %v", status_exp)
	}

	t.Log(v_exp.String())

	//Read and set status to paid
	pay.RequestStatusAsPaid()

	_, payment, err_paid := pay.Read()

	if err_paid != nil {

		t.Errorf("Failed to read!")
		return
	}

	b_paid, _ := json.Marshal(payment)
	v_paid, _ := jason.NewObjectFromBytes(b_paid)
	status_paid, _ := v_paid.GetString("status")

	if status_paid != "paid" {
		t.Errorf("Expected status PAID. Got %v", status_paid)
	}

	t.Log(v_paid.String())

	//Delete
	_, err_del := pay.Delete()

	if err_del != nil{

		t.Error(err_del)
		return
	}
}

/*
func TestRemoveFinancialPeriodRange(t *testing.T) {

	err := imiqashoserver.RemoveFinancialPeriodRange()
	if err != nil{

		t.Errorf("Failed to remove records. > %v", err)
	}
}*/