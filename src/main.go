package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
)

type dadosCliente struct {
	cpf                      string
	cpfValido                bool
	private                  int
	incompleto               int
	dataUltimaCompra         string
	ticketMedio              float64
	ticketUltimaCompra       float64
	lojaMaisFrequente        string
	lojaMaisFrequenteValido  bool
	lojaDaUltimaCompra       string
	lojaDaUltimaCompraValido bool
}

func main() {
	db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")

	dadosClientes := []dadosCliente{}
	file, err := os.Open("C:\\Users\\gabriel.anjos\\Downloads\\base_teste.txt")
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

		n := dadosCliente{
			cpf:                      line[0],
			cpfValido:                ValidaCPF(line[0]),
			private:                  private,
			incompleto:               incompleto,
			dataUltimaCompra:         line[3],
			ticketMedio:              ticketMedio,
			ticketUltimaCompra:       ticketUltimaCompra,
			lojaMaisFrequente:        line[6],
			lojaMaisFrequenteValido:  validaCNPJ(line[6]),
			lojaDaUltimaCompra:       line[7],
			lojaDaUltimaCompraValido: validaCNPJ(line[7]),
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

	defer db.Close()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

//ValidaCPF verifica se um CPF é válido
func ValidaCPF(cpf string) bool {
	cpf = strings.Replace(cpf, ".", "", -1)
	cpf = strings.Replace(cpf, "-", "", -1)
	if len(cpf) != 11 {
		return false
	}
	var eq bool
	var dig string
	for _, val := range cpf {
		if len(dig) == 0 {
			dig = string(val)
		}
		if string(val) == dig {
			eq = true
			continue
		}
		eq = false
		break
	}
	if eq {
		return false
	}

	i := 10
	sum := 0
	for index := 0; index < len(cpf)-2; index++ {
		pos, _ := strconv.Atoi(string(cpf[index]))
		sum += pos * i
		i--
	}

	prod := sum * 10
	mod := prod % 11
	if mod == 10 {
		mod = 0
	}
	digit1, _ := strconv.Atoi(string(cpf[9]))
	if mod != digit1 {
		return false
	}
	i = 11
	sum = 0
	for index := 0; index < len(cpf)-1; index++ {
		pos, _ := strconv.Atoi(string(cpf[index]))
		sum += pos * i
		i--
	}
	prod = sum * 10
	mod = prod % 11
	if mod == 10 {
		mod = 0
	}
	digit2, _ := strconv.Atoi(string(cpf[10]))
	if mod != digit2 {
		return false
	}

	return true
}

//validaCNPJ verifica se um CNPJ é válido
func validaCNPJ(cnpj string) bool {
	cnpj = strings.Replace(cnpj, ".", "", -1)
	cnpj = strings.Replace(cnpj, "-", "", -1)
	cnpj = strings.Replace(cnpj, "/", "", -1)
	if len(cnpj) != 14 {
		return false
	}

	algs := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	var algProdCpfDig1 = make([]int, 12, 12)
	for key, val := range algs {
		intParsed, _ := strconv.Atoi(string(cnpj[key]))
		sumTmp := val * intParsed
		algProdCpfDig1[key] = sumTmp
	}
	sum := 0
	for _, val := range algProdCpfDig1 {
		sum += val
	}
	digit1 := sum % 11
	if digit1 < 2 {
		digit1 = 0
	} else {
		digit1 = 11 - digit1
	}
	char12, _ := strconv.Atoi(string(cnpj[12]))
	if char12 != digit1 {
		return false
	}
	algs = append([]int{6}, algs...)

	var algProdCpfDig2 = make([]int, 13, 13)
	for key, val := range algs {
		intParsed, _ := strconv.Atoi(string(cnpj[key]))

		sumTmp := val * intParsed
		algProdCpfDig2[key] = sumTmp
	}
	sum = 0
	for _, val := range algProdCpfDig2 {
		sum += val
	}

	digit2 := sum % 11
	if digit2 < 2 {
		digit2 = 0
	} else {
		digit2 = 11 - digit2
	}
	char13, _ := strconv.Atoi(string(cnpj[13]))
	if char13 != digit2 {
		return false
	}

	return true
}
