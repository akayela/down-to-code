package main

import (
  "fmt"
  "magazine"
)

func main() {
  address := magazine.Address{Street: "123 Oak St", City: "Omaha", State: "NE", PostalCode: "68111"}
  var s magazine.Subscriber
  s.Address = address
  s.Rate = 4.99
  fmt.Println(s.Rate)
  fmt.Println(s.Address)
  
  subscriber2 := magazine.Subscriber{Name: "Audience Kayela", Rate: 4.99, Active: true}
  subscriber2.City = "Chiredzi"
  subscriber2.Street = "655 Kayela St"
  subscriber2.State = "Masv"
  subscriber2.PostalCode = "023"
  fmt.Println(subscriber2.Name)
  fmt.Println(subscriber2.Rate)
  fmt.Println(subscriber2.Active)
  fmt.Println(subscriber2.Street)
  fmt.Println(subscriber2.City)
  fmt.Println(subscriber2.State)
  fmt.Println(subscriber2.PostalCode)
}
