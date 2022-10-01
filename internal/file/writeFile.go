// url https://www.linode.com/docs/guides/creating-reading-and-writing-files-in-go-a-tutorial/#creating-a-new-file:~:text=File%3A%20./writeFile.go

package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Please provide a filename")
        return
    }

    filename := os.Args[1]
    destination, err := os.Create(filename)
    if err != nil {
        fmt.Println("os.Create:", err)
        return
    }
    defer destination.Close()

    fmt.Fprintf(destination, "[%s]: ", filename)
    fmt.Fprintf(destination, "Using fmt.Fprintf in %s\n", filename)
}