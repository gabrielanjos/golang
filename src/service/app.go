package service

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	c "github.com/gabrielanjos/golang/src/domain/clientes"
	u "github.com/gabrielanjos/golang/src/util"
)

//ReadAndManipulateFile função responsável por ler e manipular o arquivo de texto/csv
func ReadAndManipulateFile() {

	//dadosClientes := []c.DadosCliente{}
	var dadosClientes []interface{}
	file, err := os.Open("base_teste.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // Pula a primeira linha (Cabeçalho do arquivo)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		//Convertendo campos  do arquivo para o tipo correto da varíavel na struct
		private, _ := strconv.Atoi(line[1])
		incompleto, _ := strconv.Atoi(line[2])
		ticketMedio, _ := strconv.ParseFloat(line[4], 64)
		ticketUltimaCompra, _ := strconv.ParseFloat(line[5], 64)

		n := c.DadosCliente{
			Cpf:                      line[0],
			Cpfvalido:                u.ValidaCPF(line[0]),
			Private:                  private,
			Incompleto:               incompleto,
			Dataultimacompra:         line[3],
			Ticketmedio:              ticketMedio,
			Ticketultimacompra:       ticketUltimaCompra,
			Lojamaisfrequente:        line[6],
			Lojamaisfrequentevalido:  u.ValidaCNPJ(line[6]),
			Lojadaultimacompra:       line[7],
			Lojadaultimacompravalido: u.ValidaCNPJ(line[7]),
		}

		dadosClientes = append(dadosClientes, n)
	}

	c.InserirRegistros(dadosClientes)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
