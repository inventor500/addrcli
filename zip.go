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

import "fmt"

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

type QueryResult struct {
	Status string `json:"resultStatus"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zips   []Zip  `json:"zipList"`
}

func (result QueryResult) String() string {
	return fmt.Sprintf("%s, %s: %s", result.City, result.State, result.Zips)
}
