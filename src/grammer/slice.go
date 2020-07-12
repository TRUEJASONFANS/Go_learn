package main

import "fmt"

func main() {
  var myArray [10] int = [10] int {1, 2, 3,4, 5,6 ,7 ,8, 9, 10}
  var mySlice []int = myArray[:5]
  mySlice[0] = 3;

  for _, v := range myArray {
    fmt.Print(v, "");
  }

  fmt.Println();

  for _,v := range mySlice {
    fmt.Print(v, "")
  }
  fmt.Println();

  mySlice = make([]int , 5, 10);
  mySlice[0] = 0;

  for _, v := range myArray {
    fmt.Print(v, "");
  }
  fmt.Println();

  fmt.Println("cap:",cap(mySlice));
  fmt.Println("len:",len(mySlice));
}
