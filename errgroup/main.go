package main

import (
	"fmt"
	"strconv"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	strslice := []string{
		"12",
		"323",
		"3423",
		"fa",
		"fdsaf",
	}

	eg := errgroup.Group{}
	c := make(chan int)
	for _, s := range strslice {
		s := s
		eg.Go(func() error {
			return conv(c, s)
		})
	}

	go func() {
		fmt.Println(eg.Wait())

		close(c)
	}()

	var intslice []int
	for i := range c {
		intslice = append(intslice, i)
	}

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(intslice)

}

func conv(c chan<- int, s string) error {
	time.Sleep(time.Millisecond)
	i, err := strconv.Atoi(s)
	c <- i
	return err
}
