package main

import "fmt"

type student struct{
  firstname string
  lastname string
  age int
}

type secretAgent struct{
  student
  ltk bool;
}

func main(){
  s1:= student{
    firstname : "Aman",
    lastname : "tripathi",
    age : 23, // again look out for the comma we put in the last too
  }
  fmt.Println(s1)
  fmt.Println(s1.firstname, s1.age)

  sa:= secretAgent{
    student: student{// here we have to give it explictly the type,which we didn't have to do for othe variables of struct
      firstname: "james",
      lastname: "bond",
      age : 25,
    },
    ltk : true,
  }
  fmt.Println(sa)
  fmt.Println(sa.student.firstname, sa.lastname, sa.age, sa.ltk)


  // CREATING AN ANONYMOUS STRUCT
  // this is called anonymous struct cuz this struct doesn't have a name or identifier

  s3:= struct{
    first string
    last string
    age int
  }{
    first : "pushkar",
    last : "soni",
    age : 18,
  }
  fmt.Println(s3)
  fmt.Println(s3.first,s3.last,s3.age)
}
