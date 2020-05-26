package main

import (
	"fmt"
	"math/rand"
	"time"
)

//插入排序
func main() {
	randObj := rand.New(rand.NewSource(time.Now().UnixNano()))
	listAsc := randObj.Perm(1000)
	listDesc := randObj.Perm(1000)

	time1 := time.Now()
	fmt.Printf("%v\n", listAsc)
	asc(listAsc)
	fmt.Printf("正序耗时：%s\n", time.Since(time1))
	fmt.Printf("正序结果：%v\n", listAsc)

	time2 := time.Now()
	fmt.Printf("%v\n", listDesc)
	desc(listDesc)
	fmt.Printf("逆序耗时：%s\n", time.Since(time2))
	fmt.Printf("逆序结果：%v\n", listDesc)

}

//正序
func asc(list []int) {
	for i := range list {
		preIndex := i - 1
		current := list[i]
		for preIndex >= 0 && list[preIndex] > current {
			list[preIndex+1] = list[preIndex]
			preIndex--
		}
		list[preIndex+1] = current
	}
}

//逆序
func desc(list []int) {
	for i := range list {
		preIndex := i - 1
		current := list[i]
		for preIndex >= 0 && list[preIndex] < current {
			list[preIndex+1] = list[preIndex]
			preIndex--
		}
		list[preIndex+1] = current
	}
}
