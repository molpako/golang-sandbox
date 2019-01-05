package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
)

func main() {
	// 最後に改行なし
	patterns01 := []byte("111\n222\n333")
	// 最後に改行あり
	patterns02 := []byte("444\n555\n666\n")

	print(patterns01)
	print(patterns02)
}

func print(patterns []byte) {
	r := bufio.NewReader(bytes.NewReader(patterns))
	for {
		line, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		// 末尾の改行コードを削除して出力する
		fmt.Println(string(bytes.TrimRight(line, "\n")))
	}
}
