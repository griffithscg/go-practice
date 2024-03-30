package adventure_view

import (
	"os"
)

// Driver function, takes in an error object and builds up the error page
func BuildErrorPage(err error) error {
	msg := err.Error()
	templateStr := "{{define \"message\"}} " + msg + "{{end}}"
	return os.WriteFile("error-page.html", []byte(templateStr), 0644)
}
