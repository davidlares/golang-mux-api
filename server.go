package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

func main(){
  // fmt.Println("Hello, World")
  r := mux.NewRouter()
  // handling URLs
  r.HandleFunc("/user/", GetUser).Methods("GET")
  log.Println("Server running on port 8000")
  log.Fatal(http.ListenAndServe(":8000", r))
}

func GetUser(w http.ResponseWriter, r* http.Request){
  w.Write([]byte("Gorilla \n"))
}
