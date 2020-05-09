package main

import (
	"errors"
	"fmt"
	"os"
)

// 简单数组队列,未使用环形数组,性能不好
type Queue struct {
	maxSize int   //队列最大容量
	rear    int   //当前下标
	queue   []int //存储队列数据
}

//切片实现队列
func main() {
	fmt.Printf("%s: ", "队列启动，请输入队列容量")
	var maxSize int
	fmt.Scanln(&maxSize)
	queue := &Queue{
		maxSize: maxSize,
		rear:    -1,
		queue:   make([]int, 0),
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

	if queue.rear+1 == queue.maxSize {
		return errors.New("queue full")
	}
	queue.rear++
	queue.queue = append(queue.queue, data)
	return
}

//队列弹出数据
func (queue *Queue) getQueue() (data int, err error) {
	if queue.rear == -1 {
		return -1, errors.New("queue empty")
	}
	data = queue.queue[0]
	//将切片下标为0的数据删除,大批量数据会存在性能问题
	queue.queue = append(queue.queue[:0], queue.queue[1:]...)
	//下标回退
	queue.rear--
	return data, nil
}

//查看队列所有数据
func (queue *Queue) showQueue() {
	for i := 0; i <= queue.rear; i++ {
		fmt.Printf("queue[%d]%d\n", i, queue.queue[i])
	}
	fmt.Println()
}
