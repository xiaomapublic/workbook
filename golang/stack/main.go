package main

import (
	"fmt"
)

type Stack struct {
	//栈顶
	top    *Node
	length int
}

type Node struct {
	//上一个数据地址
	pre *Node
	//数据
	data interface{}
}

func main() {
	//定义头结点
	headNode := New()
	headNode.push(1)
	headNode.push(2)
	headNode.push(3)
	headNode.push(4)

	headNode.pop()
	headNode.Len()
	headNode.pop()
	headNode.pop()
	headNode.pop()
	headNode.pop()
	headNode.Len()
	headNode.pop()
	headNode.Len()

}

//创建一个栈
func New() *Stack {
	return &Stack{
		top:    nil,
		length: 0,
	}
}

//入栈
func (s *Stack) push(data interface{}) {
	n := &Node{
		pre:  s.top,
		data: data,
	}
	s.top = n
	s.length++
}

//出栈
func (s *Stack) pop() {
	if s.length == 0 {
		fmt.Println("栈空")
		return
	}
	fmt.Println(s.top.data)
	s.top = s.top.pre
	s.length--
}

//取栈长度
func (s *Stack) Len() {
	fmt.Printf("栈的长度：%d\n", s.length)
}
