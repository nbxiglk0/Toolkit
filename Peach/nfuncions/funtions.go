package nfuncions

import (
	"fmt"
	"net"
	"reflect"
	"strconv"
	"time"
)
//是否包含
//GBK OR UTF-8 字符集判断
func IsContain(item interface{}, items interface{}) bool{ //判断slice是否包含某个item
	switch reflect.TypeOf(items).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(items)
		for i :=0; i < s.Len(); i++{
			if reflect.DeepEqual(item, s.Index(i).Interface()){
				return true
			}
		}
	}
	return false
}
func Duplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}
//判断是否GBK
func IsGBK(data []byte) bool {
	length := len(data)
	var i int = 0
	for i < length {
		//fmt.Printf("for %x\n", data[i])
		if data[i] <= 0xff {
			//编码小于等于127,只有一个字节的编码，兼容ASCII吗
			i++
			continue
		} else {
			//大于127的使用双字节编码
			if  data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i + 1] >= 0x40 &&
				data[i + 1] <= 0xfe &&
				data[i + 1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}
//判断是否是UTF-8
func IsUtf8(data []byte) bool {
	for i := 0; i < len(data);  {
		if data[i] & 0x80 == 0x00 {
			// 0XXX_XXXX
			i++
			continue
		} else if num := preNUm(data[i]); num > 2 {
			// 110X_XXXX 10XX_XXXX
			// 1110_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_0XXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_10XX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_110X 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// preNUm() 返回首个字节的8个bits中首个0bit前面1bit的个数，该数量也是该字符所使用的字节数
			i++
			for j := 0; j < num - 1; j++ {
				//判断后面的 num - 1 个字节是不是都是10开头
				if data[i] & 0xc0 != 0x80 {
					return false
				}
				i++
			}
		} else  {
			//其他情况说明不是utf-8
			return false
		}
	}
	return true
}
func preNUm(data byte) int {
	str := fmt.Sprintf("%b", data)
	var i int = 0
	for i < len(str) {
		if str[i] != '1' {
			break
		}
		i++
	}
	return i
}
//判断端口是否tls
func Istls(ip string , port int)bool{
	t:=[]byte{0x16,0x03,0x01,0x00,0xb5,0x01,0x00,0x00,0xb1,0x03,0x03,0xb2,0xd3,0x4d,0xfd,0x63,0xbe,0x89,0xdb,0xe5,0x46,0xcc,0xaf,0x39,0x6e,0xba,0x63,0x63,0x75,0xce,0x30,0xda,0xe0,0x4f,0xab,0xa2,0x3e,0x50,0xea,0x41,0x20,0x10,0xc4,0x00,0x00,0x18,0xc0,0x2b,0xc0,0x2f,0xc0,0x2c,0xc0,0x30,0xc0,0x13,0xc0,0x14,0x00,0x9c,0x00,0x9d,0x00,0x2f,0x00,0x35,0x00,0x0a,0x00,0xff,0x01,0x00,0x00,0x70,0x00,0x00,0x00,0x15,0x00,0x13,0x00,0x00,0x10,0x77,0x77,0x77,0x2e,0x73,0x6f,0x2d,0x63,0x6f,0x6f,0x6c,0x73,0x2e,0x63,0x6f,0x6d,0x00,0x0b,0x00,0x04,0x03,0x00,0x01,0x02,0x00,0x0a,0x00,0x06,0x00,0x04,0x00,0x17,0x00,0x18,0x00,0x23,0x00,0x00,0x00,0x0d,0x00,0x20,0x00,0x1e,0x06,0x01,0x06,0x02,0x06,0x03,0x05,0x01,0x05,0x02,0x05,0x03,0x04,0x01,0x04,0x02,0x04,0x03,0x03,0x01,0x03,0x02,0x03,0x03,0x02,0x01,0x02,0x02,0x02,0x03,0x00,0x05,0x00,0x05,0x01,0x00,0x00,0x00,0x00,0x00,0x0f,0x00,0x01,0x01,0x00,0x10,0x00,0x0b,0x00,0x09,0x08,0x68,0x74,0x74,0x70,0x2f,0x31,0x2e,0x31}

	Target :=ip
	port = port
	Time, _ := time.ParseDuration("2s")
	conn, err := net.DialTimeout("tcp", Target+":"+strconv.Itoa(port), Time )

	if err != nil {
		fmt.Println("ERR::" + strconv.Itoa(port) + ">" + err.Error())
		return false
	}
	conn.Write(t)
	recvBuf := make([]byte, 2048)
	conn.SetReadDeadline(time.Now().Add(time.Second * 2))
	_, err = conn.Read(recvBuf[:])
	if err != nil {
		//fmt.Println("ERR::" + strconv.Itoa(port) + ">" + err.Error())
		return false
	}
	conn.SetReadDeadline(time.Time{})
/*	fmt.Println("tlsinfo:")
	fmt.Println( string(recvBuf[:]))*/
	if string(recvBuf[0:4]) == string([] byte {22,3,3,0}) {
		conn.Close()
		return true
	}else{
		conn.Close()
		return false
	}
}
//合并切片
func MergeSilce(s1 []string, s2 []string) []string{
	slice := make([]string,len(s1)+len(s2))
	copy(slice,s1)
	copy(slice[len(s1):],s2)
	return slice
}