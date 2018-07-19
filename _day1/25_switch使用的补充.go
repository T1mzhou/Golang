package main

import "fmt"

func main() {
	switch num := 4; num {
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

	score := 95
	switch {
	case score > 90:
		fmt.Println("优秀")
	case score > 80:
		fmt.Println("良好")
	case score > 70:
		fmt.Println("一般")
	default:
		fmt.Println("其它")
	}
}
