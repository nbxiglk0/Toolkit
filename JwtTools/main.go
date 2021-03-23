package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"net/url"
	"strings"
)
func main(){
	action := flag.String("action","","select action")
	content := flag.String("content","","input content")
	secret := flag.String("secret","","input secret key")
	flag.Parse()
	if *action =="encode" {
		result := encode(*content,*secret)
		fmt.Println("JWT Encode"+result)
	}
	if *action =="decode"{
		result := decode(*content)
		fmt.Println("JWT Decode:"+result)

	}else {
		flag.PrintDefaults()
	}
}

func encode(content string,secret string) string{
	decodeContent,_ := base64.StdEncoding.DecodeString(content)
	headers := strings.Split(string(decodeContent),".")[0]
	payload := Base64UrlSafeEncode([]byte(strings.Split(string(decodeContent),".")[1]))
	message := Base64UrlSafeEncode([]byte(headers))+"."+payload
		sha :=  hmac.New(sha256.New,[]byte(secret))
		sha.Write([]byte(message))
		sign := sha.Sum(nil)
		return message+"."+string(sign)
}

func decode(content string) string {
	splitcontent := strings.Split(content,".")
	types,_ := base64.StdEncoding.DecodeString(splitcontent[0])
	payload,_ := base64.StdEncoding.DecodeString(splitcontent[1])
	sign := splitcontent[2]
	decodecontent,_ := url.QueryUnescape(string(types)+string(payload))
	return decodecontent+"}"+" Sign:"+sign
}

func Base64UrlSafeEncode(source []byte) string {
	byteArr := base64.StdEncoding.EncodeToString(source)
	safeUrl := strings.Replace(string(byteArr), "/", "_", -1)
	safeUrl = strings.Replace(safeUrl, "+", "-", -1)
	safeUrl = strings.Replace(safeUrl, "=", "", -1)
	return safeUrl
}
