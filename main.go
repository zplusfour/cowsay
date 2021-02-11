// Cowsay go brrrrr

package main

import (
  "strings"
  "io/ioutil"
  "fmt"
  "os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Insert your text.")
  } else {
    data, err := ioutil.ReadFile("text/cowsay.txt")
    check(err)

    fmt.Println(strings.ReplaceAll(string(data), "xxxxx", os.Args[1]))
  }
}
