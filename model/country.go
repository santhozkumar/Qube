package model

type Country struct {
	name         string
	code         string
	states       map[string]*State
	distributors []Distributor
}

func (c *Country) add_state(name, code string) *State {
	if _, ok := c.states[name]; !ok {
		c.states[name] = &State{name: name, code: code, cities: make(map[string]*City)}
	}
	return c.states[name]
}

func (c *Country) addDistributor(distributor Distributor) {
	c.distributors = append(c.distributors, distributor)
	for _, state := range c.states {
		state.addDistributor(distributor)
		// if state.name == "Punjab" {
		// 	fmt.Println(state.name, state.distributors)
		// }
	}
}

func (c *Country) removeDistributor(distributor Distributor) error {
	i, err := distributorAt(c.distributors, distributor)
	if err != nil {
		return err
	}
	c.distributors = append(c.distributors[:i], c.distributors[i+1:]...)
	for _, state := range c.states {
		state.removeDistributor(distributor)
	}
	return nil
}

func (c *Country) isValidDistributor(distributor Distributor) bool {
	for _, d := range c.distributors {
		if d == distributor {
			return true
		}
	}
	return false
}
