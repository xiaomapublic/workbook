package main

import (
	"errors"
	"fmt"
)

type list struct {
	number int
	next   *list
	pre    *list
}

//约瑟夫问题
//设编号为1,2,3....n的n个人围坐一圈
//约定编号为k(1 <= k <= n)的人从1开始报数
//数到m的那个人出列，他的下一位又从1开始报数，数到m的那个人又出列，依次类推
//直到所有人出列为止：1.得出出队数列,2.得出最后两个人的编号
func main() {
	//人数
	n := 100
	//起始点
	start := 36
	//间隔点
	offset := 9

	head := &list{}

	for i := 1; i <= n; i++ {
		push(head, i)
	}
	show(head)
	josephusPop(head, start, offset)

	//弹出指定number数据
	//if front , data, ok := pop(head, 0); ok == nil {
	//	fmt.Printf("%d\n", data)
	//	head = front
	//} else {
	//	fmt.Println(ok)
	//}
	//if front , data, ok := pop(head, 1); ok == nil {
	//	fmt.Printf("%d\n", data)
	//	head = front
	//} else {
	//	fmt.Println(ok)
	//}
	//if front , data, ok := pop(head, 100); ok == nil {
	//	fmt.Printf("%d\n", data)
	//	head = front
	//} else {
	//	fmt.Println(ok)
	//}
	//show(head)
}

//约瑟夫弹出数据
func josephusPop(head *list, start int, offset int) {
	temp := head
	for {
		if temp.number == start {
			for {
				//判断链表是否为空
				if temp.next == nil {
					fmt.Println()
					fmt.Println("link empty")
					return
				}

				//展示当前链表数据
				help := temp
				fmt.Printf("%s: ", "now link")
				for {
					fmt.Printf("%d => ", help.number)
					if help.next == temp {
						break
					}
					help = help.next
				}
				fmt.Println()

				//下标偏移
				for i := 1; i < offset; i++ {
					temp = temp.next
				}

				//弹出数据
				fmt.Printf("pop: %d\n", temp.number)
				//删除数据
				if temp.next == temp {
					temp.number = 0
					temp.next = nil
					temp.pre = nil
				} else {
					temp.pre.next = temp.next
					temp.next.pre = temp.pre
					temp = temp.next
				}
			}
		} else if temp.next == head {
			fmt.Println("no find")
			return
		}
		temp = temp.next
	}
}

//弹出数据
func pop(head *list, number int) (*list, int, error) {
	temp := head
	for {
		if temp.next == nil {
			return head, -1, errors.New("link empty")
		} else if temp.number == number {
			data := temp.number
			if temp.next == temp {
				temp.next = nil
				temp.pre = nil
				temp.number = 0
			} else {
				temp.next.pre = temp.pre
				temp.pre.next = temp.next
				if temp == head {
					return temp.next, data, nil
				} else {
					return head, data, nil
				}
			}
		} else if temp.next == head {
			return head, -1, errors.New("no find")
		}
		temp = temp.next
	}
}

//推入数据
func push(head *list, number int) {
	tempNode := head
Loop:
	for {
		switch tempNode.next {
		case nil:
			tempNode.number = number
			tempNode.next = tempNode
			tempNode.pre = tempNode
			break Loop
		case head:
			newNode := &list{number: number, next: head}
			newNode.pre = tempNode
			tempNode.next = newNode
			head.pre = newNode
			break Loop
		}
		tempNode = tempNode.next
	}
}

//展示数据
func show(head *list) {
	temp := head
	if temp.next == nil {
		fmt.Println("link empty")
	}

	for {
		fmt.Printf("%d => ", temp.number)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	fmt.Println()
}
