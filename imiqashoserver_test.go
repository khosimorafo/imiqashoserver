package imiqashoserver_test

import (
	"testing"
	"os"
	"github.com/khosimorafo/imiqashoserver"
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

	err := a.CreateFinancialPeriodRange("2017-01-14", 12)

	if err !=nil {

		t.Errorf("Failed to create period. > %v", err)
		return
	}
}
*/

func TestReadFinancialPeriodRange(t *testing.T) {

	ps, err := a.ReadFinancialPeriodRange("open")

	if err != nil{

		t.Errorf("Failed to read records.  > %v", err)
	}

	t.Log(len(ps))
}

/*
func TestRemoveFinancialPeriodRange(t *testing.T) {

	err := a.RemoveFinancialPeriodRange()
	if err != nil{

		t.Errorf("Failed to remove records. > %v", err)
	}
}
*/