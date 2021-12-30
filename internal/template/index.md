# Architecture Decision Records

| #    | Date       | Status     | Title      |
| ---- | ---------- | ---------- | ---------- |
{{- range .}}
| {{printf "%04d" .Number}} | {{.Date}} | {{.Status}} | [{{.Title}}]({{.Filename}}) |
{{- end}}
