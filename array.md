package main
 
import (
 "fmt"
)
 
func main() {
 var name [3][3] int
 
  for x := 0; x < len(name); x++ {
	  fmt.Println()
	  for y :=0; y < len(name); y++ {
 			fmt.Print("   ",x,",",y, " =", name[y][x])
	  }
 }
}
