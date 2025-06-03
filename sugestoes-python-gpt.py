"""Sugestões do GPT de melhoreas para cada resolução sem alterar a saída"""

def problema1000():
    #Problema 1000 - Hello World
    print("Hello World!")

def problema1001():
    #Problema 1001 - Extremely Basic
    A = int(input())
    B = int(input())
    print("X = " + str(A + B))
    # Melhoria: nada alterado, código já era direto.

def problema1002():
    #Problema 1002 - Área of the Circle
    raio = float(input())
    area = 3.14159 * raio * raio
    print(f"A={area:.4f}")
    # Melhoria: uso de f-string para formatação mais clara.

def problema1003():
    #Problema 1003 - Simple Sum
    A = int(input())
    B = int(input())
    print(f"SOMA = {A + B}")
    # Melhoria: uso de f-string, evitando string + int.

def problema1004():
    #Problema 1004 - Simple Product
    A = int(input())
    B = int(input())
    print("PROD =", A * B)
    # Melhoria: clareza mantida, código já simples.

def problema1005():
    #Problema 1005 - Average 1
    a = float(input())
    b = float(input())
    media = ((a * 3.5) + (b * 7.5)) / 11
    print(f"MEDIA = {media:.5f}")
    # Melhoria: cálculo direto e f-string.

def problema1006():
    #Problema 1006 - Average 2
    A = float(input())
    B = float(input())
    C = float(input())
    media = ((A * 2) + (B * 3) + (C * 5)) / 10
    print(f"MEDIA = {media:.1f}")
    # Melhoria: mesma lógica da anterior.

def problema1007():
    #Problema 1007 - Difference
    A = int(input())
    B = int(input())
    C = int(input())
    D = int(input())
    print(f"DIFERENCA = {A * B - C * D}")
    # Melhoria: fórmula simplificada direto na saída.

def problema1008():
    #Problema 1008 - Salary
    num = int(input())
    horas = int(input())
    valor = float(input())
    salario = horas * valor
    print(f"NUMBER = {num}")
    print(f"SALARY = U$ {salario:.2f}")
    # Melhoria: cálculo separado por clareza, f-strings.

def problema1009():
    #Problema 1009 - Salary with Bonus
    nome = input()
    salario = float(input())
    vendas = float(input())
    print(f"TOTAL = R$ {salario + (vendas * 0.15):.2f}")
    # Melhoria: cálculo direto na impressão, nome ignorado por não ser usado.

def problema1010():
    #Problema 1010 - Simple Calculate
    p1 = list(map(float, input().split()))
    p2 = list(map(float, input().split()))
    total = (p1[1] * p1[2]) + (p2[1] * p2[2])
    print(f"VALOR A PAGAR: R$ {total:.2f}")
    # Melhoria: uso de map + split para entrada e f-string.

def problema1011():
    #Problema 1011 - Sphere
    R = float(input())
    volume = (4/3.0) * 3.14159 * R**3
    print(f"VOLUME = {volume:.3f}")
    # Melhoria: uso de `**3` no lugar de R*R*R.

def problema1012():
    #Problema 1012 - Area
    A, B, C = map(float, input().split())
    print(f"TRIANGULO: {A * C / 2:.3f}")
    print(f"CIRCULO: {3.14159 * C**2:.3f}")
    print(f"TRAPEZIO: {(A + B) * C / 2:.3f}")
    print(f"QUADRADO: {B**2:.3f}")
    print(f"RETANGULO: {A * B:.3f}")
    # Melhoria: organização direta da fórmula e uso de f-string.

def problema1013():
    #Problema 1013 - The Greatest
    A, B, C = map(int, input().split())
    maior = max(A, B, C)
    print(f"{maior} eh o maior")
    # Melhoria: uso de `max()` ao invés de fórmula com (a+b+abs(a-b))/2.

def problema1014():
    #Problema 1014 - Consumption
    distancia = int(input())
    combustivel = float(input())
    print(f"{distancia / combustivel:.3f} km/l")
    # Melhoria: cálculo direto na impressão.

def problema1015():
    #Problema 1015 - Distance Between Two Points
    x1, y1 = map(float, input().split())
    x2, y2 = map(float, input().split())
    distancia = ((x2 - x1)**2 + (y2 - y1)**2) ** 0.5
    print(f"{distancia:.4f}")
    # Melhoria: fórmula compacta com `** 0.5`.

def problema1016():
    #Problema 1016 - Distance
    distancia = int(input())
    print(f"{distancia * 2} minutos")
    # Melhoria: cálculo direto.

def problema1017():
    #Problema 1017 - Fuel Spent
    tempo = int(input())
    velocidade = int(input())
    gasto = (tempo * velocidade) / 12
    print(f"{gasto:.3f}")
    # Melhoria: expressão em uma linha.

def problema1018():
    #Problema 1018 - Banknotes
    valor = int(input())
    print(valor)
    for nota in [100, 50, 20, 10, 5, 2, 1]:
        qtd = valor // nota
        print(f"{qtd} nota(s) de R$ {nota},00")
        valor %= nota
    # Melhoria: uso de loop ao invés de repetição manual por nota.

def problema1019():
    #Problema 1019 - Time Conversion
    n = int(input())
    h = n // 3600
    m = (n % 3600) // 60
    s = n % 60
    print(f"{h}:{m}:{s}")
    # Melhoria: cálculo sequencial claro com divisões e módulo.

def problema1020():
    #Problema 1020 - Age in Days
    dias = int(input())
    anos = dias // 365
    dias %= 365
    meses = dias // 30
    dias %= 30
    print(f"{anos} ano(s)")
    print(f"{meses} mes(es)")
    print(f"{dias} dia(s)")
    # Melhoria: estrutura clara para decompor dias em anos, meses e dias.

def resolucao1021():
    #Problema 1021 - Banknotes and Coins
    valor = int(float(input()) * 100)  # converte para centavos para evitar erros de ponto flutuante
    notas = [10000, 5000, 2000, 1000, 500, 200]
    moedas = [100, 50, 25, 10, 5, 1]

    print("NOTAS:")
    for nota in notas:
        qtd = valor // nota
        print(f"{qtd} nota(s) de R$ {nota/100:.2f}")
        valor %= nota

    print("MOEDAS:")
    for moeda in moedas:
        qtd = valor // moeda
        print(f"{qtd} moeda(s) de R$ {moeda/100:.2f}")
        valor %= moeda
    # Melhorias:
    # - conversão para centavos (evita problemas com ponto flutuante)
    # - uso de listas e laços para reduzir código repetido
    # - f-strings para exibição formatada
