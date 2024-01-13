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
	"flag"
	"fmt"
	"os"
)

type Arguments struct {
	Subcommand string            // TODO: Change to enum?
	Args       map[string]string // TODO: Maybe not a dictionary?
}

// Parse command-line arguments
func get_args() *Arguments {
	address := flag.NewFlagSet("address", flag.ExitOnError)
	zip := flag.NewFlagSet("zip", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Copyright (C) 2024 Isaac Ganoung.")
		fmt.Printf("Usage: %s address|zip [arg0...argn]\n", os.Args[0])
		os.Exit(1)
	}

	switch os.Args[1] {
	case "address":
		company := address.String("company", "", "Company")
		address1 := address.String("address1", "", "Address line 1")
		address2 := address.String("address2", "", "Address line 2")
		city := address.String("city", "", "City")
		state := address.String("state", "", "State")
		zip := address.String("zip", "", "Zip")
		address.Parse(os.Args[2:])
		address_map := make(map[string]string)
		address_map["company"] = *company
		address_map["address1"] = *address1
		address_map["address2"] = *address2
		address_map["city"] = *city
		address_map["state"] = *state
		address_map["zip"] = *zip
		return &Arguments{"address", address_map}
		// return &Arguments{"address", }
	case "zip":
		city := zip.String("city", "", "City")
		state := zip.String("state", "", "State")
		zip.Parse(os.Args[2:])
		zip_map := make(map[string]string)
		zip_map["city"] = *city
		zip_map["state"] = *state
		return &Arguments{"zip", zip_map}
	default:
		fmt.Printf("Unsupported search type %s\n", os.Args[1])
		os.Exit(1)
	}
	return &Arguments{} // TOOD: Why does the compiler require this?
}

func main() {
	arguments := get_args()
	switch arguments.Subcommand {
	case "zip":
		result, err := zipQuery(arguments.Args["city"], arguments.Args["state"])
		if err == nil {
			fmt.Printf("\033[1m%s, %s\033[0m\n", result.City, result.State)
			if len(result.Zips) == 0 {
				fmt.Println("No results were found.")
			} else {
				for i := 0; i < len(result.Zips); i++ {
					fmt.Println(result.Zips[i])
				}
			}
		}
	case "address":
		result, err := addressQuery(arguments.Args["company"], arguments.Args["address1"], arguments.Args["address2"], arguments.Args["city"], arguments.Args["state"], arguments.Args["zip"])
		if err == nil {
			fmt.Printf("\033[1m%s, %s %s\033[0m\n", arguments.Args["city"], arguments.Args["state"], arguments.Args["zip"])
			if len(result.Addresses) == 0 {
				fmt.Println("No results were found.")
			} else {
				for i := 0; i < len(result.Addresses); i++ {
					fmt.Println(result.Addresses[i].StringFormatted())
				}
			}
		}
	}
}
