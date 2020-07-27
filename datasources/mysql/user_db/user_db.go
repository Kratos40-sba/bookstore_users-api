package user_db
// connecting our database
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)
const (
	mysql_user_username = "mysql_user_username"
	mysql_user_password = "mysql_user_password"
	mysql_user_host = "mysql_user_host"
	mysql_user_dbname = "mysql_user_dbname"
)
var (
	Client *sql.DB
	username = os.Getenv(mysql_user_username)
	password = os.Getenv(mysql_user_password)
	host = os.Getenv(mysql_user_host)
	dbname = os.Getenv(mysql_user_dbname)
)
func init(){
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		dbname,
		)
	var err error
	Client , err = sql.Open("mysql",datasourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping() ; err != nil {
		panic(err)
	}
	log.Println("database successfully started :) ")
}
