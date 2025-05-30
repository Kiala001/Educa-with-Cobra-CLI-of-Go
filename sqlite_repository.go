package main


import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStudentRepository struct {
	db *sql.DB
}

func NewSQLiteStudentRepository() (*SQLiteStudentRepository, error) {
	db, err := sql.Open("sqlite3", "student.db")
	if err != nil {
		return nil, err
	}

	repo := &SQLiteStudentRepository{db: db}

	err = repo.createTableIfNotExists()
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *SQLiteStudentRepository) createTableIfNotExists() error {
	query := `
	CREATE TABLE IF NOT EXISTS students (
		name TEXT,
		email TEXT PRIMARY KEY,
		idade INTEGER
	);`
	_, err := r.db.Exec(query)
	return err
}

func (r *SQLiteStudentRepository) Save(student Student) {
	_, err := r.db.Exec("INSERT INTO students (name, email, idade) VALUES (?, ?, ?)", student.Name, student.Email, student.Idade)
	if err != nil {
		fmt.Println("Erro ao salvar:", err)
	}
}

func (r *SQLiteStudentRepository) Count() int {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM students").Scan(&count)
	if err != nil {
		fmt.Println("Erro ao contar:", err)
		return 0
	}
	return count
}

func (r *SQLiteStudentRepository) GetAll() []Student {
	rows, err := r.db.Query("SELECT name, email, idade FROM students")
	if err != nil {
		fmt.Println("Erro ao listar:", err)
		return nil
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var s Student
		if err := rows.Scan(&s.Name, &s.Email, &s.Idade); err == nil {
			students = append(students, s)
		}
	}
	return students
}

func (r *SQLiteStudentRepository) Delete(email string) {
	_, err := r.db.Exec("DELETE FROM students WHERE email = ?", email)
	if err != nil {
		fmt.Errorf("erro ao deletar estudante: %w", err)
	}
}
