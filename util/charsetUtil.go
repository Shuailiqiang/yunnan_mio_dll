package util

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
)

func Utf8BytesToGB2312(utf8Bytes []byte) string {
	// 使用 GBK 编码器将 UTF-8 字节切片转换为 GB2312 字节切片
	gbkBytes, err := simplifiedchinese.GBK.NewEncoder().Bytes(utf8Bytes)
	if err != nil {
		log.Println("utf8ToGB2312 error  ===>"+string(gbkBytes), err)
		gbkBytes = utf8Bytes
		return string(utf8Bytes)
	}

	// 将 GB2312 字节切片转换为字符串
	gbkStr := string(gbkBytes)
	return gbkStr
}

func Utf8StrToGB2312(utf8Str string) string {
	utf8Bytes := []byte(utf8Str)

	// 使用 GBK 编码器将 UTF-8 字节切片转换为 GB2312 字节切片
	gbkBytes, err := simplifiedchinese.GBK.NewEncoder().Bytes(utf8Bytes)
	if err != nil {
		log.Println("utf8ToGB2312 error  ===>", err)
		gbkBytes = utf8Bytes
		return utf8Str
	}

	// 将 GB2312 字节切片转换为字符串
	gbkStr := string(gbkBytes)
	return gbkStr
}

func GbkStrToUtf8(gbkStr string) string {
	// 将 GBK 编码的字符串转换为字节切片
	gbkBytes := []byte(gbkStr)

	// 使用 GBK 编码器创建 Reader
	gbkReader := transform.NewReader(
		bytes.NewReader(gbkBytes),
		simplifiedchinese.GBK.NewDecoder(),
	)

	// 读取 Reader 中的内容并转换为 UTF-8 编码的字节切片
	utf8Bytes, err := ioutil.ReadAll(gbkReader)
	if err != nil {
		log.Println("GbkStrToUtf8 error ===> ", err)
		return gbkStr
	}

	// 将 UTF-8 编码的字节切片转换为字符串
	utf8Str := string(utf8Bytes)
	return utf8Str
}

func GbkBytesToUtf8(gbkBytes []byte) string {

	// 使用 GBK 编码器创建 Reader
	gbkReader := transform.NewReader(
		bytes.NewReader(gbkBytes),
		simplifiedchinese.GBK.NewDecoder(),
	)

	// 读取 Reader 中的内容并转换为 UTF-8 编码的字节切片
	utf8Bytes, err := ioutil.ReadAll(gbkReader)
	if err != nil {
		log.Println("GbkBytesToUtf8 error ===> ", err)
		return string(gbkBytes)
	}

	// 将 UTF-8 编码的字节切片转换为字符串
	utf8Str := string(utf8Bytes)
	return utf8Str
}
