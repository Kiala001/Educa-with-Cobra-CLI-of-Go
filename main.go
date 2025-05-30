package main

import (
	"fmt"

	"github.com/spf13/cobra"
)


func main() {
    var rootCMD = &cobra.Command{
        Use: "School",
        Short: "Sistema de registo de estudante",
        Run: func (cmd *cobra.Command, args []string)  {
            fmt.Println("=== BEM-VINDO AO CLI ESCOLAR ===")
        },
    }

    rootCMD.AddCommand(RegisterCmd)
    rootCMD.AddCommand(ListCmd)
    rootCMD.AddCommand(StudentCmd)
    
    rootCMD.Execute()
}