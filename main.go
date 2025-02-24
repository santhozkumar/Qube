package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func CheckAndPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func NewRow(rowSplit []string) Row {
	if len(rowSplit) != 6 {
		log.Println("invalid row", rowSplit)
	}
	var r Row
	r.cityCode = rowSplit[0]
	r.stateCode = rowSplit[1]
	r.countryCode = rowSplit[2]
	r.cityName = rowSplit[3]
	r.stateName = rowSplit[4]
	r.countryName = rowSplit[5]
	return r
}

func main() {
	f, err := os.Open("cities.csv")
	r := csv.NewReader(f)
	CheckAndPanic(err)
	db := GeographicalDatabase{countries: make(map[string]Country)}
	var i int
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if i == 0 {
			i++
			continue
		}
		// if i > 5 {
		// 	break
		// }
		r := NewRow(row)
		db.add_data(r)
		i = i + 1
	}
	// v, ok := db.countries["India"]
	// fmt.Println(v, ok)
	// fmt.Println(len(v.states))

	// create permission
	permission := Permission{d: "DISTRIBUTOR1",
		entries: []AuthorizationEntry{
			{isInclude: true, region: "India"},
			// {isInclude: true, Region: "United States"},
			// {isInclude: false, region: "Karnataka-India"},
			// {isInclude: false, region: "Chennai-Tamil Nadu-India"},
		},
	}
	err = db.add_permission(permission)
	if err != nil {
		log.Print(err)
	}
	CheckAndPanic(err)
	c, err := db.getCountry("India")
	fmt.Println(c.distributors)
	fmt.Println(db.getCity("Chennai", "Tamil Nadu", "India"))
	fmt.Println("middle")
	fmt.Println(db.check_permission("DISTRIBUTOR1", "Punjab-India"))
	fmt.Println(db.check_permission("DISTRIBUTOR1", "India"))
}
