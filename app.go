package main

import (
	"fmt"
	"time"
  "github.com/fatih/color"
  "greet/greet"
)

func main() {
	fmt.Println("Welcome to the playground!")
  red := color.New(color.FgRed).SprintFunc()
	fmt.Println("The time is", time.Now())
  name := greet.Hello("Alireza")
  fmt.Println(name)
  var username string
  fmt.Println(red("Enter Your Name:"))
  fmt.Scanf("%v", &username)
  fmt.Println("Hi", username)
}
