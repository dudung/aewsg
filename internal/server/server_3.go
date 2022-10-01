// url https://tutorialedge.net/golang/creating-simple-web-server-with-golang/#:~:text=Serving%20Content%20from%20a%20Directory

package main

import (
    "log"
    "net/http"
)

func main() {

    http.Handle("/", http.FileServer(http.Dir("./static")))

    log.Fatal(http.ListenAndServe(":8081", nil))
}