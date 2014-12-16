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
