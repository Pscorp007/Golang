package main
import "fmt"

type person struct{
  first_name string
  last_name string
  fav_icecream_flavour []string
}

type vehicle struct{
  door int
  color string
}
type truck struct{
  vehicle
  fourWheel bool
}
type sedan struct{
  vehicle
  luxury bool
}

func main(){
  p1:= person{
    first_name : "pushkar",
    last_name : "soni",
    fav_icecream_flavour : []string{"vanilla","chocolate","walnuts"},
  }
  p2:= person{
    first_name : "bunty",
    last_name : "tekwani",
    fav_icecream_flavour : []string{"american nuts","strawberry","rainberry"},
  }
  // fmt.Println(p1)
  fmt.Println(p1.first_name)
  fmt.Println(p1.last_name)
  for i,e:=range p1.fav_icecream_flavour{
    fmt.Println(i+1,e)
  }


  m:= map[string]person{
    p1.last_name : p1,
    p2.last_name : p2,
  }

  for k,v:= range m{
    fmt.Println(v.first_name)
    fmt.Println(k)
    for i,e:= range v.fav_icecream_flavour{
      fmt.Println(i+1,e)
    }
  }

  t1:= truck{
    vehicle : vehicle{
      door : 6,
      color : "blue",
    },
    fourWheel : true,
  }

  s1:= sedan{
    vehicle : vehicle{
      door : 4,
      color : "black",
    },
    luxury : true,
  }

  fmt.Println(t1)
  fmt.Println(s1)
  fmt.Println(t1.door)
  fmt.Println(s1.door)

  as:= struct{
    door int
    color string
  }{
    door : 3,
    color : "red",
  }
  fmt.Println(as)
  fmt.Println(as.color)

}
