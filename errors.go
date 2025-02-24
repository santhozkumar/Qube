package main

import "errors"

var (
	ErrCountryNotFound     = errors.New("Country Not Found")
	ErrStateNotFound       = errors.New("State Not Found")
	ErrDistributorNotFound = errors.New("Distributor Not Found")
)
