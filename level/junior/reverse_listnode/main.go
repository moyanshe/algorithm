package main

import (
	"algorithm/helpers"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var values []int
	var input string

	fmt.Println("请输入一个数组用作链表的值，格式如：1,2,3,4")
	fmt.Scan(&input)

	var inputList = strings.Split(input, ",")
	for _,v := range inputList {
		if value,err := strconv.Atoi(v);err == nil {
			values = append(values, value);
		}
	}
	
	var listNode = helpers.GenListNode(values)
	var newListNode = ReverseListNode(listNode)
	
	for newListNode != nil {
		fmt.Println(newListNode.Value)
		newListNode = newListNode.Next;
	}

}

// ReverseListNode 反转链表
func ReverseListNode(listNode *helpers.ListNode) *helpers.ListNode{
	var prev, curr, next *helpers.ListNode
	curr = listNode

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	
	return prev;
}