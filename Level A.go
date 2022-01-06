package main

import "fmt"

var number int = 0

func cycle() {
	// var count int=0
	for count := 0; count < 100000; count++ {
		number = number + 1
		count = count + 1
		fmt.Println(number)
		fmt.Println(count)
	}
}

func counter1() {
	cycle()
	fmt.Println("Counter1:", number)
}
func counter2() {
	cycle()
	fmt.Println("Counter2:", number)
}

func main() {
	counter1()
	counter2()
}
