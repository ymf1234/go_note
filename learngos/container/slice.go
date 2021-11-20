package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}
/**
  切片
 */
func main() {
	arr := [...]int{0,1,2,3,4,5,6,7}

	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	s1 := arr[2:]
	fmt.Println("修改前s1 arr[2:] = ", s1)
	s2 := arr[:]
	fmt.Println("修改前s2 arr[:] = ", s2)
	updateSlice(s1)
	fmt.Println("修改后s1")
	fmt.Println(s1)
	fmt.Println("原始值arr")
	fmt.Println(arr)

	updateSlice(s2)
	fmt.Println("修改后s2")
	fmt.Println(s2)
	fmt.Println("原始值arr")
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)


	fmt.Println("Extending slice")
	arr[0] , arr[2] = 0, 2
	fmt.Println(arr)
	s1 = arr[2:6]
	s2 = s2[3:5]
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1)=%d\n" , s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2)=%d\n" , s2, len(s2), cap(s2))


	// 添加元素
	fmt.Println("添加元素")
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	fmt.Println("s5 = ", s5)
	fmt.Println("arr =", arr)
}
