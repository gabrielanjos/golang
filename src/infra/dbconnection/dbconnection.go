package dbconnection

// criar aqui a função de criação da conexão
import (
	"fmt"
	"os"

	c "github.com/gabrielanjos/golang/domain/clientes"

	"github.com/jinzhu/gorm"

	// test comment to avoid error in _
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/joho/godotenv"
)

var db *gorm.DB //database

func init() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	//fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn

	db.Debug().AutoMigrate(&c.DadosCliente{}) //Database migration
}

//GetDB - returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}

// para chamar de outro pacote você vai utilizar..
// db := db_connection.GetDB()
// db.Create(&logSF) // vai passar como parametro tua variável que tem recebeu a struct, mas um ponteiro dela, por isso &
