-- Problema 2602 - Basic Select
SELECT name FROM customers WHERE state = 'RS';

-- Problema 2603 - Customer Adress
SELECT name, street FROM customers WHERE city = 'Porto Alegre';

-- Problema 2604 - Under 10 or Greater Than 100
SELECT id, name FROM products WHERE price < 10 OR price > 100;

-- Problema 2605 - Executive Representatives
SELECT p.name AS name, f.name AS name FROM products p
JOIN providers f ON p.id_providers = f.id
JOIN categories c ON p.id_categories = c.id
WHERE c.name = 'executive';

-- Problema 2606 - Categories
SELECT p.id AS id, p.name AS name FROM products p
JOIN categories c ON p.id_categories = c.id
WHERE c.name ILIKE 'super%';