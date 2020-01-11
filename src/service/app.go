package service

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	c "github.com/gabrielanjos/golang/domain/clientes"
	u "github.com/gabrielanjos/golang/util"
)

//ReadAndManipulateFile função responsável por ler e manipular o arquivo de texto/csv
func ReadAndManipulateFile() {

	//dadosClientes := []c.DadosCliente{}
	var dadosClientes []interface{}
	file, err := os.Open("../base_teste.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // Pula a primeira linha (Cabeçalho do arquivo)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		private, _ := strconv.Atoi(line[1])
		incompleto, _ := strconv.Atoi(line[2])
		ticketMedio, _ := strconv.ParseFloat(line[4], 64)
		ticketUltimaCompra, _ := strconv.ParseFloat(line[5], 64)

		n := c.DadosCliente{
			Cpf:                      line[0],
			CpfValido:                u.ValidaCPF(line[0]),
			Private:                  private,
			Incompleto:               incompleto,
			DataUltimaCompra:         line[3],
			TicketMedio:              ticketMedio,
			TicketUltimaCompra:       ticketUltimaCompra,
			LojaMaisFrequente:        line[6],
			LojaMaisFrequenteValido:  u.ValidaCNPJ(line[6]),
			LojaDaUltimaCompra:       line[7],
			LojaDaUltimaCompraValido: u.ValidaCNPJ(line[7]),
		}
		dadosClientes = append(dadosClientes, n)

		/*
			fmt.Println(line, len(line))
			fmt.Printf("CPF: %v\r\n", line[0])
			fmt.Printf("Valido?: %v\r\n:", ValidaCPF(line[0]))
			fmt.Printf("PRIVATE: %v\r\n", line[1])
			fmt.Printf("INCOMPLETO: %v\r\n", line[2])
			fmt.Printf("DATA DA ÚLTIMA COMPRA: %v\r\n", line[3])
			fmt.Printf("TICKET MÉDIO : %v\r\n", line[4])
			fmt.Printf("TICKET DA ÚLTIMA COMPRA: %v\r\n", line[5])
			fmt.Printf("LOJA MAIS FREQUÊNTE: %v\r\n", line[6])
			fmt.Printf("Valido?: %v\r\n:", validaCNPJ(line[6]))
			fmt.Printf("LOJA DA ÚLTIMA COMPRA: %v\r\n", line[7])
			fmt.Printf("Valido?: %v\r\n:", validaCNPJ(line[7]))
			time.Sleep(2 * time.Second)
		*/
	}

	c.InserirRegistros(dadosClientes)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
