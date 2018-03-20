package main

import (
	"fmt"
	"iot_learn"
)

var module_name string

func main() {
	switch module_name {
	case "iot_learn":
       iot_learn.Main()
	default:
		fmt.Println("invalid module")
	}
}
