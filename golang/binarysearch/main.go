package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//二分查找
func main() {
	randObj := rand.New(rand.NewSource(time.Now().UnixNano()))
	randNumMax := 5000
	arr_one := randObj.Perm(randNumMax)
	arr_two := randObj.Perm(randNumMax)

	arrOneLastIndex := len(arr_one) - 1
	arrTwoLastIndex := len(arr_two) - 1

	//randNum := randObj.Intn(randNumMax)
	randNum := 93

	time1 := time.Now()
	asc(0, arrOneLastIndex, arr_one)
	fmt.Printf("排序耗时：%v\n", time.Since(time1))
	asc(0, arrTwoLastIndex, arr_two)

	time2 := time.Now()
	cycleResult, cycleCount, ok := cycle(0, arrOneLastIndex, arr_one, randNum)
	fmt.Printf("循环方式-信息：%s, 循环次数：%d, 需要查找的数：%d, 查找到的下标：%d, 耗时：%v\n", ok, cycleCount, randNum, cycleResult, time.Since(time2))

	time3 := time.Now()
	recursiveResult, ok := recursive(0, arrTwoLastIndex, arr_two, randNum)
	fmt.Printf("递归方式-信息：%s, 需要查找的数：%d, 查找到的下标：%d, 耗时：%v\n", ok, randNum, recursiveResult, time.Since(time3))

}

//循环二分查找
func cycle(left int, right int, searchArr []int, randNum int) (result int, cycleCount int, err error) {
	l := left
	r := right
	for l <= r {
		cycleCount++
		//中间值得计算有两种方式，
		//方式一：int mid = (low +high)>>1，
		//方式二：int mid = low + ((high - low) >> 1)。
		//方式一存在溢出的风险，当low和high比较大时，有可能会导致mid的值错误，从而使程序出错。
		//方式二则可以保证生成的mid一定大于low，小于high。
		middleIndex := l + ((r - l) >> 1)
		middleValue := searchArr[middleIndex]
		if middleValue > randNum {
			r = middleIndex - 1
		} else if middleValue < randNum {
			l = middleIndex + 1
		} else {
			return middleIndex, cycleCount, errors.New("已查询到")
		}
	}

	return -1, cycleCount, errors.New("未查询到")
}

//递归二分查找
func recursive(left int, right int, searchArr []int, randNum int) (result int, err error) {
	//如果左下标大于右下标，结束递归
	if left > right {
		return -1, errors.New("未查询到")
	}

	//求出中值下标
	middleIndex := left + ((right - left) >> 1)
	middleValue := searchArr[middleIndex]

	if middleValue > randNum {
		//如果中值大于要查找的值，则在中值左边的数据块内递归查找
		result, err = recursive(left, middleIndex-1, searchArr, randNum)
	} else if middleValue < randNum {
		//如果中值小于要查找的值，则在中值右边的数据块内递归查找
		result, err = recursive(middleIndex+1, right, searchArr, randNum)
	} else {
		//如果等于则结束递归
		return middleIndex, errors.New("已查询到")
	}
	return
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
