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

	rows := doc.Find("div", "class", "detail").Find("table").Find("tbody").FindAll("tr")

	pos_rule := 0
	pos_line := 3
	pos_key := 6

	for _, row := range rows {
		file := row.Find("td", "class", "component").Text()
		rule := row.FindAll("td")[pos_rule].Find("a").Text()
		line := row.FindAll("td")[pos_line].Text()
		key := row.FindAll("td")[pos_key].Text()

		e := file + "," + line + "," + rule

		data[key] = e
	}

	return data
}

func diff(data_A map[string]string, data_B map[string]string) map[string]string {
	additions := make(map[string]string)
	for k, v := range data_B {
		_, key_exists := data_A[k]

		if !key_exists {
			additions[k] = v
		}
	}
	return additions
}

func Diff(url_A string, url_B string) map[string]string {
	data_A := get_map_metadata(url_A)
	data_B := get_map_metadata(url_B)

	additions := diff(data_A, data_B)

	return additions
}

func DiffToCSV(url_A string, url_B string) {
	additions := Diff(url_A, url_B)

	for k, v := range additions {
		fmt.Print(k)
		fmt.Print(",")
		fmt.Println(v)
	}
}
