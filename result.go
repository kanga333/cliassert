package cliassert

import (
	"bytes"
	"html/template"
)

// Result is the struct that shows assertion result.
type Result struct {
	exitStatus string
	stdout     string
	stderr     string
	successes  []string
	failures   []string
}

var resultFormat = `
{{- range $i, $v := .Failures -}}
[failure] {{ $v }}.
{{ end -}}
`

var detailFormat = `{{ "" -}}
[exit-status] {{ .ExitStatus }}

[stdout]
{{ .Stdout }}
[stderr]
{{ .Stderr }}
---
{{ range $i, $v := .Successes -}}
[success] {{ $v }}.
{{ end -}}
{{ .Result }}
{{ .NumOfCases }} cases, {{ .NumOfFailures }} failures.
`

// Show shows the standard result.
func (r *Result) Show() (string, error) {
	var buf bytes.Buffer

	tmpl, err := template.New("result").Parse(resultFormat)
	if err != nil {
		return "", err
	}
	err = tmpl.Execute(&buf, struct {
		Failures []string
	}{
		r.failures,
	})
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// ShowDetails shows the detailed result.
func (r *Result) ShowDetails() (string, error) {
	var buf bytes.Buffer

	tmpl, err := template.New("detail").Parse(detailFormat)
	if err != nil {
		return "", err
	}

	result, err := r.Show()
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(&buf, struct {
		ExitStatus    string
		Stdout        string
		Stderr        string
		Successes     []string
		Result        string
		NumOfCases    int
		NumOfFailures int
	}{
		r.exitStatus,
		r.stdout,
		r.stderr,
		r.successes,
		result,
		r.countCases(),
		r.countFailures(),
	})
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Stdout returns the stdout of the result.
func (r *Result) Stdout() string {
	return r.stdout
}

func (r *Result) countCases() int {
	return len(r.successes) + len(r.failures)
}

func (r *Result) countFailures() int {
	return len(r.failures)
}
