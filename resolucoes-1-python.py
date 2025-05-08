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