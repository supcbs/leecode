package main


import (
	"strconv"
	"fmt"
)

func permutation(s string) []string {
	if 1 == len(s) {
		return []string{s}
	}
	var results []string
	for i := 0; i < len(s); i++ {
		rightPart := s[:i] + s[i+1:]
		for _, j := range permutation(rightPart) {
			result := string(s[i]) + j
			results = append(results, result)
		}
	}
	return results
}

func getPermutation(n int) []string {
	s := ""
	for i := 1; i <= n; i++ {
		s += strconv.Itoa(i)
	}
	return permutation(s)
}

func main() {
	fmt.Println(getPermutation(4))
}
