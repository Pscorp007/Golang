package main

import "fmt"

var b bool
var a int
var c string
var d bool
// these are package level variables and they are assigned zero values by def
// ault
type mytype int
var var1 mytype
var var2 int
func main(){
  x:=42
  y:="Its a string"
  z:= true
  b=true
  fmt.Println(x,y,z,b)
  fmt.Println(x)
  fmt.Printf("%v\n",b)

  fmt.Println(a)
  fmt.Println(c)
  fmt.Println(d)
  //
  s:=fmt.Sprintf("%v\t%v\t%v\n",x,y,z)
  fmt.Println(s)

  fmt.Printf("%v\n",var1)
  fmt.Printf("%T\n",var1)
  var1=42
  fmt.Println(x)
  var2=int(var1)
  fmt.Println(var2)
  fmt.Printf("%T",var2)


}
