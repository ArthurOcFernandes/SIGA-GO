package models

import (
	"log"
	"siga/database"
)

type Student struct {
	id                   int
	name                 string
	studentRegistry      string
	socialSecurityNumber string
	email                string
	address              Address
}

func (s Student) Name() string {
	return s.name
}

func (s Student) StudentRegistry() string {
	return s.studentRegistry
}

func (s Student) SocialSecurityNumber() string {
	return s.socialSecurityNumber
}

func (s Student) Email() string {
	return s.email
}

func (s Student) Address() Address {
	return s.address
}

func ConstructStudent(name, registry, socialSecurityNumber, email string, address Address) *Student {
	return &Student{name: name, studentRegistry: registry, socialSecurityNumber: socialSecurityNumber, email: email, address: address}
}

func SearchStudentFromDataBase() []Student {
	db := database.Connection()

	selectAll, err := db.Query("SELECT * FROM students inner join addresses using (id_address)")

	if err != nil {
		log.Fatal(err)
	}

	var students []Student

	for selectAll.Next() {
		var idStudent int
		var name, studentRegistry, socialSecurityNumber, email string
		var idAddress, number int
		var street, postalCode, neighbourhood, city, state string

		err = selectAll.Scan(&idAddress, &idStudent, &name, &socialSecurityNumber, &email, &studentRegistry, &street, &number, &postalCode, &neighbourhood, &city, &state)

		if err != nil {
			panic(err.Error())
		}

		address := ConstructAdress(street, number, postalCode, neighbourhood, city, state)
		student := ConstructStudent(name, studentRegistry, socialSecurityNumber, email, *address)

		students = append(students, *student)

	}

	defer db.Close()

	return students
}
