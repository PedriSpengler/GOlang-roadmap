package main
// rodar: go run goCLICobra.go -n "João" -i 30
// go get -u github.com/spf13/cobra – Instala o pacote cobra.
// go build -o saudacao – Gera o executável.

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var nome string
var idade int

// Definindo o comando principal
var rootCmd = &cobra.Command{
	Use:   "saudacao",
	Short: "Um simples programa de saudação",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Olá, %s! Você tem %d anos.\n", nome, idade)
	},
}

func init() {
	// Adicionando flags ao comando
	rootCmd.Flags().StringVarP(&nome, "nome", "n", "Usuário", "O nome do usuário")
	rootCmd.Flags().IntVarP(&idade, "idade", "i", 18, "A idade do usuário")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}