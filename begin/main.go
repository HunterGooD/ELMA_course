package main

import "fmt"

func fooBar(current, end int, ch chan struct{}) {
	if current%3 == 0 {
		fmt.Printf("%d Bar\n", current)
	} else {
		fmt.Printf("%d Foo\n", current)
	}
	if current == end {
		ch <- struct{}{}
		return
	}
	go fooBar(current+1, end, ch)
}

func main() {
	var n int
	fmt.Print("Введите число N(конечное число итерации): ")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("N меньше или равен нулю.")
		return
	}
	ch := make(chan struct{})
	fooBar(1, n, ch)
	<-ch
}
