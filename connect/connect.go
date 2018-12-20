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

func GetUser(id string) structures.User {
  user := structures.User{}
  connection.Where("id = ?", id).First(&user)
  return user
}

func CreateUser(user structures.User) structures.User {
  connection.Create(&user) // and ID its assigned
  return user
}

func UpdateUser(id string, user structures.User) structures.User {
  currentUser := structures.User{}
  connection.Where("id = ?", id).First(&currentUser)
  currentUser.Username = user.Username
  currentUser.First_Name = user.First_Name
  currentUser.Last_Name = user.Last_Name
  connection.Save(&currentUser)
  return currentUser
}

func DeleteUser(id string){
  user := structures.User{}
  connection.Where("id = ?", id).First(&user)
  connection.Delete(&user)
}

func CreateString() string {
  return username + ":" + password + "@/" + database
}
