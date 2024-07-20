package main

import (
	wanruncmd "github.com/wanrun-develop/wanrun/cmd"
	"github.com/wanrun-develop/wanrun/configs"
)

func init() {
	configs.LoadConfig()
}

func main() {
	wanruncmd.Main()
}
