package main

import (
	"ToolKit/src/BackupDic"
	"ToolKit/src/HostA"
	"ToolKit/src/HttpScan"
	"ToolKit/src/nfuncions"
	"fmt"
	"time"
)

func main() {
	//	gui.Main()
	start := time.Now()
	callback := nfuncions.Parse()
	switch callback.Method {
	case "HttpScan":
		HttpScan.Main(callback.Argv, callback.Scanmode)
	case "BackupDic":
		BackupDic.Main(callback.Argv["keywords"])
	case "HostA":
		hosta.Hosta(callback.Argv)
	}
	cost := time.Since(start)
	fmt.Printf("The Task cost time %s", cost)
}
