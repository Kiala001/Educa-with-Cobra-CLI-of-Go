package main

type Student_Repository struct {
	students map[string]Student
}

func NewStudentRepositoryMemory() *Student_Repository {
	return &Student_Repository{students: make(map[string]Student)}
}

func (repo *Student_Repository) Save(student Student) {
	repo.students[student.Email] = student
}

func (repo *Student_Repository) Count() int{
	return len(repo.students)
}

func (repo *Student_Repository) GetAll() []Student {
	var list []Student
	for _, student := range repo.students {
		list = append(list, student)
	}
	return list
}

func (repo *Student_Repository) Delete(student_id string) {
	delete(repo.students, student_id)
}
