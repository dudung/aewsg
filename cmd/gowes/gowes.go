package main

import (
	"log"
	"net/http"
)

func main() {
  create_index_html()
  create_style_css()
  
  
  dir := http.Dir("./static")
	fs := http.FileServer(dir)
	http.Handle("/", fs)

	log.Print("Listening on :8081...")
  
	err := http.ListenAndServe(":8081", nil)
	
  if err != nil {
		log.Fatal(err)
	}
}