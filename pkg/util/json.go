package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

// JsonPrint 打印出json格式
func JsonPrint(title string, data interface{}) {
	res, _ := json.Marshal(data)
	fmt.Printf("%v: %+v\n", title, string(res))
}

// JsonFormatPrint 格式化打印出json格式
func JsonFormatPrint(title string, data interface{}) {
	fmt.Printf("%v :\n", title)
	res, _ := json.Marshal(data)
	var out bytes.Buffer
	json.Indent(&out, res, "", "    ")
	//json.Indent(&out, res, "", "\t")
	out.WriteTo(os.Stdout)
	fmt.Printf("\n")
}
