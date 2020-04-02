//Name: Vincent G.
//Date: 4/1/2020
//Description: How many eggs?

package main

import "fmt"

func dozen_Divider(){
  //declare variable as "dozens" type integer
  var dozens int
  //ask user to enter in a number
  fmt.Println("Enter a number you'd like to know how many dozens are inside of: ")
  fmt.Scanln(&dozens)
  fmt.Println()
  //divide the number by 12 and output the quotiant to show how many dozens are in the users number
  fmt.Println("There are",dozens/12,"dozens inside of",dozens)
}

func main() {
  //call the function to get
  dozen_Dvider()
}