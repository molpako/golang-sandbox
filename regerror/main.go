package main

import (
	"fmt"
	"log"
	"regexp"
	"sync"
)

func main() {

	var patterns = []string{
		"asdfsaf",
		"[1-9].*",
		"[a-z].*",
		"¥sgggg¥s",
		"relo\nrelo",
		"asdfsfdfffsaf",
		"[]{}",
	}
	regexps, err := Compile(patterns)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(patterns)
	fmt.Println(regexps)

}

func Compile(patterns []string) ([]*regexp.Regexp, []error) {
	var wg sync.WaitGroup
	rc := make(chan *regexp.Regexp)
	errc := make(chan error)
	done := make(chan struct{})

	for _, p := range patterns {
		wg.Add(1)
		go compile(&wg, rc, errc, p)
	}

	go func() {
		wg.Wait()
		done <- struct{}{}
	}()

	var regexps []*regexp.Regexp
	var errors []error
	for {
		select {
		case r := <-rc:
			regexps = append(regexps, r)
		case err := <-errc:
			errors = append(errors, err)
		case <-done:
			return regexps, errors
		}
	}

}

func compile(wg *sync.WaitGroup, rc chan *regexp.Regexp, errc chan error, p string) {
	defer wg.Done()

	regexp, err := regexp.Compile(p)
	if err != nil {
		errc <- err
		return
	}

	rc <- regexp
}
