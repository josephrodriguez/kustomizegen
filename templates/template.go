package tmpl

import (
	"bytes"
	"html/template"
)

func InterpolateTemplate(tmpl string, data interface{}) (string, error) {
	// Create a new template and parse the template string
	t := template.New("template")
	t, err := t.Parse(tmpl)
	if err != nil {
		return "", err
	}

	// Create a buffer to store the interpolated output
	var buf bytes.Buffer

	// Execute the template with the data object
	err = t.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	// Return the interpolated output as a string
	return buf.String(), nil
}
