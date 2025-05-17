#Problema 1000 - Hello World
print("Hello World!")

#Problema 1001 - Extremely Basic
A = int(input())
B = int(input())
X = A + B
print("X = " + str(X))

#Problema 1002 - Área of the Circle
n = 3.14159
raio = float(input())
area = n * (raio * raio)
print(f"A={area:.4f}") #limitar tamaho das casas decimais

#Problema 1003 - Simple Sum
A = int(input())
B = int(input())
SOMA = A + B
print(f"SOMA = {SOMA}")

#Problema 1004 - Simple Product
A = int(input())
B = int(input())
PROD = A * B
print("PROD =", PROD)

#Problema 1005 - Average 1
a = float(input())
b = float(input())
MEDIA = (((a * 3.5) + (b * 7.5)) / 11)
print(f"MEDIA = {MEDIA:.5f}")

#Problema 1006 - Average 2
A = float(input())
B = float(input())
C = float(input())
MEDIA = ((A * 2) + (B * 3) + (C * 5)) / 10
print(f"MEDIA = {MEDIA:.1f}")

#Problema 1007 - Difference
A = int(input())
B = int(input())
C = int(input())
D = int(input())
DIFERENCA = A * B - C * D
print(f"DIFERENCA = {DIFERENCA}")

#Problema 1008 - Salary
num_empregado = int(input())
horas_trabalhadas = int(input())
salario_hora = float(input())
salario = horas_trabalhadas * salario_hora
print(f"NUMBER = {num_empregado}")
print(f"SALARY = U$ {salario:.2f}")

#Problema 1009 - Salary with Bonus
nome = input()
salario = float(input())
total_vendido = float(input())
print(f"TOTAL = R$ {salario + (total_vendido / 100 * 15):.2f}")

#Problema 1010 - Simple Calculate
prod1 = input()
prod1_split = prod1.split()
prod2 = input()
prod2_split = prod2.split()

prod1_unid = int(prod1_split[1])
prod1_valor = float(prod1_split[2])
prod2_unid = int(prod2_split[1])
prod2_valor = float(prod2_split[2])
total = (prod1_unid * prod1_valor) + (prod2_unid * prod2_valor)

print(f"VALOR A PAGAR: R$ {total:.2f}")

#Problema 1011 - Sphere
pi = 3.14159
R = float(input())
print(f"VOLUME = {(4.0/3.0) * pi * (R * R * R):.3f}")

#Problema 1012 - Area
entrada = input()
splitado = entrada.split()
A = float(splitado[0])
B = float(splitado[1])
C = float(splitado[2])
    
print(f"TRIANGULO: {A * C / 2:.3f}")
print(f"CIRCULO: {3.14159 * (C*C):.3f}")
print(f"TRAPEZIO: {0.5 * (C * (A+B)):.3f}")
print(f"QUADRADO: {B*B:.3f}")
print(f"RETANGULO: {A*B:.3f}")

#Problema 1013 - The Greatest
entrada = input()
splitado = entrada.split()
A = int(splitado[0])
B = int(splitado[1])
C = int(splitado[2])

MaiorAB = (A + B + abs(A - B)) / 2
MaiorABC = (MaiorAB + C + abs(MaiorAB - C)) / 2
print(f"{int(MaiorABC)} eh o maior")

#Problema 1014 - Consumption
distancia = int(input())
combustivel_gasto = float(input())
print(f"{distancia / combustivel_gasto:.3f} km/l")

#Problema 1015 - Distance Between Two Points
entrada = input()
p1string = entrada.split()
p1 = [float(p1string[0]), float(p1string[1])]

entrada = input()
p2string = entrada.split()
p2 = [float(p2string[0]), float(p2string[1])]

distance = ((p2[0] - p1[0]) ** 2 + (p2[1] - p1[1]) ** 2) ** 0.5
print(f"{distance:.4f}")

#Problema 1016 - Distance
distance = int(input())
print(f"{distance * 2} minutos")

#Problema 1017 - Fuel Spent
horas = int(input())
velocidade = int(input())
combustivel_gasto = float(horas * velocidade) / 12
print(f"{combustivel_gasto:.3f}")

#Problema 1018 - Banknotes
entrada = int(input())
print(entrada)
notas = int(entrada / 100)
print(f"{notas} nota(s) de R$ 100,00")
entrada = entrada - notas * 100
notas = int(entrada / 50)
print(f"{notas} nota(s) de R$ 50,00")
entrada = entrada - notas * 50
notas = int(entrada / 20)
print(f"{notas} nota(s) de R$ 20,00")
entrada = entrada - notas * 20
notas = int(entrada / 10)
print(f"{notas} nota(s) de R$ 10,00")
entrada = entrada - notas * 10
notas = int(entrada / 5)
print(f"{notas} nota(s) de R$ 5,00")
entrada = entrada - notas * 5
notas = int(entrada / 2)
print(f"{notas} nota(s) de R$ 2,00")
entrada = entrada - notas * 2
print(f"{entrada} nota(s) de R$ 1,00")

