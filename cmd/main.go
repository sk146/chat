package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./web/static")))

  log.Println("Listening on :3000...")
  err :=  http.ListenAndServe(":3000", nil)
  if err != nil {
    log.Fatal(err)
  }
}
