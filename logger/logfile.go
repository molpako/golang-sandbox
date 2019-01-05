package main

import (
	"errors"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("./hh.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Fatal("error opening file :", err.Error())
	}
	log.SetOutput(f)

	callFunc()
}

func callFunc() {
	log.Println(errors.New("oooooooo"))

}
