package main

import (
	"fmt"
)

/**
题目：
实现一个函数将字符中空格换为%20

分析：
先算好长度，从尾部开始遍历进行赋值。这道题重点在于字符串是固定内存的，将空格换掉，会导致空间不够放
*/
func main() {
	str := "we are happy"
	r := replaceSpaces([]byte(str))
	fmt.Println(r)
}

/**
1.先创建一个字节buffer，计算出所需要的长度
2.从后面开始遍历进行赋值
*/
func replaceSpaces(strArr []byte) string {
	if len(strArr) <= 0 {
		return string(strArr)
	}

	needLen := len(strArr)
	for _, v := range strArr {
		vt := fmt.Sprintf("%c", v)
		if vt == " " {
			needLen += 2
		}
	}

	var ret []byte
	ret = make([]byte, needLen)
	for i := len(strArr) - 1; i >= 0; i-- {
		v := fmt.Sprintf("%c", strArr[i])
		if v == " " {
			needLen--
			ret[needLen] = '%'
			needLen--
			ret[needLen] = '2'
			needLen--
			ret[needLen] = '0'
		} else {
			needLen--
			ret[needLen] = strArr[i]
		}
	}

	fmt.Println(string(ret))
	//var buf bytes.Buffer
	//buf.WriteString()
	//buf.String()

	return string(ret)
}
