package main

import (
	"fmt"
	"io/ioutil"
)

/**
  switch
	switch 后可以没有表达式
 */
func grade (score int) string  {
	g := ""
	switch  {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60 :
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("Wrong score: %d", score))
	}
	return g
}

/**
  控制语句
  if 的 条件里可以赋值
  if 的 条件里赋值的变量作用域就在这个if语句里
 */
func main() {
	const filename = "abc.txt"
	/*contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}*/

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(20),
		grade(60),
		grade(80),
		grade(90),
		grade(101),
		)
}
