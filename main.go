//Name: Vincent G.
//Date: 4/1/2020
//Description: How many eggs?

package main

import "fmt"

func dozen(dozens int){
  //multiply the input by 12 to get the number of eggs in the number of dozens entered
  fmt.Println("There are",dozens*12,"eggs in",dozens,"dozens of eggs.")
}

func main() {
  //declare var i as integer
  var i int

  //ask user to enter in a number,scan for i, call dozen(i)
  fmt.Println("Enter in an amount of donzens to figure out how many eggs are in that many dozens.")

  fmt.Scanln(&i)

  dozen(i)
}
