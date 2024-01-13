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
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const BaseURL = "https://tools.usps.com/tools/app/ziplookup/"

func addHeaders(req *http.Request, referer string) {
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; rv:109.0) Gecko/20100101 Firefox/115.0")
	req.Header.Set("Accept", "application/json, text/javascript")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("DNT", "1")
	req.Header.Set("Sec-GPC", "1")
	req.Header.Set("Origin", "https://tools.usps.com")
	req.Header.Set("Referer", referer)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
}

// TODO: Address query

func zipQuery(city string, state string) (*ZipQueryResult, error) {
	data := url.Values{
		"city":  {city},
		"state": {state},
	}
	client := http.Client{}
	req, _ := http.NewRequest("POST", BaseURL+"zipByCityState", strings.NewReader(data.Encode()))
	addHeaders(req, "https://tools.usps.com/zip-code-lookup.htm?bycitystate")
	res, err := client.Do(req)
	if err != nil {
		return &ZipQueryResult{}, errors.New(fmt.Sprintf("Error sending request: %s", err))
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return &ZipQueryResult{}, errors.New(fmt.Sprintf("Invalid status code: %d", res.StatusCode))
	}
	var result ZipQueryResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return &ZipQueryResult{}, errors.New(fmt.Sprintf("Error decoding result: %s", err))
	}
	return &result, nil
}

func addressQuery(company string, address1 string, address2 string, city string, state string, zip string) (*AddressQueryResult, error) {
	data := url.Values{
		"companyName": {company},
		"address1":    {address1},
		"address2":    {address2},
		"city":        {city},
		"state":       {state},
		"zip":         {zip},
		"urbanCode":   {""}, // TODO: What does this actually do?
	}
	client := http.Client{}
	req, _ := http.NewRequest("POST", BaseURL+"zipByAddress", strings.NewReader(data.Encode()))
	addHeaders(req, "https://tools.usps.com/zip-code-lookup.htm?byaddress")
	res, err := client.Do(req)
	var result AddressQueryResult
	if err != nil {
		return &result, errors.New(fmt.Sprintf("Error sending request: %s", err))
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return &result, errors.New(fmt.Sprintf("Invalid status code: %d", res.StatusCode))
	}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return &AddressQueryResult{}, errors.New(fmt.Sprintf("Error decoding result: %s", err))
	}
	return &result, nil
}
