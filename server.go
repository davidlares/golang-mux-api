package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
  "./connect"
  "./structures"
)

func main(){
  // DB connection
  connect.InitializeDatabase()
  defer connect.CloseConnection()

  // fmt.Println("Hello, World")
  r := mux.NewRouter()
  // handling URLs
  r.HandleFunc("/user/{id}", GetUser).Methods("GET")
  r.HandleFunc("/user/new", NewUser).Methods("POST")

  log.Println("Server running on port 8000")
  log.Fatal(http.ListenAndServe(":8000", r))
}

func GetUser(w http.ResponseWriter, r* http.Request){
  // getting ID
  vars := mux.Vars(r)
  user_id := vars["id"]

  status := "success"
  user := connect.GetUser(user_id)
  if(user.Id <= 0){
    status = "error"
    message = "User not found."
  }
  response := structures.Response{status, user, message}
  // converting struct into response
  json.NewEncoder(w).Encode(response)
}

func NewUser(w http.ResponseWriter, r* http.Request){
  user := GetUserRequest(r)
  response := structures.Response{"success", connect.CreateUser(user), ""}
  json.NewEncoder(w).Encode(response)
}

func GetUserRequest(r* http.Request) structures.User {
  var user structures.User
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&user)
  if err != nil {
    log.Fatal(err)
  }
  return user
}
