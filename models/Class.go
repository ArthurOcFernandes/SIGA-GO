package models

type Class struct {
	id       int
	name     string
	teacher  Teacher
	students []Student
	grades   map[Student]int
}

func (c Class) Students() []Student {
	return c.students
}

func (c Class) Teacher() Teacher {
	return c.teacher
}

func (c Class) Name() string {
	return c.name
}

func (c Class) Id() int {
	return c.id
}

func ConstructClass(id int, name string, teacher Teacher, students []Student) *Class {
	class := Class{id: id, name: name, teacher: teacher, students: students}
	class.grades = make(map[Student]int)

	for _, student := range students {
		class.grades[student] = 0
	}
	return &class
}
