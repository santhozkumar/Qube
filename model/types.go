package model

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Distributor = string

type Row struct {
	cityCode    string
	stateCode   string
	countryCode string
	cityName    string
	stateName   string
	countryName string
}

func NewRow(rowSplit []string) Row {
	if len(rowSplit) != 6 {
		log.Fatalln("invalid row", rowSplit)
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

//	Permissions for DISTRIBUTOR1
//
// INCLUDE: INDIA
// INCLUDE: UNITEDSTATES
// EXCLUDE: KARNATAKA-INDIA
// EXCLUDE: CHENNAI-TAMILNADU-INDIA
type Permission struct {
	Dname             Distributor `json:"Distributor"`
	Includes          []string    `json:"Includes"`
	Excludes          []string    `json:"Excludes"`
	ParentDistributor Distributor `json:"Parent,omitempty"`
	entries           []AuthorizationEntry
}

type PermissionArray struct {
	PArray []Permission `json:"Data"`
}

func ExtractPermissionStruct() {
	f, _ := os.Open("~/Projects/Qube/authorizations.json")
	permissions := PermissionArray{}
	err := json.NewDecoder(f).Decode(&permissions)
	if err != nil {
		log.Println(err)
		fmt.Println("yup")
	}
	fmt.Println(permissions)
}

// INCLUDE: INDIA
type AuthorizationEntry struct {
	isInclude bool
	region    string
}

// type Authorization struct {
// 	permissions     []Permission
// 	distPermissions map[Distributor]*Permission
// }
