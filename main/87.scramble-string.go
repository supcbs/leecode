package main

import "fmt"

/**

题目:
给定一个字符串 s1，我们可以把它递归地分割成两个非空子字符串，从而将其表示为二叉树。
下图是字符串 s1 = "great" 的一种可能的表示形式。

    great
   /    \
  gr    eat
 / \    /  \
g   r  e   at
           / \
          a   t
在扰乱这个字符串的过程中，我们可以挑选任何一个非叶节点，然后交换它的两个子节点。
例如，如果我们挑选非叶节点 "gr" ，交换它的两个子节点，将会产生扰乱字符串 "rgeat" 。

    rgeat
   /    \
  rg    eat
 / \    /  \
r   g  e   at
           / \
          a   t
我们将 "rgeat” 称作 "great" 的一个扰乱字符串。
同样地，如果我们继续将其节点 "eat" 和 "at" 进行交换，将会产生另一个新的扰乱字符串 "rgtae" 。

    rgtae
   /    \
  rg    tae
 / \    /  \
r   g  ta  e
       / \
      t   a
我们将 "rgtae” 称作 "great" 的一个扰乱字符串。
给出两个长度相等的字符串 s1 和 s2，判断 s2 是否是 s1 的扰乱字符串。

示例 1:
输入: s1 = "great", s2 = "rgeat"
输出: true

示例 2:
输入: s1 = "abcde", s2 = "caebd"
输出: false

*/

func main() {
	s1 := "great"
	s2 := "rgeat"
	r := isScramble(s1,s2)
	fmt.Println(r)
}

/**
方法1:

首先判断字符串是否一样，一样直接返回
其次判断字符串是否相等，不相等直接返回，通过map++--来判断
最后通过递归的方式来进行判断

gr rg gr
想要相等必须满足[0:i] = [0:j] && [i:] = [j:]
          或者[0:i] =  [size-i:] && [i:] = [0:size-i]

时间复杂度：O(n)
空间复杂度：O(1)
*/
func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	if len(s1) != len(s2) {
		return true
	}

	// 因为ascii码就只有256位
	check:=make([]int, 256)
	for i:=0;i<len(s1);i++{
		check[s1[i]]++
		check[s2[i]]--
	}

	for _,v := range check {
		if v != 0 {
			return false
		}
	}

	for i:=1;i<len(s1);i++{
		if isScramble(s1[0:i],s2[0:i]) && isScramble(s1[i:],s2[i:]) {
			return true
		}
		if isScramble(s1[0:i],s2[len(s2)-i:]) && isScramble(s1[i:],s2[0:len(s2)-i]) {
			return true
		}
	}

	return false
}
