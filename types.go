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

func Sprintf(format string, colored bool, args ...interface{}) string {
	var str string
	if colored {
		str = fmt.Sprintf("\033[1m%s\033[0m", format)
	} else {
		str = format
	}
	return fmt.Sprintf(str, args...)
}

type Result interface {
	StringFormatted(bool) string
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
		result.WriteString(address.Company)
	}
	result.WriteString(address.Address1)
	if address.Address2 != "" {
		result.WriteString(address.Address2)
	}
	result.WriteString(fmt.Sprintf("%s, %s %s", address.City, address.State, address.Zip5))
	if address.Zip4 != "" {
		result.WriteString("-" + address.Zip4)
	}
	return result.String()
}

// The same as String, but with new lines
func (address Address) StringWithLines() string {
	var result strings.Builder
	if address.Company != "" {
		result.WriteString(fmt.Sprintln(address.Company))
	}
	result.WriteString(fmt.Sprintln(address.Address1))
	if address.Address2 != "" {
		result.WriteString(fmt.Sprintln(address.Address2))
	}
	result.WriteString(fmt.Sprintf("%s, %s %s", address.City, address.State, address.Zip5))
	if address.Zip4 != "" {
		result.WriteString("-" + address.Zip4)
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

func (result AddressQueryResult) StringFormatted(colored bool) string {
	if result.Status != "SUCCESS" || len(result.Addresses) == 0 {
		return Sprintf("No results", colored)
	}
	var resultString strings.Builder
	resultString.WriteString(Sprintf("Collected Addresses\n", colored))
	resultString.WriteString(result.Addresses[0].StringWithLines())
	for i := 1; i < len(result.Addresses); i++ {
		resultString.WriteString("\n\n" + result.Addresses[i].StringWithLines())
	}
	return resultString.String()
}

// Athens, GR
type CityState struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func (citystate CityState) String() string {
	return fmt.Sprintf("%s, %s", citystate.City, citystate.State)
}

type CityQueryResult struct {
	Status      string      `json:"resultStatus"`
	City        string      `json:"defaultCity"`
	State       string      `json:"defaultState"`
	Zip         string      `json:"zip5"`
	Caveat      string      `json:"defaultRecordType"`
	OtherCities []CityState `json:"citiesList"`
	NonAccept   []CityState `json:"nonAcceptList"`
}

func (result CityQueryResult) String() string {
	return fmt.Sprintf("%s: %s, %s %s (%d alternatives, %d non-accepts)", result.Status, result.City, result.State, result.Zip, len(result.OtherCities), len(result.NonAccept))
}

func (res CityQueryResult) StringFormatted(colored bool) string {
	if res.Status != "SUCCESS" {
		return Sprintf("No results", colored)
	}
	var resultString strings.Builder
	resultString.WriteString(Sprintf("%s\n", colored, res.Zip))
	resultString.WriteString(fmt.Sprintf("%s, %s %s", res.City, res.State, res.Zip))
	if res.Caveat != "" {
		resultString.WriteString(fmt.Sprintf(" (%s)", res.Caveat))
	}
	if len(res.OtherCities) > 0 {
		resultString.WriteString(Sprintf("\n\n%d Other Cities", colored, len(res.OtherCities)))
		for i := 0; i < len(res.OtherCities); i++ {
			resultString.WriteString("\n" + res.OtherCities[i].String())
		}
	}
	if len(res.NonAccept) > 0 {
		resultString.WriteString(Sprintf("\n\n%d Not Accepted Cities", colored, len(res.NonAccept)))
		for i := 0; i < len(res.NonAccept); i++ {
			resultString.WriteString("\n" + res.NonAccept[i].String())
		}
	}
	return resultString.String()
}

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

func (result ZipQueryResult) StringFormatted(colored bool) string {
	if result.Status != "SUCCESS" || len(result.Zips) == 0 {
		return Sprintf("No results", colored)
	}
	var returnValue strings.Builder
	returnValue.WriteString(Sprintf("%s, %s", colored, result.City, result.State))
	for i := 0; i < len(result.Zips); i++ {
		returnValue.WriteString("\n" + result.Zips[i].String())
	}
	return returnValue.String()
}
