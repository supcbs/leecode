package main

import "fmt"

/**

题目:
给定一个单词数组和一个长度 maxWidth，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
你应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，
使得每行恰好有 maxWidth 个字符。
要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
文本的最后一行应为左对齐，且单词之间不插入额外的空格。

说明:
单词是指由非空格字符组成的字符序列。
每个单词的长度大于 0，小于等于 maxWidth。
输入单词数组 words 至少包含一个单词。

示例:
输入:
words = ["This", "is", "an", "example", "of", "text", "justification."]
maxWidth = 16
输出:
[
   "This    is    an",
   "example  of text",
   "justification.  "
]

示例 2:
输入:
words = ["What","must","be","acknowledgment","shall","be"]
maxWidth = 16
输出:
[
  "What   must   be",
  "acknowledgment  ",
  "shall be        "
]
解释: 注意最后一行的格式应为 "shall be    " 而不是 "shall     be",
     因为最后一行应为左对齐，而不是左右两端对齐。
     第二行同样为左对齐，这是因为这行只包含一个单词。

示例 3:
输入:
words = ["Science","is","what","we","understand","well","enough","to","explain",
         "to","a","computer.","Art","is","everything","else","we","do"]
maxWidth = 20
输出:
[
  "Science  is  what we",
  "understand      well",
  "enough to explain to",
  "a  computer.  Art is",
  "everything  else  we",
  "do                  "
]

*/

func main() {
	words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	words = []string{"Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"}
	maxWidth := 20
	r := fullJustify(words, maxWidth)
	fmt.Println(r)
}

/**
方法1:

这道题主要是边界一定要算清楚

1.需要使用贪心算法，尽可能多的放一行，所以需要先确定当前行能放几个
2.如果当前行单词是一个或者是最后一行，则左对齐，后面补0
3.属于正常行，则需要算出单词直接的空格diff，由于go是向下取整，所以还需要算出其余数，加到左边
4.双指针的start此时表为last进行下一轮遍历

时间复杂度：O(n)
空间复杂度：O(1)
*/
func fullJustify(words []string, maxWidth int) []string {
	if len(words) <= 0 || maxWidth <= 0 {
		return []string{}
	}

	size := len(words)
	start := 0
	res := make([]string, 0)
	for start < size {
		// 确定当前行能放几个
		var length, last = len(words[start]), start + 1
		for last < size {
			// 这里加1是因为单词直接间隔必须有空格
			if length+len(words[last])+1 > maxWidth {
				break
			}

			length += len(words[last]) + 1
			last++
		}
		fmt.Println(start,last)

		// 判断是否是最后一行或者只有一个单词
		diff := last - 1 - start
		str := string(words[start])
		if last == size || diff == 0 {
			for i := start + 1; i < last; i++ {
				str += " " + string(words[i])
			}

			for len(str) < maxWidth {
				str += " "
			}
		} else {
			// 算出单词间隔需要补的空格数
			space := (maxWidth - length) / diff
			// 算出余数
			c := (maxWidth - length) % diff
			if start == 10 {
				fmt.Println(diff, space,c,length)
			}
			for i := start + 1; i < last; i++ {
				for j := space; j > 0; j-- {
					str += " "
				}

				// 为了均匀分配这里要用if，因为是余数所以肯定会在遍历玩之前分配完
				if c > 0 {
					str += " "
					c--
				}
				str += " " + string(words[i])
			}
		}

		res = append(res, str)
		start = last

	}
	return res
}


