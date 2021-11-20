package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
  for 循环
*/
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func swap(a, b int) (int, int) {
	b,a = a,b
	return a, b
}

func main() {
	fmt.Println(
		convertToBin(5), // 101
		convertToBin(13), // 1101
		)

	printFile("abc.txt")

	a, b := 3, 4
	a, b = swap(a, b)

	fmt.Println(a, b)
}
