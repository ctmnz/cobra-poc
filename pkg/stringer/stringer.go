package stringer

import (
  "strconv"
  "fmt"
)

func Reverse(input string) (result string) {
  for _, c := range input {
    result = string(c) + result
  }
  return result
}

func Inspect(input string, digits bool) (count int, kind string) {
  if !digits {
    return len(input), "char"
  }
  return inspectNumbers(input), "digit"
}

func inspectNumbers(input string) (count int) {
  for _, c := range input {
    _, err := strconv.Atoi(string(c))
    if err == nil {
      count++
    }
  }
  return count
}

func SayHello(input string) (result string) {
//  r := fmt.Sprintf("Hello, %s", input)
  return fmt.Sprintf("Hello, %s!", input)
}


