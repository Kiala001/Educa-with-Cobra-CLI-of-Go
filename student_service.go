package main

type Student_Service struct {
	repository Repository_Interface
}

func NewStudentService(repo *Student_Repository) *Student_Service{
	return &Student_Service{repository: repo}
}

func (s *Student_Service) Register(student Student){
	s.repository.Save(student)
}