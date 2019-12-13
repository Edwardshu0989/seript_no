package util

import (
	"math/rand"
	"strings"
	"time"
)

// 前置的0使用该字符补位
const coverChar = "P"

// 字符集，去除数字中的0和1，字母中的O和I，和补位字符P
var charset = []string{"H", "A", "2", "T", "N", "6", "D", "Z", "X", "J", "U", "F", "R", "4", "9", "C", "7",
	"5", "K", "3", "M", "W", "V", "E", "8", "S", "Y", "L", "B", "G", "Q"}

// 字符集长度
var charsetLen = int64(len(charset))

// 表示范围 = charsetLen^codeLen-1
// 例如6位邀请码的表示范围 31^6-1 = 887503680，约8.8亿
// 7位邀请码的表示范围 31^7-1 = 27512614110，约275亿
// 8位邀请码的表示范围 31^8-1 = 852891037440，约8528亿

// 将id转换为邀请码
// 输入 用户id 和 邀请码长度 得到指定长度的邀请码
// 如果用户id超出了指定长度邀请码的表示范围，将返回空字符串
func ToInviteCode(id int64, length int64) string {
	var buf = make([]string, length)
	pos := length - 1
	// 从最后以为往前计算
	for id > 0 {
		// 取余
		index := id % charsetLen
		// 对应位赋值
		buf[pos] = charset[index]
		// 商值继续计算
		id /= charsetLen
		// 前移一位
		pos -= 1

		// 计算完毕退出
		if id == 0 {
			break
		}
		// 溢出
		if pos < 0 {
			return ""
		}
	}
	if pos >= 0 {
		buf[pos] = coverChar
	}
	rand.Seed(time.Now().UnixNano())
	for i := pos - 1; i >= 0; i-- {
		if buf[i] == "" {
			buf[i] = charset[rand.Intn(int(charsetLen))]
		}
	}
	return strings.Join(buf, "")
}

// 将邀请码转换为id
// 输入 邀请码 和 邀请码长度 解析出用户id
// 如果输入的邀请码长度和指定长度不一致，将返回0
func ToId(code string, length int64) int64 {
	if int64(len(code)) != length {
		return 0
	}
	var codeArr = strings.Split(code, "")
	var id int64
	var i int64
	var j int64
	var base int64 = 1
	for i = length - 1; i >= 0; i-- {
		if codeArr[i] == coverChar {
			break
		}
		for j = 0; j < charsetLen; j++ {
			if codeArr[i] == charset[j] {
				id += j * base
			}
		}
		base *= charsetLen
	}
	return id
}
