package main

import (
	"fmt"
	"testing"
)

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


func TestListAllStudentInRepository(t *testing.T) {
	student1 := Student{
		Name: "Kiala Emanuel",
		Email: "kiala@gmail.com",
		Idade: 22,
	}

	student2 := Student{
		Name: "Rui Manuel",
		Email: "rui@gmail.com",
		Idade: 22,
	}

	student_repository := NewStudentRepositoryMemory()
	student_service := NewStudentService(student_repository)
	student_usecase := NewStudentUseCase(student_service)
	student_usecase.service.Register(student1)
	student_usecase.service.Register(student2)

	students := student_repository.GetAll()

	for key := range students {
		fmt.Println("Estudante: ", students[key].Name)
	}
}

func TestRemoveStudentInRepository(t *testing.T) {
	student := Student{
		Name: "Kiala",
		Email: "kiala@gmail.com",
		Idade: 22,
	}

	student_repository := NewStudentRepositoryMemory()
	student_service := NewStudentService(student_repository)
	student_usecase := NewStudentUseCase(student_service)
	student_usecase.service.Register(student)

	fmt.Println("Estudantes: ",student_repository.Count())
	student_usecase.service.Remove(student.Email)

	len := student_repository.Count()
	if len != 0 {
		t.Errorf("Esperava %d mas obteve %d", 0, len)
	}
}