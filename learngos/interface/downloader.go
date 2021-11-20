package main

import (
	"fmt"
	"learngos/interface/infra"
)

func getRetrieve() retrieve  {
	return infra.Retrieve{}
}

type retrieve interface {
	Get(string2 string) string
}

func main() {

	var retrieve = getRetrieve()
	fmt.Println(retrieve.Get("https://www.imooc.com/"))
}
