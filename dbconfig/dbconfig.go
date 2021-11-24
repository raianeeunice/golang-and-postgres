package dbconfig

import "fmt"

type Filme struct {
	ID    int
	Title string
	Body  []byte
}

const (
	PostgresDriver = ""

	User = ""

	Host = "localhost"

	Port = "5432"

	Password = ""

	DbName = ""

	TableName = ""
)

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)