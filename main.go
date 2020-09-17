package main

import "fmt"

func main() {
	var n int
	var str string
	fmt.Scan(&n)
	for ; n > 0; n-- {
		fmt.Scan(&str)
		if len(str) <= 2 {
			fmt.Println(str)
		}

		flag := 0
		tmp := string(str[0])
		for i := 1; i < len(str); i++ {
			switch flag {
			case 0:
				tmp = tmp + string(str[i])
				if str[i] == str[i-1] {
					flag = 1
				}
			case 1:
				if str[i] != str[i-1] {
					tmp = tmp + string(str[i])
					flag = 2
				}
			case 2:
				if str[i] != str[i-1] {
					tmp = tmp + string(str[i])
					flag = 0
				}
			}
		}
		fmt.Println(tmp)
	}
}
