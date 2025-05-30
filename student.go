package main

import "github.com/spf13/cobra"

type Student struct {
	Name string
	Email string
	Idade int
}

var StudentCmd = &cobra.Command{
	Use: "student",
	Short: "Comando relacionado ao contexto",
}