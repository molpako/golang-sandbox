package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func open() io.Reader {
	f, err := os.Open("../main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	return f
}

func main() {
	//r := open()
	f, err := os.Open("../main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	g := open()
	s := bufio.NewScanner(g)

	for s.Scan() {
		fmt.Println(s.Text())
	}
}
