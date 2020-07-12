package main
import "fmt"

func main() {
  const (
    i int = iota
    j = iota
    )
  fmt.Println(i)
  fmt.Println(j)
  k := 5
  fmt.Println(k)

  // enum
  const (
    Sunday = iota;
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
    numberOfDays
  )
  fmt.Println(numberOfDays);
}

