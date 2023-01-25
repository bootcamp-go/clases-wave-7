-- DB: empresa_db;
USE empresa_db;

-- CONSULTAS:
-- Seleccionar el nombre, el puesto de empleados y la localidad de los departamentos donde trabajan los vendedores.
SELECT e.nombre, e.puesto, d.localidad FROM empleados as e JOIN departamento as d ON e.depto_nro = d.depto_nro
WHERE e.puesto = "Vendedor";

SELECT e.nombre, e.puesto, (SELECT d.localidad FROM departamento as d WHERE d.depto_nro = e.depto_nro) as localidad
FROM empleados as e
WHERE e.puesto = "Vendedor";

-- Visualizar los departamentos con más de cinco empleados.
SELECT d.* FROM empleados as e JOIN departamento as d ON e.depto_nro = d.depto_nro
GROUP BY d.depto_nro HAVING COUNT(e.cod_emp) > 1;

-- Mostrar el nombre, salario de los empleados y nombre del departamento de empleados que tengan el mismo puesto que ‘Mito Barchuk’.
SELECT e.nombre, e.salario, d.nombre_depto FROM empleados as e JOIN departamento as d ON e.depto_nro = d.depto_nro
WHERE e.puesto = (SELECT e2.puesto FROM empleados as e2 WHERE e2.nombre = "Mito" AND e2.apellido = "Barchuk");

-- Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.
SELECT e.* FROM empleados as e WHERE e.depto_nro = (SELECT d.depto_nro FROM departamento as d WHERE d.nombre_depto = "Contabilidad")
ORDER BY e.nombre DESC;

-- Mostrar el nombre del empleado que tiene el salario más bajo.
SELECT e.nombre FROM empleados e ORDER BY e.salario LIMIT 1;
SELECT e.nombre FROM empleados as e WHERE e.salario = (SELECT MIN(e2.salario) FROM empleados as e2);

-- Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.
SELECT e.* FROM empleados as e
WHERE e.salario = (
	SELECT MAX(e2.salario) FROM empleados as e2 WHERE e2.depto_nro = (
		SELECT d.depto_nro FROM departamento as d WHERE d.nombre_depto = "Ventas"
    )
);

