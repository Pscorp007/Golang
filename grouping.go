package main

import "fmt"

func main(){
  // x := type {values} COMPOSITE LITEREAL
  // SLICE allows you to group together value of same type
  x := []int{1,2,3,4,5,6}
  fmt.Println(x)

  for i,e:=range x{// i is defined as index and e is all the values in range of x
    fmt.Println(i,e)
  }
  // SLICING A SLICE
  fmt.Println(x[:])// colon gives you ability to SLICE
  fmt.Println(x[1:])// indexing start from 0 to n-1
  fmt.Println(x[1:3])// 3 wont print value at 3 it will print values before 3 only

  // APPEND TO A SLICE
  fmt.Println("append to a slice")
  y:=[]int{8,9,10}
  x=append(x,y[0])// you have to assign it to x to see the change
  z:=append(x,y[1])// cuz assigning to z jus creates a new variable z with an appende value
  fmt.Println(x)
  fmt.Println(z)
  x=append(x,y...)// "..." used to denote all values
  fmt.Println(x)
  //DELETING FROM A SLICE
  x=append(x[:2],x[4:]...)//deleting things from pos 2 to 3 both included
  fmt.Println(x)
  // MAKE
  fmt.Println("use of make")
  k:= make([]int,10,12)// make(type,size, capacity)
  fmt.Println(k)
  fmt.Println(len(k))
  fmt.Println(cap(k))
  k[0]=67
  k[9]=999
  // k[10]=23 this is not allowed
  k=append(k,345)// you can see that the length increased
  fmt.Println(k)
  fmt.Println(len(k))
  fmt.Println(cap(k))
  k=append(k,750)
  fmt.Println(k)
  fmt.Println(len(k))
  fmt.Println(cap(k))
  k=append(k,990)// note that capacity has also increased infact doubled.
  // see in line no 22 and 23 there's a problem that everytime you insert values
  // you create new slice and store it in itself this takes lot of processing.
  // having predefined capacity saves this processing
  fmt.Println(k)
  fmt.Println(len(k))
  fmt.Println(cap(k))

  // multidimensional SLICE
  fmt.Println("multidimensional slice")
  str1:=[]string{"ronit","vannila"}
  str2:=[]string{"bunty","umbrella"}
  fmt.Println(str1)
  fmt.Println(str2)
  str:=[][]string{str1,str2}
  fmt.Println(str)

}
