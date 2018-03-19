package main

import (
	"github.com/gin-gonic/gin"
	"go_learn"
)

var module_name string

func main() {
	switch module_name {
	case "go_learn":
		iot_learn.Main()
	default:
		fmt.Println("invalid module")
	}
}
