package solution

import (
	"fmt"
	"testing"
)

func newListNode(vals []int) *ListNode {
	l := &ListNode{Val: vals[0]}
	p := l
	for _, v := range vals[1:] {
		t := &ListNode{Val: v}
		p.Next = t
		p = t
	}
	return l
}

func printListNode(l *ListNode) {
	for ; l != nil; l = l.Next {
		fmt.Printf("%v ", l.Val)
	}
	fmt.Println()
}

func Test_isPalindrome_1(t *testing.T) {
	l := newListNode([]int{1, 2})
	expected := false
	result := isPalindrome(l)
	if result != expected {
		t.Errorf("expected: %v, isPalindrome: %v", expected, result)
	}
}

func Test_isPalindrome_2(t *testing.T) {
	l := newListNode([]int{1, 2, 2, 1})
	expected := true
	result := isPalindrome(l)
	if result != expected {
		t.Errorf("expected: %v, isPalindrome: %v", expected, result)
	}
}

func Test_isPalindrome_3(t *testing.T) {
	expected := true
	result := isPalindrome(nil)
	if result != expected {
		t.Errorf("expected: %v, isPalindrome: %v", expected, result)
	}
}

func Test_isPalindrome_4(t *testing.T) {
	l := newListNode([]int{1})
	expected := true
	result := isPalindrome(l)
	if result != expected {
		t.Errorf("expected: %v, isPalindrome: %v", expected, result)
	}
}
