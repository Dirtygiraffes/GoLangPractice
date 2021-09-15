package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)

  fmt.Println("Enter a noun: ")
  scanner.Scan()
  noun := scanner.Text()

  fmt.Println("Enter a verb: ")
  scanner.Scan()
  verb := scanner.Text()

  fmt.Println("Enter an adjective: ")
  scanner.Scan()
  adjective := scanner.Text()

  fmt.Println("Enter an adverb: ")
  scanner.Scan()
  adverb := scanner.Text()

  fmt.Printf("Do you %s your %s %s %s? That's hilarious!", verb, adjective, noun, adverb)
}
