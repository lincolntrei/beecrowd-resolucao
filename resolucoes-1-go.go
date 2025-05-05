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
    fmt.Println(fmt.Sprintf("A=%.4f", A))

    //Problema 1003 - Simple Sum
    var A,B int
    fmt.Scan(&A)
    fmt.Scan(&B)
    SOMA := A + B
    fmt.Println("SOMA =", SOMA)

    //Problema 1004 - Simple Product
    var A,B int
    fmt.Scan(&A)
    fmt.Scan(&B)
    PROD := A * B
    fmt.Println("PROD =", PROD)

    //Problema 1005 - Average 1
    var A,B float32
    fmt.Scan(&A)
    fmt.Scan(&B)
    MEDIA := ((A * 3.5) + (B * 7.5)) / 11
    fmt.Println(fmt.Sprintf("MEDIA = %.5f", MEDIA))

    //Problema 1006 - Average 2
    var A,B,C float32
    fmt.Scan(&A)
    fmt.Scan(&B)
    fmt.Scan(&C)
    MEDIA := ((A * 2) + (B * 3) + (C * 5)) / 10
    fmt.Println(fmt.Sprintf("MEDIA = %.1f", MEDIA))

    //Problema 1007 - Difference
    var A,B,C,D int
    fmt.Scan(&A)
    fmt.Scan(&B)
    fmt.Scan(&C)
    fmt.Scan(&D)
    DIFERENCA := A * B - C * D
    fmt.Println(fmt.Sprintf("DIFERENCA = %d", DIFERENCA))

}