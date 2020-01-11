package clientes

import (
	"log"

	"github.com/gabrielanjos/golang/src/infra/dbconnection"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

//DadosCliente adasdas
type DadosCliente struct {
	Cpf                      string  `gorm:"column:cpf"`
	CpfValido                bool    `gorm:"column:cpfValido"`
	Private                  int     `gorm:"column:private"`
	Incompleto               int     `gorm:"column:incompleto"`
	DataUltimaCompra         string  `gorm:"column:dataUltimaCompra"`
	TicketMedio              float64 `gorm:"column:ticketMedio"`
	TicketUltimaCompra       float64 `gorm:"column:ticketUltimaCompra"`
	LojaMaisFrequente        string  `gorm:"column:lojaMaisFrequente"`
	LojaMaisFrequenteValido  bool    `gorm:"column:lojaMaisFrequenteValido"`
	LojaDaUltimaCompra       string  `gorm:"column:lojaDaUltimaCompra"`
	LojaDaUltimaCompraValido bool    `gorm:"column:lojaDaUltimaCompraValido"`
}

//TableName Nome da tabela no banco de dados
func (dadosCliente *DadosCliente) TableName() string {
	return "dadoscliente"

}

//InserirRegistros no banco de dados
func InserirRegistros(insertRecords []interface{}) {
	db := dbconnection.GetDB()
	db.Debug().AutoMigrate(&DadosCliente{}) //Database migration

	err := gormbulk.BulkInsert(db, insertRecords, 3000)
	if err != nil {
		log.Panic(err)
	}

}
