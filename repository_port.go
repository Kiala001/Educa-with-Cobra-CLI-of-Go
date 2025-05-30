package main

type Repository_Interface interface{
	Save(student Student)
	Count() int
	GetAll() []Student
	Delete(student_id string)
}