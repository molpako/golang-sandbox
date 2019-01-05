package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	name := os.Args[1]
	u, err := user.Lookup(name)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)
}
