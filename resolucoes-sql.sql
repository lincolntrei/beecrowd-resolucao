-- Problema 2602 - Basic Select
SELECT name FROM customers WHERE state = 'RS';

-- Problema 2603 - Customer Adress
SELECT name, street FROM customers WHERE city = 'Porto Alegre';

-- Problema 2604 - Under 10 or Greater Than 100
SELECT id, name FROM products WHERE price < 10 OR price > 100;