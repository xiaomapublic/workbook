package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Queue struct {
	front int
	rear  int
	queue []int
}

//切片实现队列
func main() {
	var queue Queue
	queue.front = -1
	queue.rear = -1
	queue.queue = make([]int, 0)
	randObj := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 10; i++ {
		queue.addQueue(randObj.Intn(100))
	}
	fmt.Println(queue.queue, queue.front, queue.rear)

	for {
		if data, ok := queue.getQueue(); ok {
			fmt.Printf("%d ", data)
		} else {
			fmt.Println("\n队列消耗完毕")
			break
		}
	}
	fmt.Println(queue.queue, queue.front, queue.rear)
}

//队列中添加数据
func (queue *Queue) addQueue(data int) {
	queue.queue = append(queue.queue, data)
	queue.rear++
}

//队列中获取数据
func (queue *Queue) getQueue() (data int, ok bool) {
	if queue.front == queue.rear {
		return 0, false
	} else {
		queue.front++
		return queue.queue[queue.front], true
	}
}
