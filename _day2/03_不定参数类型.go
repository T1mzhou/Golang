package main //必须

import "fmt"

func MyFunc01(a int, b int) {

}

func MyFunc02(args ...int) {
	fmt.Println("len(args) = ", len(args))
	for i := 0; i < len(args); i++ {
		fmt.Println("args[%d] = %d\n", i, args[i])
	}

	fmt.Println("===============================")

	for i, data := range args {
		fmt.Printf("args[%d] = %d\n", i, data)
	}
}

func MyFunc03(a int, args ...int) {

}

func main() {
	MyFunc01(666, 777)
	MyFunc02()
	fmt.Println("+++++++++++++++++++++++")
	MyFunc02()
	fmt.Println("+++++++++++++++++++++++")
	MyFunc02(1, 2, 3)
	MyFunc03(111, 1, 2, 3)
}
