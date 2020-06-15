package main

import (
	"fmt"
	"time"
)

func main() {
	//sortArr := rand.New(rand.NewSource(time.Now().UnixNano())).Perm(100000)
	sortArr := []int{1, 5, 6, 1, 25, -8, 9, 7, 5, 6, 99, 8, 4, 1, 5, 0, 23, 15, 16, 6}
	left := 0
	right := len(sortArr) - 1
	time1 := time.Now()
	asc(left, right, sortArr)
	fmt.Printf("正序耗时：%v\n", time.Since(time1))
	fmt.Printf("正序结果：%v\n", sortArr)

	time2 := time.Now()
	desc(left, right, sortArr)
	fmt.Printf("逆序耗时：%v\n", time.Since(time2))
	fmt.Printf("逆序结果：%v\n", sortArr)
}

//正序
func asc(left int, right int, sortArr []int) {
	//记录起始下标
	l := left
	//记录结束下标
	r := right
	//获取中间数（基准数）的下标
	middleIndex := (left + right) / 2
	//记录中间数的值
	middle := sortArr[middleIndex]

	//将数组分成两组
	//循环比较基准数左右两边的数
	for l < r {
		//如果基准数左边某数小于基准数，则比较下一个，反之则与右边交换位置
		for sortArr[l] < middle {
			l++
		}
		//如果基准数右边某数大于基准数，则比较上一个，反之则与左边数交换位置
		for sortArr[r] > middle {
			r--
		}

		//如果下标相等则表示比较完毕
		if l == r {
			l++
			r--
			break
		}
		//交换基准数左右两边数的位置
		sortArr[l], sortArr[r] = sortArr[r], sortArr[l]

		//交换后，如果左边数等于基准数，则右边下标向左移动一位，如果不移动，会陷入死循环，交换后等于基准数的值，会通过交换趋近基准数的位置
		if sortArr[l] == middle {
			r--
		}
		//交换后，如果右边数等于基准数，则左边下标向右移动一位，如果不移动，会陷入死循环，交换后等于基准数的值，会通过交换趋近基准数的位置
		if sortArr[r] == middle {
			l++
		}
	}

	//只有交换过位置后才继续递归，如果已经排好序，则不进行递归操作
	if l > left {
		asc(left, r, sortArr)
	}
	//只有交换过位置后才继续递归，如果已经排好序，则不进行递归操作
	if r < right {
		asc(l, right, sortArr)
	}
}

func desc(left int, right int, sortArr []int) {
	l := left
	r := right
	middleIndex := (left + right) / 2
	middle := sortArr[middleIndex]

	for l < r {
		for sortArr[l] > middle {
			l++
		}
		for sortArr[r] < middle {
			r--
		}

		if l == r {
			l++
			r--
			break
		}

		sortArr[l], sortArr[r] = sortArr[r], sortArr[l]

		if sortArr[l] == middle {
			r--
		}

		if sortArr[r] == middle {
			l++
		}
	}

	if l > left {
		desc(left, r, sortArr)
	}
	if r < right {
		desc(l, right, sortArr)
	}
}
