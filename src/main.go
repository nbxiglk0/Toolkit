package main

import (
	"ToolKit/src/HostA"
	"ToolKit/src/basicinfo"
	"ToolKit/src/getdict"
	"fmt"
	"time"
)

func main() {
	//	gui.Main()
	start := time.Now()
	callback := parse()
	switch callback.method {
	case "scan":
		basicinfo.Main(callback.argv, callback.scanmode)
	case "GetDict":
		getdict.Main(callback.argv["keywords"])
	case "HostA":
		hosta.Hosta(callback.argv)
	}
	cost := time.Since(start)
	fmt.Printf("The Task cost time %s", cost)
}
