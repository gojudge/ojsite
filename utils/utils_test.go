package utils

import (
	"fmt"
	"testing"
)

func Test_DecodeIoData(t *testing.T) {
	inputs := "input1\ninput2\ninput3"
	outputs := "output1\noutput2\noutput3"

	if num, json, err := DecodeIoData(inputs, outputs); err != nil {
		t.Error(err)
	} else {
		fmt.Println(json, "[num]", num)
	}
}

func Test_Tags(t *testing.T) {
	input := `测试，标签,分类，test，with, space`

	midOutput := TagsCheck(input)
	fmt.Println(midOutput)

	if midOutput != "测试,标签,分类,test,with,space" {
		t.Error("格式化错误")
	}

	tagArr := TagsParse(midOutput)
	fmt.Println(tagArr)

	if tagArr[0] != "测试" {
		t.Error("解析错误")
	}
}
