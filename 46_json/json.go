package main

import (
	"encoding/json"
	"fmt"
)

type response struct {
	Code    int
	Message string
}

type animal struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

var log = fmt.Println

func main() {
	data1, _ := json.Marshal(true) // 直接出来的数据是以比特流形式的，全是数字
	log(data1)

	data2, _ := json.Marshal(1)
	log(data2)

	data3, _ := json.Marshal(3.123)
	log(string(data3)) // 使用string将比特流转化为可读的字符串

	data4, _ := json.Marshal([]string{"apple", "banana"})
	log(string(data4))

	data5, _ := json.Marshal(map[string]int{"age": 5})
	log(string(data5))

	res := &response{
		Code:    1,
		Message: "Hello",
	}

	data6, _ := json.Marshal(res)
	log(string(data6))

	anim := &animal{
		Kind: "dog",
		Name: "Wangcai",
	}
	data7, _ := json.Marshal(anim)
	log(string(data7))

	byt := []byte(`{"num":1.23,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	log(dat)
	num := dat["num"].(float64)
	log(num)

}
