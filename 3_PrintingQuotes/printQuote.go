package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  fmt.Println("What is the quote?")

  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  quote := scanner.Text()

  fmt. Println("Who said it?")
  scanner.Scan()
  speaker := scanner.Text()

  fmt.Printf(speaker + " says, " + "\"" + quote + "\"")
}
