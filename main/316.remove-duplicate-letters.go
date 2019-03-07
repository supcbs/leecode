package main

import (
	"fmt"
	"test/leecode/utils"
)

/**

题目:
给定一个仅包含小写字母的字符串，去除字符串中重复的字母，使得每个字母只出现一次。
需保证返回结果的字典序最小（要求不能打乱其他字符的相对位置）。

示例 1:
输入: "bcabc"
输出: "abc"

示例 2:
输入: "cbacdcbc"
输出: "acdb"

*/

func main() {
	str := "c"
	str = "cbbbcaa"
	r := removeDuplicateLetters(str)
	fmt.Println(r)
}

/**
方法1:

首先需要用到map和栈，首先将所有的字母存到map中
然后遍历字符串，依次放到栈中，当发现当前的字母比栈中最后一个字母x小并且后续还存在x字母（map计数>0）
则将最后一个出栈，以此类推。
最后输出栈，倒着拼接，即为目标字符串

时间复杂度：O(n)
空间复杂度：O(1)
*/
func removeDuplicateLetters(s string) string {
	sLen := len(s)
	if sLen <= 0 {
		return ""
	}

	var stack utils.Stack
	temp := make([]byte, 26)
	for i := 0; i < sLen; i++ {
		temp[s[i]-'a']++
	}

	index := 0
	for _,v := range temp {
		if v > 0{
			index ++
		}
	}

	visit := make([]byte, 26)
	for i := 0; i < sLen; i++ {
		cur := s[i]


		for !stack.IsEmpty() && cur < stack.Last().(byte) && temp[stack.Last().(byte)-'a'] > 0 && visit[cur-'a'] == 0 {
			pop, _ := stack.Pop()
			popByte := pop.(byte)
			visit[popByte-'a'] = 0
			//temp[popByte-'a']--
			fmt.Println("pop===",pop)
		}

		if visit[cur-'a'] == 1 {
			temp[cur-'a']--
			continue
		}

		fmt.Println("push===",cur)
		stack.Push(cur)
		visit[cur-'a'] = 1
		temp[cur-'a']--

		fmt.Println(stack, temp)
		//os.Exit(1)
	}

	res := make([]byte, index)
	for !stack.IsEmpty() {
		p, _ := stack.Pop()
		pByte := p.(byte)
		index--
		res[index] = pByte
	}

	return string(res)
}

func removeDuplicateLettersOK(s string) string {
	var cnt ['z' - 'a' + 1]int
	var visited ['z' - 'a' + 1]bool
	for _, c := range s {
		cnt[c-'a'] += 1
	}

	//for _, c := range s {
	//	cnt[c-'a']--
	//}
	stack := make([]uint8, 0)
	for i, next := range s {
		//dumpStats(stack,visited,cnt)
		//fmt.Printf("stack %v\n", stack)
		//fmt.Printf("visit %+v\n", visited)
		//fmt.Printf("cnt %+v\n", cnt)
		if len(stack) == 0 {
			visited[next-'a'] = true
			stack = append(stack, s[i])
			cnt[next-'a']--
			continue
		}

		for ; len(stack) > 0; {

			top := int32(stack[len(stack)-1])

			if top > next && cnt[top-'a'] > 0 && !visited[next-'a'] {
				stack = stack[:len(stack)-1]
				visited[top-'a'] = false
			} else {
				break
			}
		}
		if !visited[next-'a'] {
			stack = append(stack, uint8(next))
			visited[next-'a'] = true
		}
		cnt[next-'a']--
	}
	ans := ""
	for i := range stack {
		ans += fmt.Sprintf("%c", stack[i])
	}
	return ans
}
