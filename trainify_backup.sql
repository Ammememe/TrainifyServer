-- MySQL dump 10.13  Distrib 8.0.37, for Win64 (x86_64)
--
-- Host: localhost    Database: trainify
-- ------------------------------------------------------
-- Server version	8.0.37

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
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `first_name` longtext NOT NULL,
  `last_name` longtext NOT NULL,
  `email` varchar(191) NOT NULL,
  `password` longtext,
  `created_at` datetime(3) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'','','amirhif123@hotmail.com','','2024-11-02 01:27:16.675'),(3,'','','','','2024-11-02 01:33:39.916'),(6,'','','test@test.com','abc123','2024-11-02 01:34:54.652'),(8,'','','test@tesst.com','$2a$10$ULGzL0nMXibJ/E/P7WZWM.ToCcbkf2Ye6EHc5E/ltDb4X1fIeIwlG','2024-11-02 01:36:14.873'),(14,'','','amirhif23@hotmail.com','$2a$10$cEIhdTPQNnbGBDMqJJzjluNGZKNPlG0ts6yqdQ.5OCb3BEwOmwLyi','2024-11-02 01:44:17.835'),(16,'','','test@tessst.com','$2a$10$nZSSmKLJh53BwU6maou33.JJYunaptry3JkH5TYAHkMpPUziR80Ta','2024-11-02 01:47:00.582'),(19,'','','test1@test.com','$2a$10$yuIopiQHiNOoGT9AL6X.GOdnD/KQePzgH7Jw2PgBGduLJh/zNHW76','2024-11-02 20:23:35.344'),(22,'','','test12@test.com','$2a$10$BzD6qiOePCdgO4IfbvZgCuCO69oOyMWCVekCPVftNwOHK7gm1E9m2','2024-11-02 20:25:16.086'),(24,'','','test123@test.com','$2a$10$hDa0SwBvfq13TZDBu4LhQ.u.ab4BCLtr7kFSU8cykQo8bpiuAy1e.','2024-11-02 20:28:25.628'),(26,'','','test5@test.com','$2a$10$7KYpFbDMlUn83mzPVTH5U.oo4Mt8iWg4PO.bAl4b0m7fSI7c7YTzS','2024-11-03 19:14:41.709');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-11-03 22:38:29
