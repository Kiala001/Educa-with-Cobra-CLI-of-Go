package main

type Student_UseCase struct {
	service *Student_Service
}

func NewStudentUseCase(service *Student_Service) *Student_UseCase {
	return &Student_UseCase{service: service}
}