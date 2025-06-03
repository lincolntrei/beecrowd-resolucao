package main

import (
    "fmt"
    "bufio"
    "math"
    "os"
    "strconv"
    "strings"
)

func main() {
    //Insira aqui a função a ser executada
}

// Problema 1000 - Hello World!
func resolucao1000() {
    fmt.Println("Hello World!")
}

// Problema 1001 - Extremaly Basic
func resolucao1001() {
    var a, b int
    fmt.Scan(&a, &b)
    fmt.Println("X =", a+b)
}

// Problema 1002 - Area of a Circle
func resolucao1002() {
    var R float64
    const n = 3.14159
    fmt.Scanln(&R)
    A := n * (R * R)
    fmt.Println(fmt.Sprintf("A=%.4f", A))
}

// Problema 1003 - Simple Sum
func resolucao1003() {
    var A, B int
    fmt.Scan(&A, &B)
    fmt.Println("SOMA =", A+B)
}

// Problema 1004 - Simple Product
func resolucao1004() {
    var A, B int
    fmt.Scan(&A, &B)
    fmt.Println("PROD =", A*B)
}

// Problema 1005 - Average 1
func resolucao1005() {
    var A, B float64
    fmt.Scan(&A, &B)
    MEDIA := ((A * 3.5) + (B * 7.5)) / 11
    fmt.Println(fmt.Sprintf("MEDIA = %.5f", MEDIA))
}

// Problema 1006 - Average 2
func resolucao1006() {
    var A, B, C float64
    fmt.Scan(&A, &B, &C)
    MEDIA := ((A * 2) + (B * 3) + (C * 5)) / 10
    fmt.Println(fmt.Sprintf("MEDIA = %.1f", MEDIA))
}

// Problema 1007 - Difference
func resolucao1007() {
    var A, B, C, D int
    fmt.Scan(&A, &B, &C, &D)
    fmt.Println(fmt.Sprintf("DIFERENCA = %d", A*B-C*D))
}

// Problema 1008 - Salario
func resolucao1008() {
    var horas, numero int
    var salario_hora float64
    fmt.Scan(&numero, &horas, &salario_hora)
    salario := salario_hora * float64(horas)
    fmt.Println(fmt.Sprintf("NUMBER = %d", numero))
    fmt.Println(fmt.Sprintf("SALARY = U$ %.2f", salario))
}

// Problema 1009 - Salary with Bonus
func resolucao1009() {
    var nome string
    var salario, total_vendido float64
    fmt.Scan(&nome, &salario, &total_vendido)
    fmt.Println(fmt.Sprintf("TOTAL = R$ %.2f", salario+(total_vendido*0.15)))
}

// Problema 1010 - Simple Calculate
//Foi necessário usar bufio ao invés de Scan, pois o Scan não lê espaço na string
func resolucao1010() {
    reader := bufio.NewReader(os.Stdin)
    prod1, _ := reader.ReadString('\n')
    prod1_partes := strings.Fields(strings.TrimSpace(prod1))
    prod2, _ := reader.ReadString('\n')
    prod2_partes := strings.Fields(strings.TrimSpace(prod2))

    prod1_unid, _ := strconv.Atoi(prod1_partes[1])
    prod1_valor, _ := strconv.ParseFloat(prod1_partes[2], 64)
    prod2_unid, _ := strconv.Atoi(prod2_partes[1])
    prod2_valor, _ := strconv.ParseFloat(prod2_partes[2], 64)

    total := (prod1_valor * float64(prod1_unid)) + (prod2_valor * float64(prod2_unid))
    fmt.Println(fmt.Sprintf("VALOR A PAGAR: R$ %.2f", total))
}

// Problema 1011 - Sphere
func resolucao1011() {
    const pi = 3.14159
    var R float64
    fmt.Scan(&R)
    fmt.Println(fmt.Sprintf("VOLUME = %.3f", (4.0/3.0)*pi*(R*R*R)))
}

// Problema 1012 - Area
func resolucao1012() {
    reader := bufio.NewReader(os.Stdin)
    entrada, _ := reader.ReadString('\n')
    entrada_lista := strings.Fields(strings.TrimSpace(entrada))

    A, _ := strconv.ParseFloat(entrada_lista[0], 64)
    B, _ := strconv.ParseFloat(entrada_lista[1], 64)
    C, _ := strconv.ParseFloat(entrada_lista[2], 64)

    fmt.Println(fmt.Sprintf("TRIANGULO: %.3f", A*C/2))
    fmt.Println(fmt.Sprintf("CIRCULO: %.3f", 3.14159*C*C))
    fmt.Println(fmt.Sprintf("TRAPEZIO: %.3f", 0.5*C*(A+B)))
    fmt.Println(fmt.Sprintf("QUADRADO: %.3f", B*B))
    fmt.Println(fmt.Sprintf("RETANGULO: %.3f", A*B))
}

// Problema 1013 - The Greatest
func resolucao1013() {
    reader := bufio.NewReader(os.Stdin)
    entrada, _ := reader.ReadString('\n')
    entrada_lista := strings.Fields(strings.TrimSpace(entrada))

    A, _ := strconv.Atoi(entrada_lista[0])
    B, _ := strconv.Atoi(entrada_lista[1])
    C, _ := strconv.Atoi(entrada_lista[2])

    MaiorAB := (A + B + int(math.Abs(float64(A-B)))) / 2
    MaiorABC := (MaiorAB + C + int(math.Abs(float64(MaiorAB-C)))) / 2
    fmt.Println(fmt.Sprintf("%d eh o maior", MaiorABC))
}

// Problema 1014 - Consumption
func resolucao1014() {
    var distancia int
    var gasolina_gasta float64
    fmt.Scan(&distancia, &gasolina_gasta)
    fmt.Println(fmt.Sprintf("%.3f km/l", float64(distancia)/gasolina_gasta))
}

// Problema 1015 - Distance Between Two Points (resolução 1)
func resolucao1015_resolucao1() {
    reader := bufio.NewReader(os.Stdin)
    entrada1, _ := reader.ReadString('\n')
    p1_strings := strings.Fields(strings.TrimSpace(entrada1))
    entrada2, _ := reader.ReadString('\n')
    p2_strings := strings.Fields(strings.TrimSpace(entrada2))

    p1x, _ := strconv.ParseFloat(p1_strings[0], 64)
    p1y, _ := strconv.ParseFloat(p1_strings[1], 64)
    p2x, _ := strconv.ParseFloat(p2_strings[0], 64)
    p2y, _ := strconv.ParseFloat(p2_strings[1], 64)

    distance := math.Sqrt(math.Pow(p2x-p1x, 2) + math.Pow(p2y-p1y, 2))
    fmt.Println(fmt.Sprintf("%.4f", distance))
}

// Problema 1015 - Distance Between Two Points (resolução 2)
type P struct {
    x float64
    y float64
}

func resolucao1015_resolucao2() {
    reader := bufio.NewReader(os.Stdin)

    entrada1, _ := reader.ReadString('\n')
    p1_strings := strings.Fields(strings.TrimSpace(entrada1))
    p1_x, _ := strconv.ParseFloat(p1_strings[0], 64)
    p1_y, _ := strconv.ParseFloat(p1_strings[1], 64)
    p1 := P{x: p1_x, y: p1_y}

    entrada2, _ := reader.ReadString('\n')
    p2_strings := strings.Fields(strings.TrimSpace(entrada2))
    p2_x, _ := strconv.ParseFloat(p2_strings[0], 64)
    p2_y, _ := strconv.ParseFloat(p2_strings[1], 64)
    p2 := P{x: p2_x, y: p2_y}

    distance := math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
    fmt.Println(fmt.Sprintf("%.4f", distance))
}

// Problema 1016 - Distance
func resolucao1016() {
    var distance int
    fmt.Scan(&distance)
    fmt.Println(fmt.Sprintf("%d minutos", distance*2))
}

// Problema 1017 - Fuel Spent
func resolucao1017() {
    var velocidade, horas int
    fmt.Scan(&horas, &velocidade)
    consumo := float64(horas*velocidade) / 12.0
    fmt.Println(fmt.Sprintf("%.3f", consumo))
}

// Problema 1018 - Banknotes
func resolucao1018() {
    var entrada int
    fmt.Scan(&entrada)
    fmt.Println(entrada)
    notas := []int{100, 50, 20, 10, 5, 2, 1}
    for _, nota := range notas {
        qtd := entrada / nota
        fmt.Println(fmt.Sprintf("%d nota(s) de R$ %d,00", qtd, nota))
        entrada %= nota
    }
}

// Problema 1019 - Time Conversion
func resolucao1019() {
    var n int
    fmt.Scan(&n)
    horas := n / 3600
    minutos := (n % 3600) / 60
    segundos := n % 60
    fmt.Println(fmt.Sprintf("%d:%d:%d", horas, minutos, segundos))
}

// Problema 1020 - Age in Days
func resolucao1020() {
    var dias int
    fmt.Scan(&dias)
    anos := dias / 365
    meses := (dias % 365) / 30
    dias = dias % 365 % 30
    fmt.Println(fmt.Sprintf("%d ano(s)", anos))
    fmt.Println(fmt.Sprintf("%d mes(es)", meses))
    fmt.Println(fmt.Sprintf("%d dia(s)", dias))
}

// Problema 1021 - Banknotes and Coins
func resolucao1021() {
    var entrada float64
    fmt.Scan(&entrada)
    valor := int(entrada * 100)
    fmt.Println("NOTAS:")
    notas := []int{10000, 5000, 2000, 1000, 500, 200}
    for _, nota := range notas {
        qtd := int(valor) / nota
        fmt.Printf("%d nota(s) de R$ %d.00\n", qtd, nota / 100)
        valor %= nota
    }
    fmt.Println("MOEDAS:")
    moedas := []int{100, 50, 25, 10, 05, 01}
    for _, moeda := range moedas {
        qtd := valor / moeda
        fmt.Printf("%d moeda(s) de R$ %.2f\n", qtd, float64(moeda) / 100)
        valor %= moeda
    }
}