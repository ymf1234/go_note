package main

import (
	"fmt"
	"learngos/retriever/mock"
	"learngos/retriever/real"
	"time"
)

type Retiever interface {
	Get(url string) string
}

func download(r Retiever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retiever
	r = mock.Retriever{"this is a fake imooc.com"}
	inspect(r)
	mockRetriever := r.(mock.Retriever)
	fmt.Println(mockRetriever.Contents)

	r = & real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	inspect(r)

	// 接口的值类型 之 Type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)




	//fmt.Println(download(r))
}

// 接口的值类型
func inspect(r Retiever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents: ",v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
