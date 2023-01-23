-- DB: bank_db
USE bank_db;

SELECT * FROM owner;
SELECT * FROM cards;


-- CONSULTAS:
-- Se necesita la información de cada owner con sus respectivas tarjetas
SELECT * FROM owner as o LEFT JOIN cards as c ON o.id = c.owner_id;

-- Se requisita la información de cada owner y la cantidad de tarjetas que posee cada uno
SELECT o.*, COUNT(c.owner_id) as nro_cards FROM owner as o LEFT JOIN cards as c ON o.id = c.owner_id
GROUP BY o.id
ORDER BY o.name, nro_cards;

-- ¿Como puedo traer la información anterior pero dependiendo si es tarjeta de credito o debito?
SELECT o.*, COUNT(c.owner_id) as nro_cards FROM owner as o LEFT JOIN cards as c ON (o.id = c.owner_id AND c.type = "DEBIT")
GROUP BY o.id;

-- *************************

-- ¿Que tarjetas posee la persona llamada "Jane Doe"?
-- SELECT c.* FROM cards as c JOIN owner as o ON o.name = "Jane Doe";
SELECT c.* FROM cards as c WHERE c.owner_id = (SELECT o.id FROM owner as o WHERE o.name = "Jane Doe")
