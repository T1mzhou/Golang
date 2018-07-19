package main

import "fmt"

func main() {
	var num int
	fmt.Printf("请按下楼层:")
	fmt.Scan(&num)

	switch num {
	case 1:
		fmt.Println("按下的是1楼")
		//break
		fallthrough
	case 2:
		fmt.Println("按下的是2楼")
		break
	case 3:
		fmt.Println("按下的是3楼")
		break
	case 4:
		fmt.Println("按下的是4楼")
		break
	case 5:
		fmt.Println("按下的是5楼")
		break
	default:
		fmt.Println("按下的是xxx楼")
	}
}
