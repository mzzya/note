package main

import (
	"fmt"
)

func main() {
	mapA := []string{"abcabcbb", "bbbbb", "pwwkew", "au", "aab", "ohvhjdml"}
	for _, m := range mapA {
		fmt.Printf("m=%s\t %d\n", m, lengthOfLongestSubstring(m))
	}
}
func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	max := 0
	var temp = make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(temp) == 0 {
			temp = append(temp, s[i])
			continue
		}
		index := -1
		exist := false
		for j, b := range temp {
			if b == s[i] {
				index = j
				exist = true
				break
			}
		}
		if exist {
			if len(temp) > max {
				max = len(temp)
			}
			if len(temp) > index+1 {
				temp = temp[index+1:]
			} else {
				temp = make([]byte, 0)
			}
		}
		temp = append(temp, s[i])
	}
	if len(temp) > max {
		max = len(temp)
	}
	return max
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ll1, ll2 := l1, l2
	result := new(ListNode)
	var tempResult = result
	var carry = 0
	for ll1 != nil || ll2 != nil {
		var ll1v, ll2v int
		if ll1 != nil {
			ll1v = ll1.Val
		}
		if ll2 != nil {
			ll2v = ll2.Val
		}
		var s = carry + ll1v + ll2v
		carry = s / 10
		tempResult.Val = s % 10
		//fmt.Printf("s%d\tc%d ll1N%#v ll2N%#v\n", s, carry, ll1.Next != nil, ll2.Next != nil)
		if ll1 != nil {
			ll1 = ll1.Next
		}
		if ll2 != nil {
			ll2 = ll2.Next
		}
		if ll1 != nil || ll2 != nil || carry == 1 {
			tempResult.Next = new(ListNode)
			tempResult = tempResult.Next
		}
	}
	if carry == 1 && tempResult != nil {
		tempResult.Val = carry
	}
	return result
}
