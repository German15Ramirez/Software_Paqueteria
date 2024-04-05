-- MySQL dump 10.13  Distrib 8.0.36, for Linux (x86_64)
--
-- Host: localhost    Database: gestion_paqueteria
-- ------------------------------------------------------
-- Server version	8.0.36

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `clientes`
--

DROP TABLE IF EXISTS `clientes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `clientes` (
  `cliente_id` int NOT NULL,
  `nombre` varchar(45) DEFAULT NULL,
  `nit` int DEFAULT NULL,
  `direccion` varchar(45) DEFAULT NULL,
  `telefono` int DEFAULT NULL,
  PRIMARY KEY (`cliente_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `clientes`
--

LOCK TABLES `clientes` WRITE;
/*!40000 ALTER TABLE `clientes` DISABLE KEYS */;
INSERT INTO `clientes` VALUES (1,'Pablo',332432,'22 av. zona 1, quetzgo',75386392),(2,'Luis',424232,'4ta calle solola',90743623),(3,'Abraham',423424,'8va avenida, zona 10, Guatemala',12567643),(4,'Carlos',32322,'10ma. calle, Piedras Negras, Huehuetenango',43765219),(5,'Amilcar',657565,'Calle Real, zona 1, Quetzaltenango',28436533);
/*!40000 ALTER TABLE `clientes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ganancias`
--

DROP TABLE IF EXISTS `ganancias`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ganancias` (
  `ganancias_id` int NOT NULL,
  `ruta_id` int DEFAULT NULL,
  `fecha` date DEFAULT NULL,
  `ingresos` float DEFAULT NULL,
  `costos` float DEFAULT NULL,
  `ganancias` float DEFAULT NULL,
  PRIMARY KEY (`ganancias_id`),
  KEY `fk_ruta_id_ganancias_idx` (`ruta_id`),
  CONSTRAINT `fk_ruta_id_ganancias` FOREIGN KEY (`ruta_id`) REFERENCES `rutas` (`ruta_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ganancias`
--

LOCK TABLES `ganancias` WRITE;
/*!40000 ALTER TABLE `ganancias` DISABLE KEYS */;
/*!40000 ALTER TABLE `ganancias` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `operadores`
--

DROP TABLE IF EXISTS `operadores`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `operadores` (
  `operador_id` int NOT NULL,
  `usuario_id` int DEFAULT NULL,
  `punto_de_control_id` int DEFAULT NULL,
  PRIMARY KEY (`operador_id`),
  KEY `fk_usuario_id_operadores_idx` (`usuario_id`),
  KEY `fk_punto_de_control_id_operadores_idx` (`punto_de_control_id`),
  CONSTRAINT `fk_punto_de_control_id_operadores` FOREIGN KEY (`punto_de_control_id`) REFERENCES `puntos_de_control` (`puntos_de_control_id`),
  CONSTRAINT `fk_usuario_id_operadores` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`usuario_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `operadores`
--

LOCK TABLES `operadores` WRITE;
/*!40000 ALTER TABLE `operadores` DISABLE KEYS */;
/*!40000 ALTER TABLE `operadores` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `paquetes`
--

DROP TABLE IF EXISTS `paquetes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `paquetes` (
  `paquete_id` int NOT NULL,
  `cliente_id` int NOT NULL,
  `peso_libras` float DEFAULT NULL,
  `destino` varchar(45) DEFAULT NULL,
  `fecha_ingreso` date DEFAULT NULL,
  `fecha_salida` date DEFAULT NULL,
  `estado` enum('En Ruta','Entregado') DEFAULT NULL,
  `ruta_id` int NOT NULL,
  PRIMARY KEY (`paquete_id`),
  KEY `fk_cliente_id_paquetes_idx` (`cliente_id`),
  KEY `fk_ruta_id_paquetes_idx` (`ruta_id`),
  CONSTRAINT `fk_cliente_id_paquetes` FOREIGN KEY (`cliente_id`) REFERENCES `clientes` (`cliente_id`),
  CONSTRAINT `fk_ruta_id_paquetes` FOREIGN KEY (`ruta_id`) REFERENCES `rutas` (`ruta_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `paquetes`
--

LOCK TABLES `paquetes` WRITE;
/*!40000 ALTER TABLE `paquetes` DISABLE KEYS */;
INSERT INTO `paquetes` VALUES (1,1,2.5,'Ciudad Capital','2023-10-11','2023-10-18','Entregado',1),(2,1,3,'Ciudad Capital','2023-11-10','2023-11-20','Entregado',1),(3,1,4,'Ciudad Capital','2023-09-19','2023-09-29','En Ruta',1),(4,1,12,'Ciudad Capital','2024-03-28','2024-04-02','Entregado',1),(5,1,1.3,'Ciudad Capital','2024-03-26','2024-04-01','En Ruta',1),(6,1,4.6,'Ciudad Capital','2024-02-10','2024-02-16','Entregado',1),(7,1,3.8,'Solola','2023-11-10','2023-11-14','En Ruta',2),(8,1,3.9,'Solola','2023-11-07','2023-11-13','Entregado',2),(9,2,2.4,'Quetzaltenango','2024-03-16','2024-03-20','Entregado',3),(10,2,2.5,'Totonicapan','2024-02-10','2024-02-19','Entregado',4),(11,2,15,'Quiche','2024-02-03','2024-02-10','En Ruta',5),(12,2,3.3,'Quetzaltenango','2024-02-10','2024-02-15','Entregado',3),(13,2,2.5,'Solola','2024-03-15','2024-03-22','En Ruta',2),(14,2,9.5,'Totonicapan','2024-03-10','2024-03-12','Entregado',4),(15,2,3.7,'Quetzaltenango','2024-01-02','2024-01-07','En Ruta',3),(16,2,4.2,'Solola','2024-02-01','2024-02-09','Entregado',2),(17,3,2.1,'Solola','2024-01-15','2024-01-25','Entregado',2),(18,3,10,'Ciudad Capital','2024-02-01','2024-02-10','Entregado',1),(19,3,2.9,'Solola','2024-01-10','2024-01-19','Entregado',2),(20,3,2.8,'Quetzaltenango','2023-10-09','2023-10-15','Entregado',3),(21,3,2.1,'Ciudad Capital','2023-12-01','2023-12-06','Entregado',1),(22,3,9,'Ciudad Capital','2023-08-10','2024-10-10','En Ruta',1),(23,3,2.4,'Quiche','2023-10-01','2023-10-10','En Ruta',5),(24,3,7.7,'Quetzaltenango','2024-02-10','2024-02-14','En Ruta',3),(25,4,7.4,'Totonicapan','2024-03-12','2024-03-18','En Ruta',4),(26,4,7.2,'Quiche','2024-02-02','2024-02-04','En Ruta',5),(27,4,7.2,'Ciudad Capital','2024-01-10','2024-01-14','En Ruta',1),(28,4,7.9,'Quetzaltenango','2023-03-10','2023-03-19','Entregado',3),(29,4,3.7,'Solola','2023-02-10','2023-02-16','Entregado',2),(30,4,4.2,'Quetzaltenango','2023-01-10','2023-01-20','Entregado',3),(31,4,4.9,'Totonicapan','2024-02-10','2024-02-18','Entregado',4),(32,4,10.3,'Quiche','2024-02-25','2024-02-29','Entregado',5),(33,5,8.2,'Quiche','2024-01-12','2024-01-17','Entregado',5),(34,5,11.8,'Quetzaltenango','2024-02-10','2024-02-14','Entregado',3),(35,5,1.1,'Solola','2023-10-12','2023-10-21','Entregado',2),(36,5,2.3,'Ciudad Capital','2023-01-12','2023-01-21','Entregado',1),(37,5,4.3,'Quetzaltenango','2024-02-10','2024-02-15','Entregado',3),(38,5,2.7,'Solola','2024-03-12','2024-03-18','En Ruta',2),(39,5,2.5,'Totonicapan','2023-01-12','2023-01-18','En Ruta',4),(40,5,2.8,'Quiche','2024-02-12','2024-02-18','En Ruta',5);
/*!40000 ALTER TABLE `paquetes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `puntos_de_control`
--

DROP TABLE IF EXISTS `puntos_de_control`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `puntos_de_control` (
  `puntos_de_control_id` int NOT NULL,
  `ruta_id` int DEFAULT NULL,
  `nombre` varchar(45) DEFAULT NULL,
  `capacidad_cola` int DEFAULT NULL,
  `tarifa_operacional` float DEFAULT NULL,
  PRIMARY KEY (`puntos_de_control_id`),
  KEY `fk_ruta_id_puntos_de_control_idx` (`ruta_id`),
  CONSTRAINT `fk_ruta_id_puntos_de_control` FOREIGN KEY (`ruta_id`) REFERENCES `rutas` (`ruta_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `puntos_de_control`
--

LOCK TABLES `puntos_de_control` WRITE;
/*!40000 ALTER TABLE `puntos_de_control` DISABLE KEYS */;
INSERT INTO `puntos_de_control` VALUES (11,1,'Punto 1 Ruta 1',10,20),(12,1,'Punto 2 Ruta 1',10,20),(13,1,'Punto 3 Ruta 1',10,20),(14,1,'Punto 4 Ruta 1',10,20),(21,2,'Punto 1 Ruta 2',10,20),(22,2,'Punto 2 Ruta 2',10,20),(23,2,'Punto 3 Ruta 2',10,20),(24,2,'Punto 4 Ruta 2',10,20),(31,3,'Punto 1 Ruta 3',10,20),(32,3,'Punto 2 Ruta 3',10,20),(33,3,'Punto 3 Ruta 3',10,20),(34,3,'Punto 4 Ruta 3',10,20),(41,4,'Punto 1 Ruta 4',10,20),(42,4,'Punto 2 Ruta 4',10,20),(43,4,'Punto 3 Ruta 4',10,20),(44,4,'Punto 4 Ruta 4',10,20),(51,5,'Punto 1 Ruta 5',10,20),(52,5,'Punto 2 Ruta 5',10,20),(53,5,'Punto 3 Ruta 5',10,20),(54,5,'Punto 4 Ruta 5',10,20);
/*!40000 ALTER TABLE `puntos_de_control` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rutas`
--

DROP TABLE IF EXISTS `rutas`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rutas` (
  `ruta_id` int NOT NULL,
  `destino` varchar(45) DEFAULT NULL,
  `tarifa_operacion` float DEFAULT NULL,
  `capacidad_maxima` int DEFAULT NULL,
  PRIMARY KEY (`ruta_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rutas`
--

LOCK TABLES `rutas` WRITE;
/*!40000 ALTER TABLE `rutas` DISABLE KEYS */;
INSERT INTO `rutas` VALUES (1,'Ciudad Capital',100,5),(2,'Solola',120,5),(3,'Quetzaltenango',80,4),(4,'Totonicapan',300,10),(5,'Quiche',250,8);
/*!40000 ALTER TABLE `rutas` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `usuarios`
--

DROP TABLE IF EXISTS `usuarios`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `usuarios` (
  `usuario_id` int NOT NULL,
  `nombre` varchar(45) DEFAULT NULL,
  `correo` varchar(45) DEFAULT NULL,
  `telefono` int DEFAULT NULL,
  `contraseña` varchar(45) DEFAULT NULL,
  `rol` enum('Recepcionista','Administrador','Operador') NOT NULL,
  PRIMARY KEY (`usuario_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `usuarios`
--

LOCK TABLES `usuarios` WRITE;
/*!40000 ALTER TABLE `usuarios` DISABLE KEYS */;
INSERT INTO `usuarios` VALUES (1232,'Alex','alex23@gmail.com',23432323,'algr34','Recepcionista'),(1233,'Jose','uemd@gmail.com',83842392,'ue23kia','Recepcionista'),(1234,'Allan','al323lan@gmail.com',89324387,'allan000__23','Recepcionista'),(1332,'Alexander','alexander@gmail.com',76325383,'alexand0293','Administrador'),(1354,'Luis Castillo','luiscasti11o@gmail.com',12324390,'castillo_@','Administrador'),(1356,'Valeria','vale98@gmail.com',32459300,'vale98_#','Administrador'),(1400,'Francisco','fran382_francisco@gmail.com',76328976,'admin123','Operador'),(1455,'Pedro','pedrito34@gmail.com',32538783,'934pedro','Operador'),(1478,'Oscar','87oscar@gmail.com',56765455,'racso78/','Operador');
/*!40000 ALTER TABLE `usuarios` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-04-04 21:59:12
