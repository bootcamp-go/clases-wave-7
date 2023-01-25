-- DB: movies_db
USE movies_db;

-- CONSULTAS:
-- ¿Que contiene las primeras 5 filas de movies?
SELECT * FROM movies LIMIT 5;

-- ¿Que peliculas tienen mas de 3 awards?
SELECT * FROM movies as mv WHERE mv.awards > 3;

-- ¿Cuantas peliculas hay por awards?
SELECT COUNT(mv.id) FROM movies as mv
GROUP BY mv.awards;

SELECT mv.awards, COUNT(mv.id) as movies_count FROM movies as mv
GROUP BY mv.awards;
