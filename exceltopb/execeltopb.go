package exceltopb

import (
	"fmt"
	"os"
	self "plugin/exceltopb/template"
	"strings"
	"text/template"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// 读取excel数据
func ReadExcelSheet(path string) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		panic(err.Error())
	}
	sheets := f.WorkBook.Sheets
	for _, sheet := range sheets.Sheet {
		goType := strings.Split(sheet.Name, "@")
		if len(goType) <= 1 {
			panic("sheet is error and lock @name or name@")
		}
		switch goType[1] {
		case Enum:
			ReadColumnBySheet(f, sheet.Name, sheet.SheetID, sheet.ID, sheet.State)
		case Const:
		default:
			fmt.Println("类型", goType[1])
		}
	}
}

// ReadColumnBySheet
func ReadColumnBySheet(f *excelize.File, name string, sheetId int, id, state string) {
	rows, err := f.GetRows(name)
	if err != nil {
		panic(err.Error())
	}
	// rows[3:]
	data := map[string]string{}
	for _, row := range rows[2:] {
		data[row[0]] = row[1]
	}
	ExcelToEnumPb(name, sheetId, id, state, data)
}

type EnumPb struct {
	Version  string
	Path     string
	Data     map[string]string
	Template string
}

// 读取excel数据转化为protocol pb
func ExcelToEnumPb(name string, sheetId int, id, state string, data map[string]string) {
	ep := new(EnumPb)
	ep.Path = "proto/common"
	ep.Data = data
	ep.Template = strings.Split(name, "@")[0]
	ep.Version = "proto3"
	t, _ := template.New(name).Parse(self.EnumTemplate)
	t.Execute(os.Stdout, ep)
}

func WriteToPb() {
}
