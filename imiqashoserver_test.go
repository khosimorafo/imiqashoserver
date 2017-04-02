package imiqashoserver_test

import (
	"testing"
	"os"
	"github.com/khosimorafo/imiqashoserver"
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

/*
func TestRemoveFinancialPeriodRange(t *testing.T) {

	err := imiqashoserver.RemoveFinancialPeriodRange()
	if err != nil{

		t.Errorf("Failed to remove records. > %v", err)
	}
}
*/