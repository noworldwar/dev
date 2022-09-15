package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// cnt := 0
	questions := [4]string{"請輸入總數 1 ~ 10,000,000：", "是否要顯示建立玩家資訊（y/n）？：", "是否要顯示中獎組合結果（y/n）？：", "是否要顯示計算所耗時（y/n）？："}
	var vals []string
	var val string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(questions[0])
	for i := 1; scanner.Scan(); i++ {
		val = scanner.Text()
		switch i {
		case 1:
			amount, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println("格式錯誤")
				i--
			} else if amount < 0 || amount > 10000000 {
				fmt.Println("請輸入總數 1 ~ 10,000,000：")
				i--
			} else {
				vals = append(vals, strings.TrimSpace(scanner.Text()))
			}
			fmt.Println(amount)
			break
		case 2:
		case 3:
		case 4:
			val = strings.TrimSpace(val)
			if val != "y" && val != "n" {
				fmt.Println("請輸入y或n")
				i--
			} else {
				vals = append(vals, strings.TrimSpace(scanner.Text()))
			}
			fmt.Println(val)
			break
		}

		if i == 4 {
			break
		}
		fmt.Println(questions[i])
	}
	fmt.Println(vals)

}
