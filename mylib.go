package main

import (
	"os"

	"github.com/anaskhan96/soup"
)

// Add ...
func Add(a int, b int) int {
	return a + b
}

func run(url string) {

	resp, err := soup.Get(url)
	if err != nil {
		os.Exit(1)
	}
	soup.HTMLParse(resp)
	// links := doc.Find("div", "id", "comicLinks").FindAll("a")
	// for _, link := range links {
	// 	fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
	// }
}
