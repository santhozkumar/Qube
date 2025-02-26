package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/santhozkumar/Qube/model"
)

func CheckAndPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	f, err := os.Open("cities.csv")
	CheckAndPanic(err)
	r := csv.NewReader(f)
	db := model.NewDataStore()
	r.Read()
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		r := model.NewRow(row)
		db.IngestData(r)
	}

	model.ExtractPermissionStruct()
	// create permission
	// permission := model.Permission{d: "DISTRIBUTOR1",
	// 	entries: []model.AuthorizationEntry{
	// 		{isInclude: true, region: "India"},
	// 		{isInclude: true, region: "United States"},
	// 		{isInclude: false, region: "Karnataka-India"},
	// 		// {isInclude: false, region: "Chennai-Tamil Nadu-India"},
	// 		{isInclude: false, region: "Tamil Nadu-India"},
	// 	},
	// }

	// err = db.AddPermission(permission)
	// if err != nil {
	// 	log.Print(err)
	// }
	// CheckAndPanic(err)
	// c, err := db.getCountry("India")
	// fmt.Println(c.distributors)
	// fmt.Println(db.getCity("Chennai", "Tamil Nadu", "India"))
	// fmt.Println("middle")
	// fmt.Println(db.check_permission("DISTRIBUTOR1", "Punjab-India"))
	// fmt.Println(db.check_permission("DISTRIBUTOR1", "India"))
	// fmt.Println(db.check_permission("DISTRIBUTOR1", "Tamil Nadu-India"))
}
