package clientes

import (
	"log"

	"github.com/gabrielanjos/golang/src/infra/dbconnection"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

//DadosCliente adasdas
type DadosCliente struct {
	Cpf                      string  `gorm:"column:cpf"`
	Cpfvalido                bool    `gorm:"column:cpfvalido"`
	Private                  int     `gorm:"column:private"`
	Incompleto               int     `gorm:"column:incompleto"`
	Dataultimacompra         string  `gorm:"column:dataultimacompra"`
	Ticketmedio              float64 `gorm:"column:ticketmedio"`
	Ticketultimacompra       float64 `gorm:"column:ticketultimacompra"`
	Lojamaisfrequente        string  `gorm:"column:lojamaisfrequente"`
	Lojamaisfrequentevalido  bool    `gorm:"column:lojamaisfrequentevalido"`
	Lojadaultimacompra       string  `gorm:"column:lojamaUltimacompra"`
	Lojadaultimacompravalido bool    `gorm:"column:lojamaUltimacompraValido"`
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
