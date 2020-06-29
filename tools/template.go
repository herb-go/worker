package tools

import (
	"text/template"
)

var Template = template.Must(template.New("main").Parse(`package {{.Name}}

//Auto generated code for hiring workers.
//DO NOT EDIT THIS FILE.
import worker "github.com/herb-go/worker"

func init() {
{{- $id := .ID}}
{{- range $key, $value := .Workers }}
	//Worker "{{$id}}.{{$value.Name}}"
	{{- if  not $value.Introduction }}
	//You can add Introduction by add comment in form WORKER({{$value.Name}}):Introduction
	{{- end}}
	worker.Hire("{{$id}}.{{$value.Name}}", &{{$value.Name}}){{- if  $value.Introduction }}.
		WithIntroduction({{$value.Introduction}})
	{{- end}}
{{ end }}
}`))
