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
	addressFlag := flag.NewFlagSet("address", flag.ExitOnError)
	zipFlag := flag.NewFlagSet("zip", flag.ExitOnError)
	cityFlag := flag.NewFlagSet("city", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Copyright (C) 2024 Isaac Ganoung.")
		fmt.Printf("Usage: %s address|city|zip [arg0...argn]\n", os.Args[0])
		os.Exit(1)
	}

	switch os.Args[1] {
	case "address":
		company := addressFlag.String("company", "", "Company")
		address1 := addressFlag.String("address1", "", "Address line 1")
		address2 := addressFlag.String("address2", "", "Address line 2")
		city := addressFlag.String("city", "", "City")
		state := addressFlag.String("state", "", "State")
		zip := addressFlag.String("zip", "", "Zip")
		addressFlag.Parse(os.Args[2:])
		addressMap := make(map[string]string)
		addressMap["company"] = *company
		addressMap["address1"] = *address1
		addressMap["address2"] = *address2
		addressMap["city"] = *city
		addressMap["state"] = *state
		addressMap["zip"] = *zip
		return &Arguments{"address", addressMap}
	case "city":
		zip := cityFlag.String("zip", "", "Zip 5")
		cityFlag.Parse(os.Args[2:])
		cityMap := make(map[string]string)
		cityMap["zip"] = *zip
		return &Arguments{"city", cityMap}
	case "zip":
		city := zipFlag.String("city", "", "City")
		state := zipFlag.String("state", "", "State")
		zipFlag.Parse(os.Args[2:])
		zipMap := make(map[string]string)
		zipMap["city"] = *city
		zipMap["state"] = *state
		return &Arguments{"zip", zipMap}

	default:
		fmt.Printf("Unsupported search type %s\n", os.Args[1])
		os.Exit(1)
	}
	return &Arguments{} // TOOD: Why does the compiler require this?
}

func main() {
	arguments := get_args()
	switch arguments.Subcommand {
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
		} else {
			print_error(err)
		}
	case "city":
		result, err := cityQuery(arguments.Args["zip"])
		if err == nil {
			fmt.Printf("\033[1m%s\033[0m\n", arguments.Args["zip"])
			fmt.Printf("%s, %s %s", result.City, result.State, result.Zip)
			if result.Caveat != "" {
				fmt.Printf(" (%s)\n", result.Caveat)
			} else {
				fmt.Println("")
			}
			if len(result.OtherCities) > 0 {
				fmt.Printf("\n\033[1m%d Other Cities\033[0m\n", len(result.OtherCities))
				for i := 0; i < len(result.OtherCities); i++ {
					fmt.Println(result.OtherCities[i])
				}
			}
			if len(result.NonAccept) > 0 {
				fmt.Printf("\n\033[1m%d Not Accepted Cities\033[0m\n", len(result.NonAccept))
				for i := 0; i < len(result.NonAccept); i++ {
					fmt.Println(result.NonAccept[i])
				}
			}
		} else {
			print_error(err)
		}
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
		} else {
			print_error(err)
		}
	}
}

func print_error(err error) {
	fmt.Printf("\033[91;1m%s\033[0m\n", err)
}
