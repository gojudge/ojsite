package utils

import (
	// "fmt"
	"github.com/gogather/com/log"
	"github.com/russross/blackfriday"
	"html"
)

func Trace(format string, v ...interface{}) {
	log.Bluef(format, v...)
}

func Warn(format string, v ...interface{}) {
	log.Warnf(format, v...)
}

func Markdown2HTML(content string) string {
	content = html.EscapeString(content)
	output := blackfriday.MarkdownCommon([]byte(content))
	return string(output)
}
