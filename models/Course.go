package models

type Course struct {
	id       int
	name     string
	students []Student
	teachers []Teacher
	classes  string
}

func (c Course) GetId() int {
	return c.id
}

func (c Course) GetName() string {
	return c.name
}

func (c Course) GetStudents() []Student {
	return c.students
}

func (c Course) GetClasses() string {
	return c.classes
}

func (c Course) GetTeachers() []Teacher {
	return c.teachers
}
