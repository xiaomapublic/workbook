package main

import "fmt"

//定义数据结构体
type HeroNode struct {
	number   int       //编号
	name     string    //名称
	nickname string    //昵称
	next     *HeroNode //下一个结点的地址
	pre      *HeroNode //上一个结点的地址
}

//单链表，实现增删改查
func main() {
	//初始化头结点
	head := &HeroNode{}

	//定义需要插入的数据
	data1 := &HeroNode{
		number:   1,
		name:     "唐僧",
		nickname: "师傅",
	}

	data2 := &HeroNode{
		number:   2,
		name:     "孙悟空",
		nickname: "泼猴",
	}

	data22 := &HeroNode{
		number:   2,
		name:     "假悟空",
		nickname: "六耳猕猴",
	}

	data3 := &HeroNode{
		number:   3,
		name:     "猪八戒",
		nickname: "老猪",
	}

	data4 := &HeroNode{
		number:   4,
		name:     "沙僧",
		nickname: "卷帘大将",
	}

	data5 := &HeroNode{
		number:   5,
		name:     "白龙马",
		nickname: "小白龙",
	}

	add(head, data5)
	add(head, data1)
	add(head, data4)
	add(head, data2)
	add(head, data3)
	add(head, data22)
	get(head)

	del(head, 2)
	get(head)

	mod(head, 4, 2)
	get(head)
}

//增：排序新增
//head 头指针
//newNode 需要新增的结点
func add(head *HeroNode, newNode *HeroNode) {
	//创建一个辅助结点[跑龙套, 帮忙]
	tempNode := head
	for {

		if tempNode.next == nil { //说明到链表的最后
			break
		} else if tempNode.next.number >= newNode.number { //说明新节点应该插入到tempNode后面
			break
		}
		//未找到则节点往后移
		tempNode = tempNode.next
	}

	newNode.next = tempNode.next
	newNode.pre = tempNode
	if tempNode.next != nil { //边界问题，如果要删除的数据为最后一个，则不会next为nil
		tempNode.next.pre = newNode
	}
	tempNode.next = newNode
}

//删
func del(head *HeroNode, number int) {
Loop:
	for {
		tempNode := head
		for {
			if tempNode.next == nil {
				//fmt.Printf("%s\n", "要删除的数据不存在")
				break Loop
			} else if tempNode.next.number == number {
				break
			}
			tempNode = tempNode.next
		}

		if tempNode.next.next != nil { //边界问题，如果要删除的数据为最后一个，则next为nil
			tempNode.next.next.pre = tempNode
		}

		tempNode.next = tempNode.next.next
	}

}

//改
func mod(head *HeroNode, number int, modNumber int) {
Loop:
	for {
		tempNode := head
		for {
			if tempNode.next == nil { //如果到链表最后还未找到则结束代码执行
				//fmt.Printf("%s\n", "要修改的数据不存在")
				break Loop
			} else if tempNode.next.number == number {
				break
			}
			tempNode = tempNode.next
		}

		modNode := tempNode.next
		modNode.number = modNumber
		tempNode.next = tempNode.next.next
		if tempNode.next != nil { //边界问题，如果要删除的数据为最后一个，则next为nil
			tempNode.next.pre = tempNode
		}

		add(head, modNode)
	}
}

//查
func get(head *HeroNode) {
	tempNode := head
	if tempNode.next == nil {
		fmt.Printf("%s\n", "链表为空")
		return
	}

	for {
		fmt.Printf("%d %s %s => ", tempNode.next.number, tempNode.next.name, tempNode.next.nickname)
		tempNode = tempNode.next
		if tempNode.next == nil {
			break
		}
	}
	fmt.Println()
}
