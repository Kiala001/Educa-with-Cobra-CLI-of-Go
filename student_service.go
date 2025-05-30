package main

type Student_Service struct {
	repository Repository_Interface
}

func NewStudentService(repo Repository_Interface) *Student_Service{
	return &Student_Service{repository: repo}
}

func (s *Student_Service) Register(student Student){
	s.repository.Save(student)
}

func (s *Student_Service) Remove(student_id string) {
	s.repository.Delete(student_id)
}