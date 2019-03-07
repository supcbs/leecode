package main

import (
	"fmt"
	"unsafe"
)

/**

题目:
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);

示例 1:
输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"

示例 2:
输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G

*/

func main() {
	str := "ABC"
	numRows := 1
	r := convert(str,numRows)
	fmt.Println(r)
}

/**
方法1:

关键点在于两个
1.创建的数组的rows数应该为 numsRows和len(str)取小的
2.需要两个表识：一个是当前的行数，另一个是是否转向

转向发生在两个地方，当row=0和numsRow-1的的时候发生转向，
转向涉及到当前行数是+1还是-1

时间复杂度：O(n)
空间复杂度：O(1)
*/
func convert(s string, numRows int) string {
	sLen := len(s)

	if sLen <= 0 || numRows == 0 {
		return ""
	}

	curRow := 0
	goDown := false

	arr := make([][]rune,numRows)
	for _,v := range s {
		fmt.Println(curRow)
		arr[curRow] = append(arr[curRow],v)
		if curRow == 0 || curRow == numRows - 1 {
			goDown = !goDown
		}
		if goDown {
			curRow ++
		} else {
			curRow --
		}

		if curRow > numRows-1 {
			curRow = numRows-1
		}
		if curRow < 0  {
			curRow =0
		}
	}

	ret := make([]byte,0)
	for _,rows := range arr {
		for _,v := range rows {
			ret = append(ret, byte(v))
		}
	}
	return string(ret)
}

func convert2(s string, numRows int) string {

	sLen := len(s)
	if sLen == 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}
	resultArray := make([]byte,sLen)
	t := numRows*2 - 2
	length := len(s)
	centerNum := int(length/t) + 1
	if length - (centerNum -1 )* t >= t / 2 {
		centerNum++
	}
	centerIndexArray := make([]int,centerNum)
	m := 0
	for i := 0; i < centerNum; i++ {
		j := i * t
		centerIndexArray[i] = j
		if j < length {
			resultArray[i] = s[j]
			m++
		}
	}
	//fmt.Println(result_array)
	for i := 1; i < numRows - 1; i++ {
		for _, c := range centerIndexArray {
			L := c - i
			R := c + i
			if L > 0 && L < length {
				resultArray[m] = s[L]
				m++
			}
			if R < length {
				resultArray[m] = s[R]
				m++
			}
		}
	}
	//fmt.Println(result_array)
	for _, c := range centerIndexArray {
		i := c + numRows - 1
		if i < length {
			resultArray[m] = s[i]
			m++
		}
	}

	return *(*string)(unsafe.Pointer(&resultArray))
}

func convert3(s string, numRows int) string {
	sRunes := []rune(s)
	lenStr := len(sRunes)
	if s == "" || lenStr < numRows || numRows < 2 {
		return s
	}

	retRunes, gapNum := make([]rune, 0, lenStr), 2*(numRows-1)
	for i := 0; i < numRows; i++ {
		nextGap := 2 * (numRows - i - 1)
		for index, step := i, 0; index < lenStr; step++ {
			retRunes = append(retRunes, sRunes[index])
			if nextGap == 0 || nextGap == gapNum {
				index += gapNum
				continue
			}

			if step%2 == 0 {
				index = index + nextGap
			} else {
				index = index + gapNum - nextGap
			}
		}
	}

	return string(retRunes)
}
