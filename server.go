package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
)

// JSON Struct
type User struct {
  Username string `json:"username"`
  First_Name string `json:"first_name"`
  Last_Name string `json:"last_name"`
}

func main(){
  // fmt.Println("Hello, World")
  r := mux.NewRouter()
  // handling URLs
  r.HandleFunc("/user/", GetUser).Methods("GET")
  log.Println("Server running on port 8000")
  log.Fatal(http.ListenAndServe(":8000", r))
}

func GetUser(w http.ResponseWriter, r* http.Request){
  // w.Write([]byte("Gorilla \n"))
  user := User {"davidlares", "David","Lares"}
  // converting struct into response
  json.NewEncoder(w).Encode(user)
}
