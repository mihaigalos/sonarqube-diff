package main

import (
	"fmt"
	"os"

	"github.com/anaskhan96/soup"
)

func Run(url string) {

	resp, err := soup.Get(url)
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	fmt.Println(doc.Find("title").Text())
	rows := doc.Find("div", "class", "detail").Find("table").Find("tbody").FindAll("tr")

	for _, row := range rows {
		file := row.Find("td", "class", "component").Text()
		line := row.FindAll("td")[3].Text()
		rule := row.FindAll("td")[0].Find("a").Text()

		fmt.Print(file)
		fmt.Print(",")
		fmt.Print(line)
		fmt.Print(",")
		fmt.Print(rule)

		fmt.Println()
	}
}
