package main

import "fmt"
import "time"

func main() {
	i := 0
	for {
		i++
		time.Sleep(time.Second)
		if i == 5 {
			continue
		}
		fmt.Println("i = ", i)
	}
}
