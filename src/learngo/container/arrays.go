package main

import "fmt"

func printArray(arr [5]int) {
	fmt.Println("函数 printArray()")
	for k, v := range arr{
		fmt.Println(k, v)
	}
}

func main() {
	/**
	  数组
	 */
	var arr1 [5]int
	arr2 := [3]int{1,3,5}
	arr3 := [...]int{2,4,6,8,10}

	var grid [4][5]int // 4行5列

	fmt.Println(arr1, arr2, arr3, grid)

	// 遍历数组
	fmt.Println("第一种")
	for i:= 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	fmt.Println("第二种")
	for i:= range arr3{
		fmt.Println(arr3[i])
	}

	fmt.Println("第三种")
	for k, v := range arr3{
		fmt.Println(k, v)
	}


	printArray(arr1)
}
