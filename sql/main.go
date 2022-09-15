package main

import "fmt"

func main() {
	tableIds := []string{"table1", "table2"}
	sql := `UPDATE live.table_setting SET enabled = ? WHERE table_id IN (`
	for i, v := range tableIds {
		sql += `'` + v + `'`
		if len(tableIds) == 1 {
			break
		}
		if i < len(tableIds)-1 {
			sql += `,`
		}
	}
	sql += `);`
	fmt.Println(sql)
}
