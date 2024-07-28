package models

type Address struct {
	id            int
	street        string
	number        int
	postalCode    string
	neighbourhood string
	city          string
	state         string
}

func (a Address) getStreet() string {
	return a.street
}

func (a Address) getNumber() int {
	return a.number
}

func (a Address) getPostalCode() string {
	return a.postalCode
}

func (a Address) getNeighbourhood() string {
	return a.neighbourhood
}

func (a Address) getCity() string {
	return a.city
}

func (a Address) getState() string {
	return a.state
}

func ConstructAdress(street string, number int, postalCode, neighbourhood, city, state string) *Address {
	return &Address{street: street, number: number, postalCode: postalCode, neighbourhood: neighbourhood, city: city, state: state}
}
