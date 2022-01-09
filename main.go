package main

import (
  "fmt"
  //"github.com/rs/zerolog/log"
)


// interface define some functionalities
type speaker interface {
  speak()
}

type person struct {
  name string
  email string
}

type cat struct {
  name string
}

func (p person) speak() {
  fmt.Println(p.name)
}

func saysomething(h speaker){
  h.speak()
}

func (c cat) speak (){
  fmt.Println("meow")
}

func main()  {

  p := person{name:"Jack", email:"s@g.com"}
  saysomething(p)

  c := cat{}
  saysomething(c)
  // fmt.Println("Hello World...")
  // log.Info().Msg("Hi")
}
