package main

import (
  "fmt"
  "runtime"
)

var z int8
var a float64
// generally we use float64 for floating point variables

const l= 32
// another method
const (
  m= 34.32 // this is untyped constant
  n string = "James Bond"// this is a typed constant
)
//IOTA
const(
  e=iota
  f
  g
  h
)
const(
  i=iota
  j
  k
)
const(
  _ = iota
  kb=1<<(iota*10)
  mb=1<<(iota*10)
  gb=1<<(iota*10)
)


func main(){
  x:=7
  y:=25
  fmt.Println(x>y)
  z=23
  //z=-129 int8 allows int upto -128 to 127 only
  z=127
  // uint8 allows 0 to 255
  fmt.Println(z)
  fmt.Println(runtime.GOOS)// operating system
  fmt.Println(runtime.GOARCH)// architecture

  //STRING TYPE
  fmt.Println("STRING TYPE")
  s:=" hello pushkar "
  bs:=[]byte(s)// byte is an alias for uint8
               // rune is an alias for int32
  fmt.Println(bs)
  fmt.Printf("%T\n",bs)

  for i:=0; i<len(s); i++{
    fmt.Printf("%#U",s[i])// #U prints utf code for each character
  }
  fmt.Println("")
  for i, v:=range s{
    fmt.Println(i,v)
  }
  fmt.Println("")
  rawstring:= `this is raw
"but this is not"
string`
  fmt.Println(rawstring)

  // constant keyword
  fmt.Println("CONSTANT")
  fmt.Println(l)
  fmt.Println(m)
  fmt.Println(n)
  fmt.Printf("%T\n",l)
  fmt.Printf("%T\n",m)
  fmt.Printf("%T\n",n)

  //iota
  fmt.Println("")
  fmt.Println(e)
  fmt.Println(f)
  fmt.Println(g)
  fmt.Println(h)
  fmt.Println(i)
  fmt.Println(j)
  fmt.Println(k)

  fmt.Printf("%d\t\t\t%b\n",kb,kb)
  fmt.Printf("%d\t\t\t%b\n",mb,mb)
  fmt.Printf("%d\t\t%b\n",gb,gb)
}
