package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID    string  `json:"id,id1"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	ID         string      `json:"id"`
	Items      []OrderItem `json:"items"`
	Quantity   int         `json:"quantity"`
	TotalPrice float64     `json:"total_price"`
}

func main() {
	marshal()
	parseNLP()
}

// 转为json
func marshal() {
	o := Order{
		ID:         "1234",
		Quantity:   3,
		TotalPrice: 30,
		Items: []OrderItem{
			{
				ID:    "item_1",
				Name:  "learn go",
				Price: 15,
			},
			{
				ID:    "item_2",
				Name:  "learn go",
				Price: 15,
			},
		},
	}

	// 转成json
	b, err := json.Marshal(o)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", b)
}

// 解析json
func unmarshal() {
	// 解析json
	s := `{"id":"1234","items":[{"id":"item_1","name":"learn go","price":15},{"id":"item_2","name":"learn go","price":15}],"quantity":3,"total_price":30}`
	var o1 Order
	err := json.Unmarshal([]byte(s), &o1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", o1)
}

func parseNLP() {
	res := `{
"data": [
    {
        "synonym":"",
        "weight":"0.6",
        "word": "真丝",
        "tag":"材质"
    },
    {
        "synonym":"",
        "weight":"0.8",
        "word": "韩都衣舍",
        "tag":"品牌"
    },
    {
        "synonym":"连身裙;联衣裙",
        "weight":"1.0",
        "word": "连衣裙",
        "tag":"品类"
    }
]
}`

	m := struct {
		Data []struct{
			Synonym string `json:"synonym"`
			Weight string `json:"weight"`
			Word string `json:"word"`
			Tag string `json:"tag"`
		} `json:"data"`
	}{}
	err := json.Unmarshal([]byte(res), &m)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", m.Data[2].Synonym)
}
