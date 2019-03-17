package main

import "fmt"

var k=26
//k:=25  you cannot do this this is the differnce  between var and short dec operator
var l int
//declare a variable of type int
var s string = "shaken not stirred"
// this is a static programming language
// a variable is declared to store values of certain type only
// this is not a dynamic programming language
var name string
// this declares a variable of string type "" called zero value
var age= 42

//Creating your own type
var a int =76
type hotdog int
var b hotdog
func main() {
	fmt.Println("hello follow me on twitter")
  foo()
  fmt.Println("print sometheing",23,true)
  n,e := fmt.Println("my name",24,false)
  fmt.Println(n)
  fmt.Println(e)
  m,_:=fmt.Println("prit")
  fmt.Println(m)

  x:=25     // declaration(x:) + assignment(=) // called short declaration operator
  // x has local scope only
  //short declaration operator can only be used for creating new variables
  // you cannot use short declaration operator on already declared variables
  
  fmt.Println(x)
  x=27     // assignment
  fmt.Println(x)
  z:="james,bond"
  fmt.Println(z)
  var y=36
  fmt.Println(y)
  fmt.Println(k)
  fmt.Println(s)
  fmt.Println("the value of y is",name,"yes")
  fmt.Println(age)

  // PRINTING
  fmt.Printf("type: %T\n",age)
  fmt.Printf("binary : %b\n",age)
  fmt.Printf("base 2 : %x\n",age)
  fmt.Print("vase ")// diff b/w println and print is pln gives a line
  fmt.Printf("hexadecimal : %#x\n",age)
  // sprint
  s:= fmt.Sprintf("%T\t%b\t%x\t%#x\n ",age,age,age,age)
  // we assigned above output to a string
  fmt.Println(s)
  fmt.Printf("%v\n",age) // most important prints value in the default format
  //fmt.Fprint() printing to a file

  //CREATING YOUR OWN type
  b=43
  fmt.Println(a)
  fmt.Println(b)
  fmt.Printf("%T\n",b)
  //a=b not allowed as b is of type hotdog
  a=int(b)// allowed and this is called conversion not casting
  // also not that brackets are put around b not around int as in c++
  fmt.Println(a)
}

func foo(){
  fmt.Println("i am in foo")
}
