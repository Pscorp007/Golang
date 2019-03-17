package main
import "fmt"

const(
  m = 23
  n float64 = 45.3
)

const(
  _=iota
  g
  h
  i
  j
)

func main(){
  a:=45
  fmt.Printf("%d %b %#x\n",a,a,a)
  b:=25
  c:= (a<b)
  fmt.Println(c)
  c= (a<=b)
  fmt.Println(c)
  c= (a>b)
  fmt.Println(c)
  c= (a>=b)
  fmt.Println(c)
  c= (a==b)
  fmt.Println(c)
  c= (a!=b)
  fmt.Println(c)

  fmt.Println(m)
  fmt.Println(n)

  d:=45
  fmt.Printf("%d\t%b\t%#x\n",d,d,d)
  e:=d<<2
  fmt.Printf("%d\t%b\t%#x\n",e,e,e)

  rawstring:= `this is raw
" but this is not"
string`
  fmt.Println(rawstring)

  fmt.Println(g,h,i,j)
}
