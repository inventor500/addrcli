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
	"fmt"
	"strings"
)

type Zip struct {
	Zip5   string `json:"zip5"`
	Caveat string `json:"recordType"`
}

func (zip Zip) String() string {
	if zip.Caveat == "" {
		return zip.Zip5
	}
	return fmt.Sprintf("%s (%s)", zip.Zip5, zip.Caveat)
}

type ZipQueryResult struct {
	Status string `json:"resultStatus"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zips   []Zip  `json:"zipList"`
}

func (result ZipQueryResult) String() string {
	return fmt.Sprintf("%s, %s: %s", result.City, result.State, result.Zips)
}

type Address struct {
	Company  string `json:"companyName"`
	Address1 string `json:"addressLine1"`
	Address2 string `json:"addressLine2"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip5     string `json:"zip5"`
	Zip4     string `json:"zip4"`
	County   string `json:"countyName"`
}

func (address Address) String() string {
	var result strings.Builder
	if address.Company != "" {
		result.WriteString(fmt.Sprintf("%s, ", address.Company))
	}
	result.WriteString(fmt.Sprintf("%s, ", address.Address1))
	if address.Address2 != "" {
		result.WriteString(fmt.Sprintf("%s, ", address.Address2))
	}
	result.WriteString(fmt.Sprintf("%s, %s %s", address.City, address.State, address.Zip5))
	if address.Zip4 != "" {
		result.WriteString(fmt.Sprintf("-%s", address.Zip4))
	}
	return result.String()
}

// The same as String, but with new lines
func (address Address) StringFormatted() string {
	var result strings.Builder
	if address.Company != "" {
		result.WriteString(fmt.Sprintf("%s\n", address.Company))
	}
	result.WriteString(fmt.Sprintf("%s\n", address.Address1))
	if address.Address2 != "" {
		result.WriteString(fmt.Sprintf("%s\n", address.Address2))
	}
	result.WriteString(fmt.Sprintf("%s, %s %s", address.City, address.State, address.Zip5))
	if address.Zip4 != "" {
		result.WriteString(fmt.Sprintf("-%s\n", address.Zip4))
	}
	return result.String()
}

type AddressQueryResult struct {
	Status    string    `json:"resultStatus"`
	Addresses []Address `json:"addressList"`
}

func (result AddressQueryResult) String() string {
	return fmt.Sprintf("%s: %s", result.Status, result.Addresses)
}
