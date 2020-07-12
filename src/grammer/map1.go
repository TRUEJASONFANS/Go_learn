package main

import "fmt"
type PersonInfo struct {
  ID string
  Name string
  Address string
}
func main() {
  var personDB map[string] PersonInfo;
  personDB = make(map[string] PersonInfo)
  personDB["12345"] = PersonInfo{"12345", "www","Room 203"}

  person, ok := personDB["12345"]
  if ok {
    fmt.Println("Found person", person.Name, person.Address,person.ID)
  } else {
    fmt.Println("Did not find person with ID 1234")
  }

  // delete
  delete(personDB, "12345")
  person, ok = personDB["12345"]
  if ok {
    fmt.Println("Found person", person.Name, person.Address,person.ID)
  } else {
    fmt.Println("Did not find person with ID 12345")
  }
}
