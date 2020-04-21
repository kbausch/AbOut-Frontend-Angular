-- assessment_light_popluate.sql is a script for the C# unit tests to populate 
-- the tables more quickly than the assessment_populate.sql script.
-- Make sure to update it after modifying assessment_populate.sql.

USE assessment;

SET FOREIGN_KEY_CHECKS=0;

TRUNCATE TABLE `semesters`;
INSERT INTO `semesters` VALUES (1,'Fall 2016'),(4,'Fall 2017'),(7,'Fall 2018'),(10,'Fall 2019'),(2,'Spring 2017'),(5,'Spring 2018'),(8,'Spring 2019'),(11,'Spring 2020'),(3,'Summer 2017'),(6,'Summer 2018'),(9,'Summer 2019');

TRUNCATE TABLE `intervals`;
INSERT INTO `intervals` VALUES (1,1,NULL),(2,1,3),(3,4,6),(4,7,9),(5,4,9);

TRUNCATE TABLE `programs`;
INSERT INTO `programs` VALUES (1,'SE','Software Engineering',11,1),(2,'CS','Computer Science',11,1),(3,'EE','Electrical Engineering',11,1);

TRUNCATE TABLE `users`;
INSERT INTO `users` VALUES (1,'kbausch','Kaleb','Bausch',0),(2,'dbrush','Diedrich','Brush',0),(3,'dcaron','Dalton','Caron',0),(4,'xdolence','Xaavan','Dolence\n',0),(5,'mfrisbee','Marcus','Frisbee',0),(6,'ehudges','Eli','Hodges',0),(7,'jjenkins','Jackson','Jenkins',0),(8,'wmackubbin','Wyatt','Mackubbin',0),(9,'jmichelotti','Jake','Michelotti',1),(10,'jvesco','Jacob','Vesco',1),(11,'cschahczenski','Celia','Schahczenski',0);

TRUNCATE TABLE `subjects`;
INSERT INTO `subjects` VALUES (4,'CSCI'),(5,'ESOF'),(6,'M');

TRUNCATE TABLE `courses`;
INSERT INTO `courses` VALUES (1,4,'135','Fundamentals of Computing I',1),(2,4,'136','Fundamentals of Computing II',1),(3,4,'246','Discrete Structures',1),(4,4,'255','Embedded Systems',1),(5,5,'326','Software Maintenance',1),(6,4,'172','Calculus I',1),(7,5,'411','Verification and Validation',1);

TRUNCATE TABLE `program_courses`;
INSERT INTO `program_courses` VALUES (1,1,1,1),(2,2,1,1),(3,1,2,1),(4,2,2,1),(5,1,3,1),(6,2,3,1),(10,1,4,1),(11,2,4,1),(12,3,4,1),(13,1,5,1),(14,2,5,1),(15,2,6,1),(16,3,6,1),(17,1,6,1),(18,1,7,1);

TRUNCATE TABLE `prefixes`;
INSERT INTO `prefixes` VALUES (2,'CAC'),(1,'EAC'),(3,'SAC');

TRUNCATE TABLE `outcomes`;
INSERT INTO `outcomes` VALUES (1,1,'1','an ability to identify, formulate, and solve complex engineering problems by applying principles of engineering, science, and mathematics',1),(2,1,'2','an ability to apply engineering design to produce solutions that meet specified needs with consideration of public health, safety, and welfare, as well as global, cultural, social, environmental, and economic factors',1),(3,2,'1','an ability to analyze a complex computing problem and to apply principles of computing and other relevant disciplines to identify solutions',1);

TRUNCATE TABLE `program_outcomes`;
INSERT INTO `program_outcomes` VALUES (1,1,1,1),(2,2,1,1),(3,3,2,1);

TRUNCATE TABLE `measures`;
-- No data exists yet for this table.

TRUNCATE TABLE `permissions`;
INSERT INTO `permissions` VALUES (1,1,1,1,0),(2,2,2,1,0),(3,3,3,0,0),(4,4,1,0,1),(5,4,2,0,1),(6,4,3,0,0);

SET FOREIGN_KEY_CHECKS=1;