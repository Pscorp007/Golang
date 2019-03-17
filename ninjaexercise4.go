package main
import "fmt"

func main(){
  x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }
  fmt.Println(x)
	fmt.Printf("%T\n", x)
  fmt.Println(x[:5])
  fmt.Println(x[5:])
  fmt.Println(x[2:7])
  fmt.Println(x[1:6])

  x= append(x,52)
  fmt.Println(x)
  x= append(x,52,53,54)
  y:=[]int{56,57,58,59,60}
  x=append(x,y...)
  fmt.Println(x)
  x = append(x[:10])
  fmt.Println(x)
  x= append(x[:3],x[6:]...)
  fmt.Println(x)

  s:= make([]string,0,100)
  s= append(s,` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
  fmt.Println(s)
  fmt.Println(len(s))
  fmt.Println(cap(s))

  sl:= [][]string{{"James", "Bond", "Shaken, not stirred"},{"Miss", "Moneypenny", "Helloooooo, James"}}
  fmt.Println(sl)

  for i,e:= range sl{
    fmt.Println("record: ",i)
    for _,f:= range e{
      fmt.Println(f)
    }
  }

  mp:= map[string][]string{
    "bond_james" : []string{"shaken,not stirred","martinis","Women"},
    "moneypenny_miss" : []string{"james bond","literature","computer science"},
    "no_dr" : []string{"being evil","ice cream","sunsets"},
  }
  // fmt.Println(mp)
  for k,e:= range mp{
    fmt.Println(k)
    for _,f:=range e{
      fmt.Println("\t",f)
    }
  }
  fmt.Println()
  mp["pushkar_soni"]=[]string{"coding","nan","creamroll"}
  for k,e:= range mp{
    fmt.Println(k)
    for _,f:=range e{
      fmt.Println("\t",f)
    }
  }

  fmt.Println()
  delete(mp,"pushkar_soni")
  for k,e:= range mp{
    fmt.Println(k)
    for _,f:=range e{
      fmt.Println("\t",f)
    }
  }
}
