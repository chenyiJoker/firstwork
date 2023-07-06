package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	for i := 0; i < 100; i++ {
		go printName("张三", ch, 1)

		<-ch
		go printName("李四", ch, 2)

		<-ch
		go printName("王五", ch, 3)

		<-ch
		go printName("赵六", ch, 4)

		<-ch
	}
}

func printName(name string, ch chan int, expectedOrder int) {
	fmt.Println(name)

	ch <- expectedOrder
}
