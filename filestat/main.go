package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	name := os.Args[1]
	f, _ := os.Open(name)
	s, _ := f.Stat()
	stat := s.Sys().(*syscall.Stat_t)
	fmt.Println("uid:", stat.Uid)
	fmt.Println("gid:", stat.Gid)
}
