package hosta

//Host碰撞 Tool
import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

var ip []string
var domain []string
var Protocols = []string{"http", "https"}
var wg = sync.WaitGroup{}

func Hosta(argv map[string]string) {
	ipfile := argv["ipfile"]
	domainfile := argv["DomainFile"]
	readFile(ipfile, domainfile)
	process()
}
func readFile(ipfile string, domainfile string) {
	file, err := os.OpenFile(ipfile, os.O_WRONLY|os.O_CREATE, 777)
	//	domainfile, err := os.Open(domainfile)
	if err != nil {
		fmt.Printf("Read File Error")
		os.Exit(0)
	}
	filescan := bufio.NewScanner(file)
	for filescan.Scan() {
		ip = append(ip, filescan.Text())
	}
	file, err = os.OpenFile(domainfile, os.O_WRONLY|os.O_CREATE, 777)
	if err != nil {
		fmt.Printf("Read File Error")
		os.Exit(0)
	}
	filescan = bufio.NewScanner(file)
	for filescan.Scan() {
		domain = append(domain, filescan.Text())
	}
	defer file.Close()
	return
}

func process() {
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}} //忽略https证书
	client := &http.Client{Transport: tr, Timeout: time.Duration(3 * time.Second)}
	for i := 0; i < 2; i++ {
		for j := 0; j < len(ip); j++ {
			Protocol := Protocols[i]
			url := Protocol + "://" + ip[j]
			for k := 0; k < len(domain); k++ {
				wg.Add(1)
				go request(url, client, domain[k])
				wg.Wait()
			}
		}
	}
}
func request(url string, client *http.Client, host string) {
	req, nil := http.NewRequest("GET", url, nil)
	oriresp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request" + url + "Error")
		return
	}
	orilength := oriresp.ContentLength
	req.Header.Add("Host", host)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request" + url + "Error")
		return
	}
	resplength := resp.ContentLength
	if resplength != orilength {
		fmt.Println("URL:" + url + "访问成功," + "Host:" + host)
	} else {
		fmt.Println("URL:" + url + "访问失败")
	}
	wg.Done()
}
