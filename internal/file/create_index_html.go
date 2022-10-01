/*
  create_index_html.go
  Create index.html file as static content for serving
  
  Sparisoma Viridi
  https://github.com/dudung
  
  20221001 create this by modifying some examples (Tsoukalos, 2022).
*/
package main

import (
    "fmt"
    "os"
)

func main() {
  fname := "index.html"
  dest, err := os.Create(fname)
  if err != nil {
    fmt.Println("os.Create:", err)
    return
  }
  defer dest.Close()
  
  fmt.Fprintf(dest, "<!doctype html>\n")
  fmt.Fprintf(dest, "<html lang='en'>\n")
  fmt.Fprintf(dest, "  <head>\n")
  fmt.Fprintf(dest, "    <title>Hello Wordl</title>\n")
  fmt.Fprintf(dest, "  </head>\n")
  fmt.Fprintf(dest, "  <body>\n")
  fmt.Fprintf(dest, "    <h1>Hello, Wordl!</h1>\n")
  fmt.Fprintf(dest, "  </body>\n")
  fmt.Fprintf(dest, "</html>\n")
}