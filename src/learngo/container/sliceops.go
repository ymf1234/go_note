package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("v=%v len=%d, cap=%d\n", s, len(s), cap(s))
}

/**
  切片操作
 */
func main() {
	/**
	 创建切片
	 */
	fmt.Println("创建切片")
	var s []int // 默认值 nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i + 1)
	}

	fmt.Println(s)

	s1 := []int{2,4,6,8}
	printSlice(s1)

	s2 := make([]int, 16)
	printSlice(s2)

	s3 := make([]int, 10, 32)
	printSlice(s3)


	fmt.Println("拷贝切片")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("删除切片")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	front := s2[0]
	s2 = s2[1:]

	fmt.Println("front = ", front)
	printSlice(s2)

	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]

	fmt.Println("tail = ", tail)
	printSlice(s2)



}
