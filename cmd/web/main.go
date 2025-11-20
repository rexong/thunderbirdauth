package main

import (
	"fmt"

	"thunderbird.zap/idp/internal/config"
)

func main() {
	config := config.Init()
	fmt.Println(config.DbPath())
}
