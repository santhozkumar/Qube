package main

type City struct {
	name         string
	code         string
	distributors []Distributor
}

func (c *City) addDistributor(distributor Distributor) {
	c.distributors = append(c.distributors, distributor)
}

func distributorAt(distributors []Distributor, d Distributor) (int, error) {
	for i, distributor := range distributors {
		if distributor == d {
			return i, nil
		}
	}

	return -1, ErrDistributorNotFound
}

func (c *City) removeDistributor(d Distributor) error {
	i, err := distributorAt(c.distributors, d)
	if err != nil {
		return err
	}
	c.distributors = append(c.distributors[:i], c.distributors[i+1:]...)
	return nil
}

func (c *City) isValidDistributor(distributor Distributor) bool {
	for _, d := range c.distributors {
		if d == distributor {
			return true
		}
	}
	return false
}

// func (c *City) removeDistributor(d Distributor) error {
// 	for i, cityDistributor := range c.distributors {
// 		if cityDistributor == d {
// 			c.distributors = append(c.distributors[:i], c.distributors[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return ErrDistributorNotFound
// }
