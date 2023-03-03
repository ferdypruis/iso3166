// Code generated {{.generator}}; DO NOT EDIT.

package iso3166

const (
	_  Country = iota
{{- range .countries}}
	{{.Alpha2}}         // {{.Name}} ({{.Alpha3}}/{{.Numeric}})
{{- end}}
)

var countries = [...]struct {
	// ISO 3166-1 two-letter alphabetic code
	alpha2 string
	// ISO 3166-1 three-letter alphabetic code
	alpha3 string
	// ISO 3166-1 numeric code
	numeric string
	// English name
	name string
}{
{{- range .countries}}
	{{.Alpha2}}: {alpha2: "{{.Alpha2}}", alpha3: "{{.Alpha3}}", numeric: "{{.Numeric}}", name: {{.Name | printf "%q"}}},
{{- end}}
}

var fromAlpha2 = map[string]Country{
{{- range .countries}}
    "{{.Alpha2}}": {{.Alpha2}},
{{- end}}
}

var fromAlpha3 = map[string]Country{
{{- range .countries}}
    "{{.Alpha3}}": {{.Alpha2}},
{{- end}}
}

var fromNumeric = map[string]Country{
{{- range .countries}}
    "{{.Numeric}}": {{.Alpha2}},
{{- end}}
}

var fromName = map[string]Country{
{{- range .countries}}
    "{{.Name}}": {{.Alpha2}},
{{- end}}
}
