package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Data struct {
	HostID      string `json:"hostID"`
	TableID     string `json:"tableID"`
	Name        string `json:"name"`
	GameType    int    `json:"gameType"`
	Description string `json:"description"`
	TableLimit  string `json:"tableLimit"`
	ShoeID      string `json:"shoeID"`
	Phone       string `json:"phone"`
	Video       Video  `json:"video"`
	Meta        string `json:"meta"`
}

type GameTable struct {
	Data []Data `json:"data"`
}

type Video struct {
	Disabled  string `json:"disabled"`
	Wide_1080 string `json:"wide_1080"`
	Wide_360  string `json:"wide_360"`
	Wide_720  string `json:"wide_720"`
	Zoom_1080 string `json:"zoom_1080"`
}

func main() {

	file, _ := ioutil.ReadFile("gametable_prod.json")

	var gt GameTable
	_ = json.Unmarshal([]byte(file), &gt)
	sql := ""

	for i := 0; i < len(gt.Data); i++ {

		if i == 0 || (i >= 10 && i%10 == 0) {
			sql += "INSERT INTO live.table_setting (table_id,tb_name,tb_type_id,enabled) VALUES \r\n"
		}
		// sql += "INSERT INTO live.table_setting (table_id,tb_name,tb_type_id,enabled) VALUES \r\n"

		if i%10 != 0 {
			sql += ","
		}

		sql += "('" + gt.Data[i].TableID + "','" + gt.Data[i].Name + "'," + strconv.Itoa(gt.Data[i].GameType) + ",1) \r\n"

		if i != 0 && i%10 == 9 {
			sql += ";\r\n"
		}
	}
	fmt.Println(sql)
	fmt.Println("\r\ncount: ", len(gt.Data))
}
