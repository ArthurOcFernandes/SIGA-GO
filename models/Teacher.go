package models

import (
	"log"
	"siga/database"
)

type Teacher struct {
	id                   int
	name                 string
	socialSecurityNumber string
	email                string
	address              Address
}

func ConstructTeacher(id int, name string, socialSecurityNumber string, email string, address Address) *Teacher {
	return &Teacher{id: id, name: name, socialSecurityNumber: socialSecurityNumber, email: email, address: address}
}

func (t *Teacher) Name() string {
	return t.name
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

func searchTeacherFromDataBase() []Teacher {

	selectAll := database.SelectAllUserJoinAddress("teachers")

	var teachers []Teacher

	for selectAll.Next() {
		var idTeacher, idAddress, addressNumber int
		var name, socialSecurityNumber, email, addressStreet, addressPostalCode, addressNeighbourhood, addressCity, addressState string

		err := selectAll.Scan(&idAddress, &idTeacher, &name, &socialSecurityNumber, &email, &addressStreet, &addressNumber, &addressPostalCode, &addressNeighbourhood, &addressCity, &addressState)

		if err != nil {
			log.Fatal(err)
		}

		address := ConstructAdress(addressStreet, addressNumber, addressPostalCode, addressNeighbourhood, addressCity, addressState)
		teacher := ConstructTeacher(idTeacher, name, socialSecurityNumber, email, *address)

		teachers = append(teachers, *teacher)
	}

	return teachers
}
