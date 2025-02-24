package main

// import "fmt"

type State struct {
	name         string
	code         string
	cities       map[string]*City
	distributors []Distributor
}

func (s *State) addDistributor(distributor Distributor) {
	s.distributors = append(s.distributors, distributor)
	for _, city := range s.cities {
		city.addDistributor(distributor)
		// if city.name == "Chennai" {
		// fmt.Println(city.name, city.distributors)
		// }
	}
}

func (s *State) removeDistributor(distributor Distributor) error {
	i, err := distributorAt(s.distributors, distributor)
	if err != nil {
		return err
	}
	s.distributors = append(s.distributors[:i], s.distributors[i+1:]...)
	for _, city := range s.cities {
		city.removeDistributor(distributor)
	}
	return nil
}

func (s *State) isValidDistributor(distributor Distributor) bool {
	// fmt.Println(s.name, s.distributors)
	for _, d := range s.distributors {
		if d == distributor {
			return true
		}
	}
	return false
}

func (s *State) add_city(name, code string) *City {
	if _, ok := s.cities[name]; !ok {
		s.cities[name] = &City{name: name, code: code}
	}
	return s.cities[name]
}
