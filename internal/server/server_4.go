// url 
https://www.alexedwards.net/blog/serving-static-sites-with-go#:~:text=is%20wonderfully%20compact%3A-,File%3A%20main.go,-package%20main%0A%0Aimport

package main

import (
	"log"
	"net/http"
)

func main() {
  dir := http.Dir("./static")
	fs := http.FileServer(dir)
	http.Handle("/", fs)

	log.Print("Listening on :8081...")
  
	err := http.ListenAndServe(":8081", nil)
	
  if err != nil {
		log.Fatal(err)
	}
}