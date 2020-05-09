package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	front   int
	rear    int
	queue   []int
}

//切片实现队列
func main() {
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
		queue:   make([]int, 0),
	}

	var key string
	for {
		fmt.Printf("front: %d, rear: %d \n", queue.front, queue.rear)
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

	if queue.rear+1 == queue.maxSize {
		return errors.New("queue full")
	}
	queue.rear++
	queue.queue = append(queue.queue, data)
	return
}

//队列弹出数据
func (queue *Queue) getQueue() (data int, err error) {
	if queue.front == queue.rear {
		return -1, errors.New("queue empty")
	}
	queue.front++
	return queue.queue[queue.front], nil
}

//查看队列所有数据
func (queue *Queue) showQueue() {
	for i := queue.front + 1; i <= queue.rear; i++ {
		fmt.Printf("queue[%d]%d\n", i, queue.queue[i])
	}
	fmt.Println()
}
