package main
 
import (
    "fmt"
)
 
func main() {

    //Problema 1000 - Hello World!
	fmt.Println("Hello World!")

	//Problema 1001 - Extremaly Basic
	var a,b int
    fmt.Scan(&a)
    fmt.Scan(&b)
    x := a + b
    fmt.Println("X =", x)

	//Problema 1002 - Area of a Circle
	var R float64
    const n = 3.14159
    fmt.Scanln(&R)
    A := n * (R * R)
    fmt.Printf("A=%.4f", A)

}