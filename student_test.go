package main

import "testing"

func TestRegisterStudentInRepository (t *testing.T) {
	student := Student{
		Name: "Kiala",
		Email: "kiala@gmail.com",
		Idade: 22,
	}

	student_repository := NewStudentRepositoryMemory()
	student_service := NewStudentService(student_repository)
	student_usecase := NewStudentUseCase(student_service)
	student_usecase.service.Register(student)

	student_in_repo := student_repository.Count()

	if student_in_repo != 1 {
		t.Errorf("Esperava %d mas, obteve %d", 1, student_in_repo)
	}
}