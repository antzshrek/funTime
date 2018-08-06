/*

write a program that converts from Fahrenheit into Celsius.

*/


package main
import "fmt"

func main(){
	var F float64
	fmt.Println("Enter the value you'll wanna convert from Fahrenheit into Celsius: ")
	fmt.Scanf("%f", &F)
	C := (F - 32) * 5/9
	fmt.Println("The equivalent of the Number you enter is: ", C, "°C")
}
