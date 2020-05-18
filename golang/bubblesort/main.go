package main

import (
	"fmt"
	"math/rand"
	"time"
)

//冒泡排序，优化算法
func main() {
	randObj := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := 36000
	//获取拥有随机数的切片
	var listAscOptimization []int
	var listAscOrdinary []int
	var listDescOptimization []int
	var listDescOrdinary []int
	listAscOptimization = randObj.Perm(n)
	listAscOrdinary = make([]int, n)
	listDescOptimization = make([]int, n)
	listDescOrdinary = make([]int, n)
	//copy 使用时，list2会复制的数量是根据list2初始大小规定的，如果list2初始为空，则不会从list1中复制一个元素进去
	//目标切片(list2)必须分配过空间且足够承载复制的元素个数，并且来源和目标的类型必须一致
	copy(listAscOrdinary, listAscOptimization)
	copy(listDescOptimization, listAscOptimization)
	copy(listDescOrdinary, listAscOptimization)

	//正序优化算法
	t1 := time.Now()
	sortAscOptimization(listAscOptimization)
	//计算耗时
	elapsed1 := time.Since(t1)
	fmt.Printf("正序优化算法耗时%v\n", elapsed1)

	//正序普通算法
	t2 := time.Now()
	sortAscOrdinary(listAscOrdinary)
	//计算耗时
	elapsed2 := time.Since(t2)
	fmt.Printf("正序普通算法耗时%v\n", elapsed2)

	//倒序优化算法
	t3 := time.Now()
	sortDescOptimization(listDescOptimization)
	//计算耗时
	elapsed3 := time.Since(t3)
	fmt.Printf("倒序优化算法耗时%v\n", elapsed3)

	//倒序普通算法
	t4 := time.Now()
	sortDescOrdinary(listDescOrdinary)
	//计算耗时
	elapsed4 := time.Since(t4)
	fmt.Printf("倒序普通算法耗时%v\n", elapsed4)

}

//正序-优化版
func sortAscOptimization(list []int) {
	//标志位，如果循环遍历没有变更数据，表示排序已完成
	flag := true
	//记录最后一次交换数据的位置，下次循环只用比较到上一次循环之前的位置
	k := len(list) - 1
	pos := 0
	for flag {
		flag = false
		for j := 0; j < k; j++ {
			if list[j] > list[j+1] {
				list[j+1], list[j] = list[j], list[j+1]
				flag = true
				pos = j
			}
		}
		k = pos
	}
}

//正序-普通版
func sortAscOrdinary(list []int) {
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-1; j++ {
			if list[j] > list[j+1] {
				list[j+1], list[j] = list[j], list[j+1]
			}
		}
	}
}

//倒序-优化版
func sortDescOptimization(list []int) {
	k := len(list) - 1
	pos := 0
	flag := true
	for flag {
		flag = false
		for j := 0; j < k; j++ {
			if list[j] < list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				flag = true
				pos = j
			}
		}
		k = pos
	}
}

//倒序-普通版
func sortDescOrdinary(list []int) {
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-1; j++ {
			if list[j] < list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}
