package main

import (
    "fmt"
    "os"
)

func check(e error) {
    if (e != nil) {
        panic(e)
    }
}

func main() {
    var path string = "input.txt"
    data, err := os.ReadFile(path)
    check(err)
    fmt.Println(data)
}
