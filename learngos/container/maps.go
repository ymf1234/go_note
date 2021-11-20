package main

import "fmt"

func main() {
	m := map[string]string{
		"name": "cc",
		"lang": "golang",
		"site": "site",
	}

	fmt.Println(m)

	m2 := make(map[string] int) // m2 == empty map
	fmt.Println(m2)

	var m3 map[string]int // m3 == nil
	fmt.Println(m3)

	fmt.Println("map 遍历")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("获取值")
	lang, ok := m["lang"]
	fmt.Println(lang, ok)

	if lang, ok := m["lan"]; ok {
		fmt.Println(lang)
	} else {
		fmt.Println("key 不存在")
	}

	fmt.Println("删除元素")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
