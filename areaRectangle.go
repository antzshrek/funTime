/*
write a program that calculate the area of a rectangle.
*/

package main
import "fmt"


func main(){
    var input1 float64
    fmt.Print("Enter the length of the rectangle: ")
    fmt.Scanf("%f", &input1)
    var input2 float64
    fmt.Print("Enter a width of the rectangle: ")
    fmt.Scanf("%f", &input2)
    area := input1 * input2
    fmt.Println("The area of the rectangle is: ", area)
}

