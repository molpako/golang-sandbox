package main

import "fmt"

func main() {
	numbers := []int{53541233, 21235343, 11421443, 5423123}
	c := make(chan []int)
	go factorize(numbers, c)

	for i := range c {
		fmt.Printf("receive: %d\n", i)
	}
}

// 素因数分解する関数
func factorize(numbers []int, c chan<- []int) {
	for _, number := range numbers {
		var a []int
		for i := 1; i < number+1; i++ {
			if number%i == 0 {
				a = append(a, i)
			}
		}
		c <- a
	}
	// 送信側がチャネルをクローズする
	close(c)
}
