package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type Winloss struct {
	PlayerID string `json:"playerID"`
	Username string `json:"username"`
	Amount   int    `json:"amount"`
}

func main() {

	jsonStr := `{"lose":[{"playerid":"C56227QR3AU6URAE6940","username":"willis","amount":-4500},{"playerid":"mike","username":"mike","amount":-6500}],"win":[{"playerid":"willis","username":"willis1","amount":3500},{"playerid":"willis2","username":"willis","amount":4500},{"playerid":"mike","username":"mike","amount":7500}]}`

	var m map[string]interface{}
	_ = json.Unmarshal([]byte(jsonStr), &m)

	list, _ := m["win"].([]interface{})
	fmt.Println(list)

	sort.Slice(list, func(i, j int) bool {
		// fmt.Println(list[i].(map[string]interface{})["amount"])
		fi := list[i].(map[string]interface{})["amount"].(float64)
		fj := list[j].(map[string]interface{})["amount"].(float64)
		return fi > fj
	})

	fmt.Println(list)
}
