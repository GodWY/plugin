package template

const EnumTemplate = ` 
syntax = "{{.Version}}";

package emum;

option go_package = "{{.Path}};

// {{.Template}}
enum {{.Template}}{
	{{range $key, $value := .Data}}
		{{$key}} = {{$value}}
	{{end}}
}
`
