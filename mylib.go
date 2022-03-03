package main

import (
	"fmt"
	"os"

	"github.com/anaskhan96/soup"
)

func get_map_metadata(url string) map[string]string {
	data := make(map[string]string)
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
		key := row.FindAll("td")[6].Text()

		e := file + "," + line + "," + rule

		data[key] = e
	}

	return data
}

func Run(url_A string) {

	data_A := get_map_metadata(url_A)
	data_B := get_map_metadata(url_A)
	fmt.Println(data_A)
	fmt.Println(data_B)
}
