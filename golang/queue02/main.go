package main

import (
	"errors"
	"fmt"
	"os"
)

// 数组队列,使用环形数组
type Queue struct {
	maxSize int     //队列最大容量
	front   int     //首部
	rear    int     //尾部
	queue   [10]int //存储队列数据
}

//切片实现队列
func main() {
	fmt.Printf("%s: ", "队列启动，请输入队列容量")
	var maxSize int
	fmt.Scanln(&maxSize)
	queue := &Queue{
		maxSize: maxSize + 1,
		front:   0,
		rear:    0,
	}

	var key string
	for {
		fmt.Printf("参数信息 queue: %+v \n", queue)
		fmt.Println("1、添加数据入队列，请输入add")
		fmt.Println("2、获取队列数据，请输入get")
		fmt.Println("3、展示队列，请输入show")
		fmt.Println("4、退出系统，请输入exit")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("请输入需要添加的数字")
			var num int
			fmt.Scanln(&num)
			if ok := queue.addQueue(num); ok != nil {
				fmt.Println(ok)
			}
		case "get":
			if data, ok := queue.getQueue(); ok != nil {
				fmt.Println(ok)
			} else {
				fmt.Println(data)
			}
		case "show":
			queue.showQueue()
		case "exit":
			os.Exit(0)
		}
		fmt.Println()
	}
}

//队列添加数据
func (queue *Queue) addQueue(data int) (error error) {
	if queue.isFull() == true {
		return errors.New("queue full")
	}

	queue.queue[queue.rear] = data
	queue.rear = (queue.rear + 1) % queue.maxSize
	return
}

//队列弹出数据
func (queue *Queue) getQueue() (data int, err error) {
	if queue.isEmpty() == true {
		return -1, errors.New("queue empty")
	}
	data = queue.queue[queue.front]
	queue.front = (queue.front + 1) % queue.maxSize
	return
}

//查看队列所有数据
func (queue *Queue) showQueue() {
	//获取队列数组个数
	size := queue.getSize()
	if size == 0 {
		fmt.Println("队列为空")
	}
	tmpFront := queue.front
	for i := 0; i < size; i++ {
		fmt.Printf("queue[%d]%d\n", tmpFront, queue.queue[tmpFront])
		tmpFront = (tmpFront + 1) % queue.maxSize
	}
	fmt.Println()
}

//判断队列是有已满
func (queue *Queue) isFull() bool {
	return (queue.rear+1)%queue.maxSize == queue.front
}

//判断队列是否为空
func (queue *Queue) isEmpty() bool {
	return queue.front == queue.rear
}

//获取队列中元素
func (queue *Queue) getSize() int {
	return (queue.rear + queue.maxSize - queue.front) % queue.maxSize
}
