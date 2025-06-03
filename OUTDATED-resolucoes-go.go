package main
 
import (
    "fmt"
)
 
// PRIMEIRA VERSÃO DE ALGUMAS RESOLUÇÕES, COM CÓDIGO MAIS SIMPLES E SEM FUNÇÕES

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

    //Problema 1011 - Sphere
    const pi = 3.14159
    var R float64
    fmt.Scan(&R)
    fmt.Println(fmt.Sprintf("VOLUME = %.3f", (4.0 / 3.0) * pi * (R*R*R)))

    //Problema 1012 - Area
        //IMPORTAR "strconv", "strings", "os" E "bufio"
    reader := bufio.NewReader(os.Stdin)
    entrada, _ := reader.ReadString('\n')
    entrada_lista := strings.Fields(strings.TrimSpace(entrada))
    
    A, _ := strconv.ParseFloat(entrada_lista[0], 64)
    B, _ := strconv.ParseFloat(entrada_lista[1], 64)
    C, _ := strconv.ParseFloat(entrada_lista[2], 64)
    
    fmt.Println(fmt.Sprintf("TRIANGULO: %.3f", A * C / 2))
    fmt.Println(fmt.Sprintf("CIRCULO: %.3f", 3.14159 * (C*C)))
    fmt.Println(fmt.Sprintf("TRAPEZIO: %.3f", 0.5 * (C * (A+B))))
    fmt.Println(fmt.Sprintf("QUADRADO: %.3f", B*B))
    fmt.Println(fmt.Sprintf("RETANGULO: %.3f", A*B))

    //Problema 1013 - The Greatest
        //IMPORTAR "strconv", "strings", "os", "bufio" E "math"
    reader := bufio.NewReader(os.Stdin)
    entrada, _ := reader.ReadString('\n')
    entrada_lista := strings.Fields(strings.TrimSpace(entrada))
    
    A, _ := strconv.Atoi(entrada_lista[0])
    B, _ := strconv.Atoi(entrada_lista[1])
    C, _ := strconv.Atoi(entrada_lista[2])
    
    AB := float64(A) - float64(B)
    MaiorAB := (A + B + int(math.Abs(AB))) / 2
    ABC := float64(MaiorAB) - float64(C)
    MaiorABC := (MaiorAB + C + int(math.Abs(ABC))) / 2
    fmt.Println(fmt.Sprintf("%d eh o maior", MaiorABC))

    //Problema 1014 - Consumption
    var distancia int
    var gasolina_gasta float64
    fmt.Scan(&distancia)
    fmt.Scan(&gasolina_gasta)
    fmt.Println(fmt.Sprintf("%.3f km/l", float64(distancia) / gasolina_gasta))

    //Problema 1015 - Distance Between Two Points
        //RESOLUÇÃO 1 - usando interface (lista de diferentes tipos)
        //IMPORTAR "strconv", "strings", "os", "bufio" E "math"
    reader := bufio.NewReader(os.Stdin)
    entrada1, _ := reader.ReadString('\n')
    p1_strings := strings.Fields(strings.TrimSpace(entrada1))
    entrada2, _ := reader.ReadString('\n')
    p2_strings := strings.Fields(strings.TrimSpace(entrada2))
    
    p1 := make([]interface{}, 2)
    p1[0], _ = strconv.ParseFloat(p1_strings[0], 64)
    p1[1], _ = strconv.ParseFloat(p1_strings[1], 64)
    p2 := make([]interface{}, 2)
    p2[0], _ = strconv.ParseFloat(p2_strings[0], 64)
    p2[1], _ = strconv.ParseFloat(p2_strings[1], 64)
    
    distance := math.Sqrt(math.Pow(p2[0].(float64) - p1[0].(float64), 2) + math.Pow(p2[1].(float64) - p1[1].(float64), 2))
    fmt.Println(fmt.Sprintf("%.4f", distance))

        //RESOLUÇÃO 2 - usando struct
        //IMPORTAR "strconv", "strings", "os", "bufio" E "math"
        //Criar struct entre função main e imports
    type P struct {
        x float64
        y float64
    }

    reader := bufio.NewReader(os.Stdin)
    
    entrada1, _ := reader.ReadString('\n')
    p1_strings := strings.Fields(strings.TrimSpace(entrada1))
    p1_x, _ := strconv.ParseFloat(p1_strings[0], 64)
    p1_y, _ := strconv.ParseFloat(p1_strings[1], 64)
    p1 := P{
        x:  p1_x,
        y:  p1_y,
    }
    
    entrada2, _ := reader.ReadString('\n')
    p2_strings := strings.Fields(strings.TrimSpace(entrada2))
    p2_x, _ := strconv.ParseFloat(p2_strings[0], 64)
    p2_y, _ := strconv.ParseFloat(p2_strings[1], 64)
    p2 := P{
        x:  p2_x,
        y:  p2_y,
    }
    
    distance := math.Sqrt(math.Pow(p2.x - p1.x, 2) + math.Pow(p2.y - p1.y, 2))
    fmt.Println(fmt.Sprintf("%.4f", distance))

    //Problema 1016 - Distance
    var distance int
    fmt.Scan(&distance)
    fmt.Println(fmt.Sprintf("%d minutos", distance * 2))

    //Problema 1017 - Fuel Spent
    var velocidade,horas int
    fmt.Scan(&horas)
    fmt.Scan(&velocidade)
    consumo := float64(horas * velocidade) / 12.0
    fmt.Println(fmt.Sprintf("%.3f", consumo))

    //Problema 1018 - Banknotes
    var entrada,notas int
    fmt.Scan(&entrada)
    fmt.Println(entrada)
    notas = int(entrada / 100)
    fmt.Println(fmt.Sprintf("%d nota(s) de R$ 100,00", notas))
    entrada = entrada - (notas * 100)
    notas = int(entrada / 50)
    fmt.Println(fmt.Sprintf("%d nota(s) de R$ 50,00", notas))
    entrada = entrada - (notas * 50)
    notas = int(entrada / 20)
    fmt.Println(fmt.Sprintf("%d nota(s) de R$ 20,00", notas))
    entrada = entrada - (notas * 20)
    notas = int(entrada / 10)
    fmt.Println(fmt.Sprintf("%d nota(s) de R$ 10,00", notas))
    entrada = entrada - (notas * 10)
    notas = int(entrada / 5)
    fmt.Println(fmt.Sprintf("%d nota(s) de R$ 5,00", notas))
    entrada = entrada - (notas * 5)
    notas = int(entrada / 2)
    fmt.Println(fmt.Sprintf("%d nota(s) de R$ 2,00", notas))
    entrada = entrada - (notas * 2)
    fmt.Println(fmt.Sprintf("%d nota(s) de R$ 1,00", entrada))

    //Problema 1019 - Time Conversion
    var n int
    fmt.Scan(&n)
    horas := int(n / 3600)
    minutos := int((n - (horas * 3600)) / 60)
    segundos := int(n - (horas * 3600) - (minutos * 60))
    fmt.Println(fmt.Sprintf("%d:%d:%d", horas, minutos, segundos))

    //Problema 1020 - Age in Days
    var dias int
    fmt.Scan(&dias)
    anos := int(dias / 365)
    meses := int((dias - anos * 365) / 30)
    dias = dias - anos * 365 - meses * 30
    fmt.Println(fmt.Sprintf("%d ano(s)", anos))
    fmt.Println(fmt.Sprintf("%d mes(es)", meses))
    fmt.Println(fmt.Sprintf("%d dia(s)", dias))

}