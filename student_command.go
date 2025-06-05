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

var repo, _ = NewSQLiteStudentRepository()

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
			
			for key := range students {
				fmt.Println("Nome: ",students[key].Name,",Email: ",students[key].Email," e Idade: ",students[key].Idade)
			}
		}
		
	},
}

var DeleteCmd = &cobra.Command{
	Use: "delete",
	Short: "Deleta um estudante do repositório",
	Run: func (cmd *cobra.Command, args []string)  {
		Email := email

		service := NewStudentService(repo)
		usecase := NewStudentUseCase(service)

		usecase.service.Remove(Email)
		fmt.Println("Estudante com o codigo ",Email," deletado com sucesso.")
	},
}

func init() {
	RegisterCmd.Flags().StringVar(&nome, "nome", "", "Nome do estudante")
	RegisterCmd.Flags().StringVar(&email, "email", "", "Email do estudante")
	RegisterCmd.Flags().IntVar(&idade, "idade", 0, "Idade do estudante")

	RegisterCmd.MarkFlagRequired("nome")
	RegisterCmd.MarkFlagRequired("email")
	RegisterCmd.MarkFlagRequired("idade")

	DeleteCmd.Flags().StringVar(&email, "id", "", "Identificador único do estudante")
	DeleteCmd.MarkFlagRequired("id")
	
	StudentCmd.AddCommand(RegisterCmd)
	StudentCmd.AddCommand(ListCmd)
	StudentCmd.AddCommand(DeleteCmd)
}