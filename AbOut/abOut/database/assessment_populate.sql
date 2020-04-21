-- MySQL dump 10.13  Distrib 8.0.18, for Win64 (x86_64)
--
-- Host: localhost    Database: assessment
-- ------------------------------------------------------
-- Server version	5.7.28-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

USE `assessment`;

--
-- Table structure for table `courses`
--

DROP TABLE IF EXISTS `courses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `courses` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `subject_id` int(11) NOT NULL,
  `number` char(5) COLLATE utf8_bin NOT NULL,
  `name` varchar(70) COLLATE utf8_bin NOT NULL,
  `interval_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `courses_interval_id_idx` (`interval_id`),
  KEY `subject_id_idx` (`subject_id`),
  CONSTRAINT `courses_interval_id` FOREIGN KEY (`interval_id`) REFERENCES `intervals` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `subject_id` FOREIGN KEY (`subject_id`) REFERENCES `subjects` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `courses`
--

LOCK TABLES `courses` WRITE;
/*!40000 ALTER TABLE `courses` DISABLE KEYS */;
INSERT INTO `courses` VALUES (1,4,'135','Fundamentals of Computing I',1),(2,4,'136','Fundamentals of Computing II',1),(3,4,'246','Discrete Structures',1),(4,4,'255','Embedded Systems',1),(5,5,'326','Software Maintenance',1),(6,4,'172','Calculus I',1),(7,5,'411','Verification and Validation',1);
/*!40000 ALTER TABLE `courses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `intervals`
--

DROP TABLE IF EXISTS `intervals`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `intervals` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `begin` int(11) NOT NULL,
  `end` int(11) DEFAULT NULL COMMENT 'Readly I want too allow nulls.',
  PRIMARY KEY (`id`),
  KEY `begin_idx` (`begin`),
  KEY `end_idx` (`end`),
  CONSTRAINT `begin` FOREIGN KEY (`begin`) REFERENCES `semesters` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `end` FOREIGN KEY (`end`) REFERENCES `semesters` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `intervals`
--

LOCK TABLES `intervals` WRITE;
/*!40000 ALTER TABLE `intervals` DISABLE KEYS */;
INSERT INTO `intervals` VALUES (1,1,NULL),(2,1,3),(3,4,6),(4,7,9),(5,4,9);
/*!40000 ALTER TABLE `intervals` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `measures`
--

DROP TABLE IF EXISTS `measures`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `measures` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `program_outcomes_id` int(11) NOT NULL,
  `program_courses_id` int(11) NOT NULL,
  `interval_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `program_outcomes_id_idx` (`program_outcomes_id`),
  KEY `program_courses_id_idx` (`program_courses_id`),
  KEY `interval_id_idx` (`interval_id`),
  CONSTRAINT `measures_interval_id` FOREIGN KEY (`interval_id`) REFERENCES `intervals` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `program_courses_id` FOREIGN KEY (`program_courses_id`) REFERENCES `program_courses` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `program_outcomes_id` FOREIGN KEY (`program_outcomes_id`) REFERENCES `program_outcomes` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `measures`
--

LOCK TABLES `measures` WRITE;
/*!40000 ALTER TABLE `measures` DISABLE KEYS */;
/*!40000 ALTER TABLE `measures` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `outcomes`
--

DROP TABLE IF EXISTS `outcomes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `outcomes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `prefix_id` int(11) NOT NULL,
  `identifier` char(5) COLLATE utf8_bin NOT NULL,
  `text` varchar(300) COLLATE utf8_bin NOT NULL,
  `interval_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `outcomeID_UNIQUE` (`id`),
  KEY `interval_id_idx` (`interval_id`),
  KEY `prefix_id` (`prefix_id`),
  CONSTRAINT `outcomes_interval_id` FOREIGN KEY (`interval_id`) REFERENCES `intervals` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `prefix_id` FOREIGN KEY (`prefix_id`) REFERENCES `prefixes` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `outcomes`
--

LOCK TABLES `outcomes` WRITE;
/*!40000 ALTER TABLE `outcomes` DISABLE KEYS */;
INSERT INTO `outcomes` VALUES (1,1,'1','an ability to identify, formulate, and solve complex engineering problems by applying principles of engineering, science, and mathematics',1),(2,1,'2','an ability to apply engineering design to produce solutions that meet specified needs with consideration of public health, safety, and welfare, as well as global, cultural, social, environmental, and economic factors',1),(3,2,'1','an ability to analyze a complex computing problem and to apply principles of computing and other relevant disciplines to identify solutions',1);
/*!40000 ALTER TABLE `outcomes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `prefixes`
--

DROP TABLE IF EXISTS `prefixes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `prefixes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `text` char(5) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `outcomePrefix_UNIQUE` (`text`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `prefixes`
--

LOCK TABLES `prefixes` WRITE;
/*!40000 ALTER TABLE `prefixes` DISABLE KEYS */;
INSERT INTO `prefixes` VALUES (2,'CAC'),(1,'EAC'),(3,'SAC');
/*!40000 ALTER TABLE `prefixes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `program_courses`
--

DROP TABLE IF EXISTS `program_courses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `program_courses` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `program_id` int(11) NOT NULL,
  `course_id` int(11) NOT NULL,
  `interval_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `program_id_idx` (`program_id`),
  KEY `course_id_idx` (`course_id`),
  KEY `program_courses_interval_id_idx` (`interval_id`),
  CONSTRAINT `course_id` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `program_courses_interval_id` FOREIGN KEY (`interval_id`) REFERENCES `intervals` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `program_courses_program_id` FOREIGN KEY (`program_id`) REFERENCES `programs` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `program_courses`
--

LOCK TABLES `program_courses` WRITE;
/*!40000 ALTER TABLE `program_courses` DISABLE KEYS */;
INSERT INTO `program_courses` VALUES (1,1,1,1),(2,2,1,1),(3,1,2,1),(4,2,2,1),(5,1,3,1),(6,2,3,1),(10,1,4,1),(11,2,4,1),(12,3,4,1),(13,1,5,1),(14,2,5,1),(15,2,6,1),(16,3,6,1),(17,1,6,1),(18,1,7,1);
/*!40000 ALTER TABLE `program_courses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `program_outcomes`
--

DROP TABLE IF EXISTS `program_outcomes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `program_outcomes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `program_id` int(11) NOT NULL,
  `outcome_id` int(11) NOT NULL,
  `interval_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `programOutcomeID_UNIQUE` (`id`),
  KEY `program_id_idx` (`program_id`),
  KEY `outcome_id_idx` (`outcome_id`),
  KEY `interval_id_idx` (`interval_id`),
  CONSTRAINT `outcome_id` FOREIGN KEY (`outcome_id`) REFERENCES `outcomes` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `program_outcomes_interval_id` FOREIGN KEY (`interval_id`) REFERENCES `intervals` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `program_outcomes_program_id` FOREIGN KEY (`program_id`) REFERENCES `programs` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `program_outcomes`
--

LOCK TABLES `program_outcomes` WRITE;
/*!40000 ALTER TABLE `program_outcomes` DISABLE KEYS */;
INSERT INTO `program_outcomes` VALUES (1,1,1,1),(2,2,1,1),(3,3,2,1);
/*!40000 ALTER TABLE `program_outcomes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `programs`
--

DROP TABLE IF EXISTS `programs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `programs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `abbrev` char(5) COLLATE utf8_bin NOT NULL,
  `name` varchar(45) COLLATE utf8_bin NOT NULL,
  `current_semester_id` int(11) NOT NULL,
  `interval_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `current_semester_id_idx` (`current_semester_id`),
  KEY `interval_id_idx` (`interval_id`),
  CONSTRAINT `current_semester_id` FOREIGN KEY (`current_semester_id`) REFERENCES `semesters` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `interval_id` FOREIGN KEY (`interval_id`) REFERENCES `intervals` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `programs`
--

LOCK TABLES `programs` WRITE;
/*!40000 ALTER TABLE `programs` DISABLE KEYS */;
INSERT INTO `programs` VALUES (1,'SE','Software Engineering',11,1),(2,'CS','Computer Science',11,1),(3,'EE','Electrical Engineering',11,1);
/*!40000 ALTER TABLE `programs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `semesters`
--

DROP TABLE IF EXISTS `semesters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `semesters` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` char(11) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `semesterID_UNIQUE` (`id`),
  UNIQUE KEY `semesterName_UNIQUE` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `semesters`
--

LOCK TABLES `semesters` WRITE;
/*!40000 ALTER TABLE `semesters` DISABLE KEYS */;
INSERT INTO `semesters` VALUES (1,'Fall 2016'),(4,'Fall 2017'),(7,'Fall 2018'),(10,'Fall 2019'),(2,'Spring 2017'),(5,'Spring 2018'),(8,'Spring 2019'),(11,'Spring 2020'),(3,'Summer 2017'),(6,'Summer 2018'),(9,'Summer 2019');
/*!40000 ALTER TABLE `semesters` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `subjects`
--

DROP TABLE IF EXISTS `subjects`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `subjects` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `text` char(5) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `subjects`
--

LOCK TABLES `subjects` WRITE;
/*!40000 ALTER TABLE `subjects` DISABLE KEYS */;
INSERT INTO `subjects` VALUES (4,'CSCI'),(5,'ESOF'),(6,'M');
/*!40000 ALTER TABLE `subjects` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cas` char(25) COLLATE utf8_bin NOT NULL,
  `first` char(25) COLLATE utf8_bin NOT NULL,
  `last` char(25) COLLATE utf8_bin NOT NULL,
  `superuser` tinyint NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `userID_UNIQUE` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'kbausch','Kaleb','Bausch',0),(2,'dbrush','Diedrich','Brush',0),(3,'dcaron','Dalton','Caron',0),(4,'xdolence','Xaavan','Dolence\n',0),(5,'mfrisbee','Marcus','Frisbee',0),(6,'ehudges','Eli','Hodges',0),(7,'jjenkins','Jackson','Jenkins',0),(8,'wmackubbin','Wyatt','Mackubbin',0),(9,'jmichelotti','Jake','Michelotti',1),(10,'jvesco','Jacob','Vesco',1),(11,'cschahczenski','Celia','Schahczenski',0);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

--
-- Table structure for table `permissions`
--

DROP TABLE IF EXISTS `permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `permissions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `program_id` int(11) NOT NULL,
  `is_manager` TINYINT NOT NULL DEFAULT 0,
  `is_observer` TINYINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `permissions_user_id_idx` (`user_id`),
  KEY `permissions_program_id_idx` (`program_id`),
  CONSTRAINT `permissions_program_id` FOREIGN KEY (`program_id`) REFERENCES `programs` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `permissions_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `permissions`
--

LOCK TABLES `permissions` WRITE;
/*!40000 ALTER TABLE `permissions` DISABLE KEYS */;
INSERT INTO `permissions` VALUES (1,1,1,1,0),(2,2,2,1,0),(3,3,3,0,0),(4,4,1,0,1),(5,4,2,0,1),(6,4,3,0,0);
/*!40000 ALTER TABLE `permissions` ENABLE KEYS */;
UNLOCK TABLES;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-02-13 11:51:33
