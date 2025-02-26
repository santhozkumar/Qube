package model

import (
	"strings"
)

type GeographicalDatabase struct {
	countries map[string]*Country
}

func (db *GeographicalDatabase) getCountry(name string) (*Country, error) {
	v, ok := db.countries[name]
	if !ok {
		return &Country{}, ErrCountryNotFound
	}
	return v, nil
}

func (db *GeographicalDatabase) getState(stateName, countryName string) (*State, error) {
	country, err := db.getCountry(countryName)
	if err != nil {
		return &State{}, err
	}
	v, ok := country.states[stateName]
	if !ok {
		return &State{}, ErrStateNotFound
	}
	return v, nil
}

func (db *GeographicalDatabase) getCity(cityName, stateName, countryName string) (*City, error) {
	state, err := db.getState(stateName, countryName)
	if err != nil {
		return &City{}, err
	}
	v, ok := state.cities[cityName]
	if !ok {
		return &City{}, ErrStateNotFound
	}
	return v, nil
}

func (db *GeographicalDatabase) check_permission(distributor, region string) bool {

	var countryName, stateName, cityName string
	regionSplit := strings.Split(region, "-")
	switch len(regionSplit) {
	case 1:
		countryName = regionSplit[0]
		country, err := db.getCountry(countryName)
		if err != nil {
			return false
		}
		return country.isValidDistributor(distributor)
	case 2:
		countryName, stateName = regionSplit[1], regionSplit[0]
		state, err := db.getState(stateName, countryName)
		if err != nil {
			return false
		}
		return state.isValidDistributor(distributor)
	case 3:
		countryName, stateName, cityName = regionSplit[2], regionSplit[1], regionSplit[0]
		city, err := db.getCity(cityName, stateName, countryName)
		if err != nil {
			return false
		}
		return city.isValidDistributor(distributor)
	}

	return false
}

func (db *GeographicalDatabase) AddPermission(p Permission) error {
	for _, entry := range p.entries {
		var countryName, stateName, cityName string
		regionSplit := strings.Split(entry.region, "-")
		switch len(regionSplit) {
		case 1:
			countryName = regionSplit[0]
			country, err := db.getCountry(countryName)
			if err != nil {
				return err
			}
			if entry.isInclude {
				country.addDistributor(p.Dname)
			} else {
				country.removeDistributor(p.Dname)
			}
		case 2:
			countryName, stateName = regionSplit[1], regionSplit[0]
			state, err := db.getState(stateName, countryName)
			if err != nil {
				return err
			}
			if entry.isInclude {
				state.addDistributor(p.Dname)
			} else {
				state.removeDistributor(p.Dname)
			}
		case 3:
			countryName, stateName, cityName = regionSplit[2], regionSplit[1], regionSplit[0]
			city, err := db.getCity(cityName, stateName, countryName)
			if err != nil {
				return err
			}
			if entry.isInclude {
				city.addDistributor(p.Dname)
			} else {
				city.removeDistributor(p.Dname)
			}
		}
	}
	return nil
}

func (db *GeographicalDatabase) IngestData(r Row) {
	if _, ok := db.countries[r.countryName]; !ok {
		db.countries[r.countryName] = &Country{name: r.countryName, code: r.countryCode, states: make(map[string]*State)}
	}
	country := db.countries[r.countryName]
	state := country.add_state(r.stateName, r.stateName)
	_ = state.add_city(r.cityName, r.cityCode)
}

func NewDataStore() GeographicalDatabase {
	return GeographicalDatabase{countries: make(map[string]*Country)}
}
