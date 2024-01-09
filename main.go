// Copyright (C) 2024  Isaac Ganoung

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func submitQuery(city string, state string) QueryResult {
	data := url.Values{
		"city":  {city},
		"state": {state},
	}
	client := http.Client{}
	req, _ := http.NewRequest("POST", "https://tools.usps.com/tools/app/ziplookup/zipByCityState", strings.NewReader(data.Encode()))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; rv:109.0) Gecko/20100101 Firefox/115.0 ")
	req.Header.Set("Accept", "application/json, text/javascript")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("DNT", "1")
	req.Header.Set("Sec-GPC", "1")
	req.Header.Set("Origin", "https://tools.usps.com")
	req.Header.Set("Referer", "https://tools.usps.com/zip-code-lookup.htm?bycitystate")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln("Could not get url!")
	}
	if res.StatusCode != 200 {
		log.Fatalf("Invalid status code %d", res.StatusCode)
	}
	var result QueryResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error decoding result: %s", err)
	}
	defer res.Body.Close()
	return result
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "\033[31;1mUsage: %s city state\033[0m\n", os.Args[0])
	} else {
		result := submitQuery(os.Args[1], os.Args[2])
		fmt.Printf("\033[1m%s, %s\033[0m\n", result.City, result.State)
		for i := 0; i < len(result.Zips); i++ {
			fmt.Println(result.Zips[i])
		}
	}
}
