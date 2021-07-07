package main

import "plugin/exceltopb"

const EnumTemplate = ` 
syntax = "{{.Version}}";

package emum;

option go_package = "{{.Path}};

// {{.Template}}
enum {{.Template}}{
	{{range $index, $element := .Arry}}
		{{$element}} = {{$index}}
	{{end}}
}
`

type Info struct {
	Version  string
	Path     string
	Arry     []string
	Template string
}

func main() {
	// i := &Info{
	// 	Version:  "proto3",
	// 	Path:     "/gen/path",
	// 	Arry:     []string{"a", "b"},
	// 	Template: "test",
	// }
	// t, _ := template.New("oo").Parse(EnumTemplate)
	// t.Execute(os.Stdout, i)
	exceltopb.ReadExcelSheet("../excel/test.xlsx")
}
