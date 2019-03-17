package main

import "fmt"

func main(){

  if x:=2; x==2{
    fmt.Println("this condition was true")
  }
 // fmt.Println(x)// this conditon gives error cuz scope of x is in the if condition only
  fmt.Println("this is a statement");/* semicolon*/ fmt.Println("this another statemet")
}
