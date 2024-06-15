package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
    f2, err := os.Open("no-file.txt")
    if err != nil {
        log.Println("err happened", err)                                          // the log will log to this file.
    }
    defer f2.Close()
	fmt.Println("done")
}

