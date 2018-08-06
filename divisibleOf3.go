/*
Write a program that prints out all the numbers between 1 and 100 that are evenly divisible by 3 (i.e., 3, 6, 9, etc.).

*/

package main
import "fmt"


func main() {
	for j := 1; j <= 100; j++ {
		if j % 3 == 0{
			fmt.Println(j)
		}else{
			fmt.Println(j, " is not evenly divisible by 3 ")
		}
	}
   }
