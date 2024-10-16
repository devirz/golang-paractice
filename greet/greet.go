package greet

import "fmt"

func Hello(name string) string {
  msg := fmt.Sprintf("Hi %v!", name)
  return msg
}
