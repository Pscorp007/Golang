package main

import "fmt"

func main(){
  m := map[string]int{// map[key type]valuetype{ key:value, key:value,}//note the comma in the end
    "Aditya":17,
    "Naresh": 19,
  }
  fmt.Println(m)
  fmt.Println(m["Aditya"])
  fmt.Println(m["Naresh"])
  fmt.Println(m["bunty"])// not bunty do not exist so zero value is provided by the compiler by default
  // thres a way to check if it exist or not
  v,okb:=m["bunty"]// v stores the value , okb tells if it exist or not
  fmt.Println(v)
  fmt.Println(okb)

  if value,present:=m["Aditya"]; present{
    fmt.Println("it is present",value)
  }

  m["bunty"]=24// now this includes bunty

  // printing all key value pairs
  fmt.Println("")
  for k,v := range m{
    fmt.Println(k, v)
  }
  // the above for loop technique can be used for slices too
  xi := []int{4,5,6,7,8}
  for i,v := range xi{
    fmt.Println(i, v)
  }

  // DELETION FROM A MAP
  delete(m,"bunty")// delete(map,key)
  delete(m,"ronit")// no error can delete non existing members too
  // your can also check and delet using above technique
}
