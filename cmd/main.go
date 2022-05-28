// This program generates contributors.go. It can be invoked by running
// go generate
package main

//go:generate go run main.go

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type Input struct {
	NullType string
	StdType  string
	AltType  string
}

type Data struct {
	Inputs []Input
}

func main() {
	data := Data{Inputs: []Input{
		{
			NullType: "NullInt32",
			StdType:  "Int32",
			AltType:  "Default",
		},
		{
			NullType: "NullInt32",
			StdType:  "Int32",
			AltType:  "Panic",
		},
		{
			NullType: "NullInt32",
			StdType:  "Int32",
			AltType:  "Error",
		},
		{
			NullType: "NullInt16",
			StdType:  "Int16",
			AltType:  "Default",
		},
		{
			NullType: "NullInt16",
			StdType:  "Int16",
			AltType:  "Panic",
		},
		{
			NullType: "NullInt16",
			StdType:  "Int16",
			AltType:  "Error",
		},
		{
			NullType: "NullInt64",
			StdType:  "Int64",
			AltType:  "Default",
		},
		{
			NullType: "NullInt64",
			StdType:  "Int64",
			AltType:  "Panic",
		},
		{
			NullType: "NullInt64",
			StdType:  "Int64",
			AltType:  "Error",
		},

		{
			NullType: "NullFloat64",
			StdType:  "Float64",
			AltType:  "Default",
		},
		{
			NullType: "NullFloat64",
			StdType:  "Float64",
			AltType:  "Panic",
		},
		{
			NullType: "NullFloat64",
			StdType:  "Float64",
			AltType:  "Error",
		},

		{
			NullType: "NullString",
			StdType:  "String",
			AltType:  "Default",
		},
		{
			NullType: "NullString",
			StdType:  "String",
			AltType:  "Panic",
		},
		{
			NullType: "NullString",
			StdType:  "String",
			AltType:  "Error",
		},

		{
			NullType: "NullBool",
			StdType:  "Bool",
			AltType:  "Default",
		},
		{
			NullType: "NullBool",
			StdType:  "Bool",
			AltType:  "Panic",
		},
		{
			NullType: "NullBool",
			StdType:  "Bool",
			AltType:  "Error",
		},
		{
			NullType: "NullByte",
			StdType:  "Byte",
			AltType:  "Default",
		},
		{
			NullType: "NullByte",
			StdType:  "Byte",
			AltType:  "Panic",
		},
		{
			NullType: "NullByte",
			StdType:  "Byte",
			AltType:  "Error",
		},
	},
	}
	funcMap := template.FuncMap{
		"ToLower": strings.ToLower,
	}

	t := template.Must(template.New("type2type").Funcs(funcMap).Parse(queueTemplate))
	f, err := os.Create("../sqlNulls.go")
	if err != nil {
		log.Println("create file: ", err)
		return
	}
	defer f.Close()
	if err := t.Execute(f, data); err != nil {
		panic(err)
	}
}

var queueTemplate = `package type2type
import (
	"database/sql"
	"errors"
	"fmt"
)
var (
{{ range $item := .Inputs }}
{{ if eq $item.AltType "Error" }}
ErrInvalid{{ .NullType }}To{{ $item.StdType }}Or{{ $item.AltType }} = errors.New("failed to convert from {{$item.NullType}} to {{ $item.StdType }} or {{ $item.AltType}}")
{{ end }} {{ end }}
)
{{ range $item :=  .Inputs}}
{{ if eq $item.AltType "Error" }}
	func {{ $item.NullType }}To{{ $item.StdType }}Or{{ $item.AltType }}(nt sql.{{ $item.NullType }}) ({{ $item.StdType | ToLower }},error) {

{{ else }}
	func {{ $item.NullType }}To{{ $item.StdType }}Or{{ $item.AltType }}(nt sql.{{ $item.NullType }}) {{ $item.StdType | ToLower }} {


{{ end }}
	var v {{ $item.StdType | ToLower }}
	if nt.Valid {
		v = nt.{{ $item.StdType }}
	}
	{{ if eq $item.AltType "Panic" }}
	if !nt.Valid{
	panic(ErrInvalid{{ .NullType }}To{{ $item.StdType }}OrError)
}
{{ end }}
{{ if eq $item.AltType "Error" }}
	if !nt.Valid{
return v, ErrInvalid{{ .NullType }}To{{ $item.StdType }}Or{{ $item.AltType }}
}
return v, nil
{{ else }}
return v

{{ end }}
}
{{ end }}

`
