package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
)

func main() {
  fmt.Println("What is the input string?")

  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  input := scanner.Text()

  inputSliced := strings.Split(input, "")

  fmt.Printf(input + " has %d characters.", len(inputSliced))
}
