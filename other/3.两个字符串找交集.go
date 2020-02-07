package main

import "fmt"

/**
1.用map存储
2.遍历第二个字符串
 */
func main() {
	str1 := "abc1234def789 0g"
	str2 := "ba 209834cdxyz"
	mStr := make([]byte,0)

	fmt.Println(str1[0])
	fmt.Println([]byte(str1)[0])
	temp := make(map[byte]int)
	for i:=0;i<len(str1);i++{
		temp[str1[i]] = 1
	}
	fmt.Println(temp)
	for i:=0;i<len(str2);i++{
		if _,ok := temp[str2[i]]; ok {
			mStr = append(mStr, str2[i])
		}
	}

	fmt.Println(string(mStr))
}

//func main2() {
//	str1 := "abc1234def789 0g"
//	str2 := "ba 209834cdxyz"
//	mStr := make([]byte,0)
//
//	for i:=0;i<len(str2);i++{
//		if _,ok := temp[str2[i]]; ok {
//			mStr = append(mStr, str2[i])
//		}
//	}
//
//	fmt.Println(string(mStr))
//}

//func quickSort(s string, left int, right int) string {
//	if left >= right {
//		return s
//	}
//
//	x,y := left,right
//	k := s[left]
//
//}
