package main

import (
	"errors"
	"fmt"
)

//定义数据结构体
type circlesingleNode struct {
	data interface{}       //数据
	next *circlesingleNode //下一个结点的地址
	pre  *circlesingleNode //上一个结点的地址
}

//链表实现环形队列
func main() {
	//初始化头结点
	head := &circlesingleNode{}

	push(head, 5)
	push(head, 1)
	push(head, 4)
	push(head, 2)
	push(head, 3)
	push(head, 22)
	push(head, "s")
	get(head)

	if data, ok := pop(head); ok != nil {
		fmt.Println(ok)
	} else {
		fmt.Println("pop data: ", data)
	}
	get(head)

	if data, ok := pop(head); ok != nil {
		fmt.Println(ok)
	} else {
		fmt.Println("pop data: ", data)
	}
	get(head)

	if data, ok := pop(head); ok != nil {
		fmt.Println(ok)
	} else {
		fmt.Println("pop data: ", data)
	}
	get(head)

	if data, ok := pop(head); ok != nil {
		fmt.Println(ok)
	} else {
		fmt.Println("pop data: ", data)
	}
	get(head)

	if data, ok := pop(head); ok != nil {
		fmt.Println(ok)
	} else {
		fmt.Println("pop data: ", data)
	}
	get(head)

	if data, ok := pop(head); ok != nil {
		fmt.Println(ok)
	} else {
		fmt.Println("pop data: ", data)
	}
	get(head)

	if data, ok := pop(head); ok != nil {
		fmt.Println(ok)
	} else {
		fmt.Println("pop data: ", data)
	}
	get(head)

	if data, ok := pop(head); ok != nil {
		fmt.Println(ok)
	} else {
		fmt.Println("pop data: ", data)
	}
	get(head)
}

//推入数据
func push(head *circlesingleNode, data interface{}) {
	//创建一个辅助结点[跑龙套, 帮忙]
	tempNode := head
	newNode := &circlesingleNode{
		data: data,
		next: nil,
		pre:  nil,
	}
	for {

		if tempNode.next == nil { //说明到链表的最后
			break
		} else if tempNode.next == head { //说明新节点应该插入到tempNode后面
			break
		}
		//未找到则节点往后移
		tempNode = tempNode.next
	}

	newNode.next = head
	newNode.pre = tempNode
	tempNode.next = newNode
	head.pre = newNode
}

//弹出数据
func pop(head *circlesingleNode) (data interface{}, err error) {
	tempNode := head

	if tempNode.next == nil {
		return 0, errors.New("queue empty")
	}

	data = tempNode.next.data
	if tempNode.next.next != head {
		tempNode.next = tempNode.next.next
	} else {
		tempNode.pre = nil
		tempNode.next = nil
	}
	return data, nil
}

//查
func get(head *circlesingleNode) {
	tempNode := head
	if tempNode.next == nil {
		fmt.Printf("%s\n", "link empty")
		return
	}

	for {
		fmt.Printf("%v => ", tempNode.next.data)
		tempNode = tempNode.next
		if tempNode.next == head {
			break
		}
	}
	fmt.Println()
}
