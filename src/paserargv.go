package main

import (
	"fmt"
	"os"
)
import "flag"

type callback struct {
	method string
	argv map[string]string
	scanmode string
}
var Scanmode string

func parse() *callback { //参数解析
	filepath := flag.String("filepath","","input domainfile path")
	function := flag.String("func","None","Select function")
	ports := flag.String("ports","80,443","Input Ports range(default 80,443)")
	IpFile := flag.String("ipfile", "ip.txt", "ipfile,Default dir ./ip.txt")
	DoMainFile := flag.String("domainfile", "domain.txt", "domainfile,Default dir ./domain.txt")
	iprange := flag.String("iprange","","input IP range")
	keyword :=  flag.String("keywords","","input keywords")
	var backInfo callback
	resu := make(map[string]string)
	backInfo.argv =resu
	flag.Parse()//解析参数
	switch *function {
	case "Scan"://获取url GET 信息
			//导入域名
			//fmt.Println(*filepath)
			backInfo.method = "Scan"
			backInfo.argv["ports"] = * ports
			backInfo.argv["filepath"] = *filepath
			if *filepath != ""&& *iprange== "" {
				backInfo.scanmode = "domain"
			} else if  *iprange != ""&&*filepath == ""{
				backInfo.scanmode = "ip"
			} else{
			fmt.Println("Please input domain.txt or iprange")
			os.Exit(0)
		}
			return &backInfo
	case "Dict"://关键字字典
		backInfo.method = "GetDict"
		backInfo.argv["keywords"] = *keyword
		return &backInfo
	case "HostA"://HostA碰撞
		backInfo.method = "HostA"
		backInfo.argv["ipfile"] = *IpFile
		backInfo.argv["DomainFile"] = *DoMainFile
		return &backInfo
 	default:
		fmt.Println("Unkown function select")
		flag.PrintDefaults()
		usage()
		os.Exit(0)
	}
	return nil
}

func usage(){
	usage := "functions List:\n 1.Scan \n 2.Dict \n 3.HostA \n " +
		"Ex:\n " +
		"ToolKit.exe -func Scan -filepath domain.txt (-ports 80,443,...)\n " +
		"ToolKit.exe -func Scan -iprange 192.168.1.1-255\n " +
		"ToolKit.exe -func Dict -keywords baidu,Tenxun\n" +
		"ToolKit.exe -func HostA -ipfile ip.txt(default) -domainfile domain.txt(default)"
	fmt.Println(usage)
}