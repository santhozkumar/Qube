package main

// import "fmt"

// import "log"
type Distributor = string

type Film struct {
	name         string
	distributors []Distributor
}

type Row struct {
	cityCode    string
	stateCode   string
	countryCode string
	cityName    string
	stateName   string
	countryName string
}

//	Permissions for DISTRIBUTOR1
//
// INCLUDE: INDIA
// INCLUDE: UNITEDSTATES
// EXCLUDE: KARNATAKA-INDIA
// EXCLUDE: CHENNAI-TAMILNADU-INDIA

type AuthorizationEntry struct {
	isInclude bool
	region    string
}

type Authorization struct {
	permissions []Permission
}

type Permission struct {
	d                 Distributor
	entries           []AuthorizationEntry
	parentDistributor *Permission
}

//
// func distributorAt(distributors []Distributor, d Distributor) (int, error) {
// 	for i, distributor := range distributors {
// 		if distributor == d {
// 			return i, nil
// 		}
// 	}
// 	return -1, ErrDistributorNotFound
//
// }
