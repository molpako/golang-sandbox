package main

import (
	"encoding/csv"
	"io"
	"os"
)

type records [][]string

func main() {
	rs := records{
		{"aaa", "bbb", "ccc"},
		{"aaa", "bbb", "ccc"},
		{"aaa", "bbb", "ccc"},
	}

	rs.change()

}

func (rs records) change() {

	var csvRecords records
	for _, r := range rs {
		r[0] = "changed"
	}
	rs.writeCSV(os.Stdout)
}

// writeCSV writes records to CSV
func (rs records) writeCSV(w io.Writer) error {
	cw := csv.NewWriter(w)
	defer cw.Flush()

	return cw.WriteAll(rs)
}
