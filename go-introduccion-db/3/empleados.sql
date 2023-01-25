DROP DATABASE IF EXISTS empresa_db;
CREATE DATABASE IF NOT EXISTS empresa_db;

USE empresa_db;

-- DDL
CREATE TABLE IF NOT EXISTS `departamento` (
	`depto_nro` varchar(7) NOT NULL,
    `nombre_depto` varchar(100) NOT NULL,
    `localidad` varchar(250) NOT NULL,
    PRIMARY KEY (`depto_nro`)
);

CREATE TABLE IF NOT EXISTS `empleados` (
	`cod_emp` varchar(6) NOT NULL,
    `nombre` varchar(50) NOT NULL,
    `apellido` varchar(50) NOT NULL,
    `puesto` varchar(50) NOT NULL,
    `fecha_alta` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `salario` int unsigned NOT NULL DEFAULT 0,
    `comision` int unsigned NOT NULL DEFAULT 0,
    `depto_nro` varchar(7) NOT NULL,
    PRIMARY KEY (`cod_emp`),
    KEY `idx_fk_depto_nro` (`depto_nro`),
    CONSTRAINT `fk_depto_nro` FOREIGN KEY (`depto_nro`) REFERENCES `departamento` (`depto_nro`)
);

-- DATASET
INSERT INTO `departamento` (`depto_nro`, `nombre_depto`, `localidad`) VALUES ("D-000-1", "Software", "Los Tigres");
INSERT INTO `departamento` (`depto_nro`, `nombre_depto`, `localidad`) VALUES ("D-000-2", "Sistemas", "Guadalupe");
INSERT INTO `departamento` (`depto_nro`, `nombre_depto`, `localidad`) VALUES ("D-000-3", "Contabilidad", "La Roca");
INSERT INTO `departamento` (`depto_nro`, `nombre_depto`, `localidad`) VALUES ("D-000-4", "Ventas", "Plata");

INSERT INTO `empleados` (`cod_emp`, `nombre`, `apellido`, `puesto`, `fecha_alta`, `salario`, `comision`, `depto_nro`) VALUES ("E-0001", "César", "Piñero", "Vendedor", "2018-05-12", 80000, 15000, "D-000-4");
INSERT INTO `empleados` (`cod_emp`, `nombre`, `apellido`, `puesto`, `fecha_alta`, `salario`, `comision`, `depto_nro`) VALUES ("E-0002", "Yosep", "Kowaleski", "Analista", "2015-07-14", 140000, 0, "D-000-2");
INSERT INTO `empleados` (`cod_emp`, `nombre`, `apellido`, `puesto`, `fecha_alta`, `salario`, `comision`, `depto_nro`) VALUES ("E-0003", "Mariela", "Barrios", "Director", "2014-06-05", 185000, 0, "D-000-3");
INSERT INTO `empleados` (`cod_emp`, `nombre`, `apellido`, `puesto`, `fecha_alta`, `salario`, `comision`, `depto_nro`) VALUES ("E-0004", "Jonathan", "Aguilera", "Vendedor", "2015-06-03", 85000, 10000, "D-000-4");
INSERT INTO `empleados` (`cod_emp`, `nombre`, `apellido`, `puesto`, `fecha_alta`, `salario`, `comision`, `depto_nro`) VALUES ("E-0005", "Daniel", "Brezezicki", "Vendedor", "2018-03-03", 83000, 10000, "D-000-4");
INSERT INTO `empleados` (`cod_emp`, `nombre`, `apellido`, `puesto`, `fecha_alta`, `salario`, `comision`, `depto_nro`) VALUES ("E-0006", "Mito", "Barchuk", "Presidente", "2014-06-05", 190000, 0, "D-000-3");
INSERT INTO `empleados` (`cod_emp`, `nombre`, `apellido`, `puesto`, `fecha_alta`, `salario`, `comision`, `depto_nro`) VALUES ("E-0007", "Emilio", "Galarza", "Desarrollador", "2014-08-02", 60000, 0, "D-000-1");