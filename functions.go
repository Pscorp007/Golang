package main

import "fmt"

func main(){
  s1:= createstring("pushkar")
  fmt.Println(s1)
  //IN GOLANG WE CAN HAVE MULTIPLE RETURN FROM FUNCITON
  a,b:= tango("basket","ball","lnmiit")
  fmt.Println(a)
  fmt.Println(b)
  foo("james",1,2,3,4,5,6,7)
  xi:= []int{3,4,5}
  foo("james",xi...)// this infinite dot is necessary for passing slice
  // this ... is called unfuling we always have to unfurl the slice to pass it to func

  defer foo1()//defer keyword
  // the function written after defer is always executed after exit of container function
  // here in this case foo1 will be executed after main ends
  bar()

  // METHOD FUNCTION
  sa1:=secretagent{
    student : student{
      f_name: "pushkar",
      l_name: "soni",
    },
    ltk: true,
  }

  sa1.methodfunction()

  //ANONYMOUS FUNCTION
  // func (parameters){ code }(arguments)
  func (){
    fmt.Println("anonymous function 1 ran")
  }()
  func (t int){
    fmt.Println("anony function with value: ",t)
  }(45)
  // FUNCTION EXPRESSION i.e,. assigning function to a variable
  f:= func (m int){
    fmt.Println("function expression with value : ",m)
  }
  f(1920)

  // FUNCTION RETURNING A FUNCTION

  i:=bar2()// bar 2 will return a function
  j:=i() // try to remove this line then see what happens
  fmt.Println(j)
  // another way to do above three lines is fmt.Println(bar2()())

  // CLOSURE
  // note the scope of x
  g:= incrementor()
  h:= incrementor()
  fmt.Println(g())
  fmt.Println(g())
  fmt.Println(g())
  fmt.Println(h())
  fmt.Println(h())

}




func createstring( s string ) string{ // func function_name( parameters) type_of_func{ }
  return fmt.Sprint("hello ", s)
}

func tango(s string, s1 string, s2 string) (string,bool){// returning multiple parameters
  a:= fmt.Sprint(s2," is playing ", s, s1)
  b:= true
  return a,b// syntax for return no brackets
}

func foo(s string, x ...int){// PASSING INFINITE INTS may be zero
  // NOTE: when you pass infinite type parameters then it should be the last parameters
  // i.e,. foo(x ...int,s string)-> not allowed
  // but   foo(s string, x ...int)-> is allowed
  fmt.Println(s)
  sum:=0
  for i,v:= range x{// here  i is necessary so if you don't want to use is then write for _,v:=range x{}
    fmt.Printf("%v-%v,",i,v)
    sum+=v;
  }
  fmt.Println(sum)
}

func foo1(){
  fmt.Println("I am foo1")
}
func bar(){
  fmt.Println("I am bar")
}

type student struct{
  f_name string
  l_name string
}
type secretagent struct{
  student
  ltk bool
}

//  METHODS
// syntax for method functions
// func (r reciever) identifier(parameters) (return statements){ code }

func (s secretagent) methodfunction(){
  fmt.Println(s)
}

func bar2() func() int{// function bar2 whose return type is a function of return int
    return func() int{
        return 420
      }
  }

func incrementor() func() int{
  var x int // x initialized to zero
    return func () int{
      x++
      return x
    }
  }
