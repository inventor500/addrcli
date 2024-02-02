# addrcli

This is a simple program that retrieves the list of possible zip codes for a given city / state combination.

This program uses the USPS website as a back end, and thus (probably) only works for addresses in the United States.


## Usage

There are three supported query types: address, zip, and city.

### Address
The `address` query looks for an address when given only partial information. This corresponds to the [byaddress](https://tools.usps.com/zip-code-lookup.htm?byaddress) query on the USPS website.

There are six possible flags for this query: `company`, `address1`, `address2`, `city`, `state`, and `zip`. All flags are optional, but if none are provided, there will be no search results.

Example usage:
```shell
$ addrcli address --address1 '620 Eighth Avenue' --city "New York" --state "NY"
# New York, NY 
# 620 8TH AVE
# NEW YORK, NY 10018-1618
# ...
```

### City
The `city` query looks for cities matching a given zip code. This corresponds to the [citybyzipcode](https://tools.usps.com/zip-code-lookup.htm?citybyzipcode) query on the USPS website.

There is one possible flag for this query: `zip`. This is the 5 digit zip code for the city.

Example usage:
```shell
$ addrcli city --zip 10018
# 10018
# NEW YORK, NY 10018 (STANDARD)
#
# 5 Not Accepted Cities
# MANHATTAN, NY
# NEW YORK CITY, NY
# NY, NY
# NY CITY, NY
# NYC, NY
```

### Zip
The `zip` query looks for zip codes matching a given city/state combination. This corresponds to the [bycitystate](https://tools.usps.com/zip-code-lookup.htm?bycitystate) query on the USPS website.

There are two possible flags for this query: `city` and `state`.

Example usage:
```shell
$ addrcli zip --city "New York" --state NY
# NEW YORK, NY
# 10001
# 10002
# 10003
# 10004
# 10005
# ...
```
