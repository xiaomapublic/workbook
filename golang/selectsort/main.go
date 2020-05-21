package main

import (
	"fmt"
	"math/rand"
	"time"
)

//选择排序
func main() {
	var listAsc, listDesc []int
	randObj := rand.New(rand.NewSource(time.Now().UnixNano()))
	listAsc = randObj.Perm(10)
	listDesc = randObj.Perm(10)
	asc(listAsc)
	desc(listDesc)
	fmt.Printf("%v\n%v\n", listAsc, listDesc)

}

//正序
func asc(list []int) {
	for i := 0; i < len(list)-1; i++ {
		tmp := list[i]
		index := i
		for j := i + 1; j < len(list); j++ {
			if tmp > list[j] {
				tmp = list[j]
				index = j
			}
		}
		list[i], list[index] = list[index], list[i]
	}
}

//逆序
func desc(list []int) {
	for i := 0; i < len(list); i++ {
		tmp := list[i]
		k := i
		for j := i + 1; j < len(list); j++ {
			if tmp < list[j] {
				tmp = list[j]
				k = j
			}
		}
		list[i], list[k] = list[k], list[i]
	}
}
