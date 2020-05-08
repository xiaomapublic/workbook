package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type ValNode struct {
	row int
	col int
	val int
}

//稀疏数组演示
func main() {
	// 创建原始切片数组
	var chessMap [13][13]int
	chessMap[1][1] = 1
	chessMap[2][2] = 2
	chessMap[3][3] = 3
	chessMap[4][4] = 4
	chessMap[5][5] = 5
	chessMap[10][10] = 10

	//打印原始数组
	for _, v1 := range chessMap {
		for _, v2 := range v1 {
			fmt.Printf("%d ", v2)
		}
		fmt.Println()
	}
	fmt.Println()

	sparseaMap := make([]ValNode, 0)
	// 根据切片数组创建稀疏数组
	for i, v1 := range chessMap {
		for j, v2 := range v1 {
			if v2 != 0 {
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseaMap = append(sparseaMap, valNode)
			}
		}
	}

	file, _ := os.OpenFile("sparsearray/sparsearray.data", os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()

	// 方式一：带缓存的写入
	writer := bufio.NewWriter(file)
	//稀疏数组存盘
	for _, v := range sparseaMap {
		// 方式二：直接写入
		//file.WriteString(fmt.Sprintf("%d %d %d\n", v.row, v.col, v.val))

		//writer.Write([]byte(fmt.Sprintf("%d %d %d\n", v.row, v.col, v.val)))
		writer.WriteString(fmt.Sprintf("%d %d %d\n", v.row, v.col, v.val))
	}
	writer.Flush()

	//从文件中读取稀疏数组，还原后打印
	fileOpen, _ := os.Open("sparsearray/sparsearray.data")
	defer fileOpen.Close()
	reader := bufio.NewReader(fileOpen)

	var chessNewMap [13][13]int

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		arr := strings.Split(str, " ")
		row, _ := strconv.Atoi(arr[0])
		col, _ := strconv.Atoi(arr[1])
		val, _ := strconv.Atoi(arr[2])
		chessNewMap[row][col] = val
	}

	//打印还原后的数组
	for _, v1 := range chessMap {
		for _, v2 := range v1 {
			fmt.Printf("%d ", v2)
		}
		fmt.Println()
	}

}
