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
    var A,B float64
    fmt.Scan(&A)
    fmt.Scan(&B)
    MEDIA := ((A * 3.5) + (B * 7.5)) / 11
    fmt.Println(fmt.Sprintf("MEDIA = %.5f", MEDIA))

    //Problema 1006 - Average 2
    var A,B,C float64
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

    //Problema 1008 - Salario
    var horas,numero int
    var salario_hora,salario float64
    fmt.Scan(&numero)
    fmt.Scan(&horas)
    fmt.Scan(&salario_hora)
    salario = salario_hora * float64(horas)
    fmt.Println(fmt.Sprintf("NUMBER = %d", numero))
    fmt.Println(fmt.Sprintf("SALARY = U$ %.2f", salario))

    //Problema 1009 - Salary with Bonus
    var nome string
    var salario,total_vendido float64
    fmt.Scan(&nome)
    fmt.Scan(&salario)
    fmt.Scan(&total_vendido)
    fmt.Println(fmt.Sprintf("TOTAL = R$ %.2f", salario + (total_vendido / 100 * 15)))

    //Problema 1010 - Simple Calculate
    //IMPORTAR "strconv", "strings", "os" E "bufio"
    //Foi necessário usar bufio ao invés de Scan, pois o Scan não lê espaço na string
    reader := bufio.NewReader(os.Stdin)
    prod1, _ := reader.ReadString('\n') // o _ ignora a variavel
    prod1_partes := strings.Fields(strings.TrimSpace(prod1))
    prod2, _ := reader.ReadString('\n')
    prod2_partes := strings.Fields(strings.TrimSpace(prod2))
    
    prod1_unid, err1 := strconv.Atoi(prod1_partes[1])
    prod1_valor, err2 := strconv.ParseFloat(prod1_partes[2], 64)
    prod2_unid, err3 := strconv.Atoi(prod2_partes[1])
    prod2_valor, err4 := strconv.ParseFloat(prod2_partes[2], 64)
    
    if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
        fmt.Println("erro")
        return
    }
    total := (prod1_valor * float64(prod1_unid)) + (prod2_valor * float64(prod2_unid))

    fmt.Println(fmt.Sprintf("VALOR A PAGAR: R$ %.2f", total))

}