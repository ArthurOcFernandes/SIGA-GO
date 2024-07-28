package models

type Teacher struct {
	id                   int
	socialSecurityNumber string
	email                string
	address              Address
}

func ConstructTeacher(socialSecurityNumber string, email string, address Address) *Teacher {
	return &Teacher{socialSecurityNumber: socialSecurityNumber, email: email, address: address}
}

func (t *Teacher) Id() int {
	return t.id
}

func (t *Teacher) SocialSecurityNumber() string {
	return t.socialSecurityNumber
}
func (t *Teacher) Email() string {
	return t.email
}
func (t *Teacher) Address() Address {
	return t.address
}
