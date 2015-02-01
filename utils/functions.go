package utils

import (
	// "fmt"
	"github.com/gogather/com"
	"github.com/gogather/com/log"
	"github.com/russross/blackfriday"
	"html"
	"strings"
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

// decode input and output data, "\n" seg
func DecodeIoData(inputs string, outputs string) (num int, json string, err error) {
	in := strings.Split(inputs, "\n")
	out := strings.Split(outputs, "\n")

	if len(in) > len(out) {
		num = len(out)
	} else {
		num = len(in)
	}

	ins := in[0:num]
	outs := out[0:num]

	data := map[string]interface{}{
		"input":   ins,
		"outputs": outs,
	}

	json, err = com.JsonEncode(data)

	return num, json, err
}

// parse and check tags
func TagsCheck(tags string) string {
	tags = strings.Replace(tags, " ", "", -1)
	tags = strings.Replace(tags, "，", ",", -1)
	return tags
}

// parse tags
func TagsParse(tags string) []string {
	return strings.Split(tags, ",")
}

// pack client message
func MsgPack(data map[string]interface{}) string {
	json, err := com.JsonEncode(data)
	if err != nil {
		return ""
	} else {
		return json
	}
}
