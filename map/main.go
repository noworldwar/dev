package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// var m1 = map[string]float64{
	// 	"Bell Labs": 40.15615,
	// 	"Google":    51.15610,
	// }

	// for k1, v1 := range m1 {
	// 	fmt.Println(k1, ": ", v1)
	// }

	var m2 map[string]interface{}

	fmt.Println(m2)
	// map[]

	m3 := make(map[string]interface{})
	m3["name"] = "mike"

	m3["zero"] = 0

	fmt.Println(m3)
	// map[name:mike zero:0]

	data, err := json.Marshal(m3)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	// {"name":"mike","zero":0}

	m4 := make(map[string]interface{})

	m4["name"] = "Patrick"
	m4["age"] = 18
	m4["weight"] = 80.984

	err = json.Unmarshal(data, &m4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m4)
	// map[age:18 name:mike weight:80.984 zero:0]

}
