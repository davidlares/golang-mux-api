package connect

import (
  "log"
  "github.com/jinzhu/gorm"
 _ "github.com/jinzhu/gorm/dialects/mysql"
 "../structures"
)

var connection *gorm.DB
const engine_sql string = "mysql"

// should be secrets or ENV variables
const username string = "root"
const password string = ""
const database string = "go_api"

func InitializeDatabase(){
  connection = ConnectORM(CreateString())
  log.Println("DB Connection successfully created")
}

func CloseConnection(){
  connection.Close()
  log.Println("DB Connection successfully closed")
}

// establishing ORM connection with MySQL
func ConnectORM(stringConnection string) *gorm.DB {
  connection, err := gorm.Open(engine_sql, stringConnection)
  if err != nil{
    log.Fatal(err)
    return nil
  }
  return connection
}

func CreateString() string {
  return username + ":" + password + "@/" + database
}

func GetUser(id string) structures.User {
  user := structures.User{}
  connection.Where("id = ?", id).First(&user)
  return user
}
