// url https://www.geeksforgeeks.org/how-to-delete-or-remove-a-file-in-golang/

// Golang program to illustrate how to 
// remove files from the default directory
package main
   
import (
    "log"
    "os"
)
   
func main() {
  // Removing file from the directory
  // Using Remove() function
  
  e1 := os.Remove("index.html")
  if e1 != nil {
      log.Fatal(e1)
  }
  
  e2 := os.Remove("style.css")
  if e2 != nil {
      log.Fatal(e2)
  }
}