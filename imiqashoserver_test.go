package imiqashoserver_test

import (
	"testing"
	"os"
	"github.com/khosimorafo/imiqashoserver"

	"github.com/jinzhu/now"
	"time"
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

func TestReadFinancialPeriodRange(t *testing.T) {

	ps, err := imiqashoserver.ReadFinancialPeriodRange("open")

	if err != nil{

		t.Errorf("Failed to read records.  > %v", err)
	}

	t.Log(len(ps))
}
*/

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

/*
func TestRemoveFinancialPeriodRange(t *testing.T) {

	err := imiqashoserver.RemoveFinancialPeriodRange()
	if err != nil{

		t.Errorf("Failed to remove records. > %v", err)
	}
}
*/