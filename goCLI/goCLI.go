package main

// rodar: go run goCLI.go -nome="João" -idade=30


import (
	"fmt"
	"flag"
)

func goCLI(){
	// define as flags
	nome := flag.String("nome", "Usuário", "O nome do usuário")
	idade := flag.Int("idade", 20, "A idade do usuáiro")
	// parse as flags
	flag.Parse()
	// print as flags
	fmt.Printf("Olá, %s! Você tem %d anos.\n", *nome, *idade)
}

func main() {
	goCLI()
}