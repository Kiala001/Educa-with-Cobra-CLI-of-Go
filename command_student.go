package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	nome string
	email string
	idade int
)
var repo = NewStudentRepositoryMemory()

var RegisterCmd = &cobra.Command{
	Use: "register",
	Short: "Registra um novo estudante",
	Run: func (cmd *cobra.Command, args []string)  {
		student := Student{
			Name: nome,
			Email: email,
			Idade: idade,
		}

		service := NewStudentService(repo)
		usecase := NewStudentUseCase(service)

		usecase.service.Register(student)

		len := repo.Count()

		if len == 0 {
			fmt.Println("Erro ao registar o estudante.")
		}

		fmt.Printf("Estudante %s registado com email %s e idade %d.", nome, email, idade)
		fmt.Println()
	},
}

var ListCmd = &cobra.Command{
	Use: "list",
	Short: "Lista todos estudantes cadastados.",
	Run: func (cmd *cobra.Command, args []string)  {

		if repo.Count() == 0 {
			fmt.Println("Nenhum estudante registado.")
		}

		if repo.Count() > 0 {
			students := repo.GetAll()
			fmt.Println("Lista de Estudantes")
			for stud := range students {
				fmt.Println(stud)
			}
		}
		
	},
}


func init() {
	RegisterCmd.Flags().StringVar(&nome, "nome", "", "Nome do estudante")
	RegisterCmd.Flags().StringVar(&email, "email", "", "Email do estudante")
	RegisterCmd.Flags().IntVar(&idade, "idade", 0, "Idade do estudante")

	RegisterCmd.MarkFlagRequired("nome")
	RegisterCmd.MarkFlagRequired("email")
	RegisterCmd.MarkFlagRequired("idade")

	StudentCmd.AddCommand(RegisterCmd)
	StudentCmd.AddCommand(ListCmd)
}