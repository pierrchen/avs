package tmpl

// Vendorsetup is the template for vendorsetup.mk
const Initrc = `
{{- range .Imports -}}
import {{.}}
{{end}}

{{- range .Actions }}
on {{.Triggers}}
    {{- range .Commands}}
    {{.}}
    {{- end}}
{{end}}


{{- range .Services}}
service {{.Name}} {{.Path}} {{.Args}}
   {{- range .Options}}
    {{.}}
   {{- end }}
{{- end }}`
