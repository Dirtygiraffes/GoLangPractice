package main

import (
  "fmt"
  "bufio"
  "os"
)


func main() {
  fmt.Println("What is your name?")

  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()

  name := scanner.Text()

  fmt.Println("Hello, " + name + ", nice to meet you!")
}
