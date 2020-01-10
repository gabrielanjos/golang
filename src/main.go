package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("C:\\Users\\gabriel.anjos\\Downloads\\base_teste.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // Pula a primeira linha (Cabeçalho do arquivo)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		fmt.Println(line, len(line))
		fmt.Printf("CPF: %v\r\n", line[0])
		fmt.Printf("PRIVATE: %v\r\n", line[1])
		fmt.Printf("INCOMPLETO: %v\r\n", line[2])
		fmt.Printf("DATA DA ÚLTIMA COMPRA: %v\r\n", line[3])
		fmt.Printf("TICKET MÉDIO : %v\r\n", line[4])
		fmt.Printf("TICKET DA ÚLTIMA COMPRA: %v\r\n", line[5])
		fmt.Printf("LOJA MAIS FREQUÊNTE: %v\r\n", line[6])
		fmt.Printf("LOJA DA ÚLTIMA COMPRA: %v\r\n", line[7])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
