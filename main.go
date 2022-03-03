package main

import (
	"os"
)

func main() {
	url_A := os.Args[1]
	url_B := os.Args[2]
	DiffToCSV(url_A, url_B)

}
