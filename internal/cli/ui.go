package cli

import (
	"fmt"

	"github.com/fatih/color"
)

// var SuccessMsg = color.New(color.FgGreen).PrintlnFunc()
// var ErrorMsg = color.New(color.FgRed).PrintlnFunc()
// var Warning = color.New(color.FgYellow).PrintlnFunc()
// var Title = color.New(color.Bold, color.FgCyan).SprintFunc()
// var Status = color.New(color.FgGreen).SprintFunc()
var Meta = color.New(color.FgHiBlack).SprintFunc()

// Generic builder
func Style(attrs ...color.Attribute) func(a ...interface{}) string {
	return color.New(attrs...).SprintFunc()
}

// Print version
func PrintStyle(attrs ...color.Attribute) func(a ...interface{}) {
	return color.New(attrs...).PrintlnFunc()
}

var (
	Bold      = Style(color.Bold)
	Dim       = Style(color.Faint)
	Underline = Style(color.Underline)

	Green  = Style(color.FgGreen)
	Red    = Style(color.FgRed)
	Yellow = Style(color.FgYellow)
	Cyan   = Style(color.FgCyan)
	Gray   = Style(color.FgHiBlack)

	BoldCyan  = Style(color.Bold, color.FgCyan)
	BoldGreen = Style(color.Bold, color.FgGreen)
)

func Success(args ...any) {
	fmt.Println(Green("✔ "), fmt.Sprint(args...))
}

func Error(args ...any) {
	fmt.Println(Red("✖ "), fmt.Sprint(args...))
}

func Warning(args ...any) {
	fmt.Println(Yellow("⚠ "), fmt.Sprint(args...))
}

func Info(args ...any) {
	fmt.Println(Cyan("ℹ "), fmt.Sprint(args...))
}

func Section(title string) string {
	return BoldCyan(title)
}

func Muted(text string) string {
	return Gray(text)
}

func PrintStatusWithColorAndText(code int) string {
	var text string

	switch code {
	case 200:
		text = "200 OK"
	case 201:
		text = "201 Created"
	case 204:
		text = "204 No Content"
	case 400:
		text = "400 Bad Request"
	case 401:
		text = "401 Unauthorized"
	case 403:
		text = "403 Forbidden"
	case 404:
		text = "404 Not Found"
	case 500:
		text = "500 Internal Server Error"
	default:
		text = fmt.Sprintf("%d", code)
	}

	switch {
	case code >= 200 && code < 300:
		return Green(text)
	case code >= 300 && code < 400:
		return Cyan(text)
	case code >= 400 && code < 500:
		return Yellow(text)
	default:
		return Red(text)
	}
}

func statusText(code int) string {
	switch code {
	case 200:
		return "200 OK"
	case 201:
		return "201 Created"
	case 204:
		return "204 No Content"
	case 400:
		return "400 Bad Request"
	case 401:
		return "401 Unauthorized"
	case 403:
		return "403 Forbidden"
	case 404:
		return "404 Not Found"
	case 500:
		return "500 Internal Server Error"
	default:
		return fmt.Sprintf("%d", code)
	}
}

func FormatStatus(code int) string {
	text := statusText(code)

	switch {
	case code >= 200 && code < 300:
		return Green(text)
	case code >= 300 && code < 400:
		return Cyan(text)
	case code >= 400 && code < 500:
		return Yellow(text)
	default:
		return Red(text)
	}
}
