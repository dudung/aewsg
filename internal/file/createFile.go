// url https://www.linode.com/docs/guides/creating-reading-and-writing-files-in-go-a-tutorial/#creating-a-new-file:~:text=File%3A%20./createFile.go

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
    var _, err = os.Stat(filename)

    if os.IsNotExist(err) {
        file, err := os.Create(filename)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()
    } else {
        fmt.Println("File already exists!", filename)
        return
    }

    fmt.Println("File created successfully", filename)
}