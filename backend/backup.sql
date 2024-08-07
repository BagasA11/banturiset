-- MariaDB dump 10.19  Distrib 10.4.27-MariaDB, for Win64 (AMD64)
--
-- Host: localhost    Database: banturiset
-- ------------------------------------------------------
-- Server version	10.4.34-MariaDB-1:10.4.34+maria~ubu2004

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `balances`
--

DROP TABLE IF EXISTS `balances`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `balances` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `saldo` float NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(50) NOT NULL,
  `desc` longtext DEFAULT NULL,
  `proposal_url` varchar(170) DEFAULT NULL,
  `klirens_url` varchar(170) DEFAULT NULL,
  `fund_until` datetime(3) NOT NULL,
  `dead_line` datetime(3) NOT NULL,
  `milestone` tinyint(4) NOT NULL DEFAULT 1,
  `tkt_level` tinyint(4) NOT NULL DEFAULT 1,
  `cost` float NOT NULL,
  `status` smallint(6) NOT NULL DEFAULT 0,
  `pesan_revisi` longtext DEFAULT NULL,
  `fraud` tinyint(1) NOT NULL DEFAULT 0,
  `pengajuan_id` bigint(20) unsigned NOT NULL,
  `peneliti_id` bigint(20) unsigned NOT NULL,
  `balance_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_balances_deleted_at` (`deleted_at`),
  KEY `fk_balances_pengajuan` (`pengajuan_id`),
  KEY `fk_balances_peneliti` (`peneliti_id`),
  CONSTRAINT `fk_balances_peneliti` FOREIGN KEY (`peneliti_id`) REFERENCES `penelitis` (`id`),
  CONSTRAINT `fk_balances_pengajuan` FOREIGN KEY (`pengajuan_id`) REFERENCES `pengajuans` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `balances`
--

LOCK TABLES `balances` WRITE;
/*!40000 ALTER TABLE `balances` DISABLE KEYS */;
/*!40000 ALTER TABLE `balances` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `budget_details`
--

DROP TABLE IF EXISTS `budget_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `budget_details` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `deskripsi` longtext DEFAULT NULL,
  `tahap` tinyint(3) unsigned NOT NULL DEFAULT 1,
  `cost` float NOT NULL DEFAULT 0,
  `project_id` bigint(20) unsigned DEFAULT NULL,
  `percent` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `start` datetime(3) DEFAULT NULL,
  `end` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_projects_budget_details` (`project_id`),
  CONSTRAINT `fk_projects_budget_details` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `budget_details`
--

LOCK TABLES `budget_details` WRITE;
/*!40000 ALTER TABLE `budget_details` DISABLE KEYS */;
INSERT INTO `budget_details` VALUES (1,'Membeli Radar Cuaca: Radar cuaca akan digunakan untuk mengambil data cuaca',1,1000000,6,0,NULL,NULL),(2,'Biaya Observasi Lapangan',1,500000,6,0,NULL,NULL);
/*!40000 ALTER TABLE `budget_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `donasis`
--

DROP TABLE IF EXISTS `donasis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `donasis` (
  `id` varchar(191) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `status` longtext NOT NULL,
  `jml` float NOT NULL,
  `fee` float NOT NULL,
  `method` varchar(191) NOT NULL DEFAULT 'OVO',
  `donatur_id` bigint(20) unsigned DEFAULT NULL,
  `balance_id` bigint(20) unsigned DEFAULT NULL,
  `project_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_donasis_deleted_at` (`deleted_at`),
  KEY `fk_donasis_donatur` (`donatur_id`),
  KEY `fk_balances_donasi` (`balance_id`),
  KEY `fk_projects_donasi` (`project_id`),
  CONSTRAINT `fk_balances_donasi` FOREIGN KEY (`balance_id`) REFERENCES `balances` (`id`),
  CONSTRAINT `fk_donasis_donatur` FOREIGN KEY (`donatur_id`) REFERENCES `donaturs` (`id`),
  CONSTRAINT `fk_projects_donasi` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `donasis`
--

LOCK TABLES `donasis` WRITE;
/*!40000 ALTER TABLE `donasis` DISABLE KEYS */;
INSERT INTO `donasis` VALUES ('invoice-383fg1e','2024-06-17 22:42:19.780','2024-06-17 22:42:19.780',NULL,'PENDING',25000,834.75,'OVO',1,NULL,6),('invoice-5641heb','2024-06-17 22:49:14.367','2024-06-17 22:49:14.366',NULL,'PENDING',25000,834.75,'OVO',1,NULL,6),('invoice-72g3883','2024-06-17 22:37:21.597','2024-06-17 22:37:21.596',NULL,'PENDING',25000,834.75,'OVO',1,NULL,6),('invoice-bf2eh75','2024-06-17 22:49:18.411','2024-06-17 22:49:18.411',NULL,'PENDING',25000,834.75,'OVO',1,NULL,6),('invoice-bfhb26c','2024-06-17 22:48:28.826','2024-06-17 22:48:28.825',NULL,'PENDING',25000,834.75,'OVO',1,NULL,6),('invoice-cf37a3b','2024-06-17 22:37:15.414','2024-06-17 22:37:15.413',NULL,'PENDING',25000,834.75,'OVO',1,NULL,6);
/*!40000 ALTER TABLE `donasis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `donaturs`
--

DROP TABLE IF EXISTS `donaturs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `donaturs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_users_donatur` (`user_id`),
  CONSTRAINT `fk_users_donatur` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `donaturs`
--

LOCK TABLES `donaturs` WRITE;
/*!40000 ALTER TABLE `donaturs` DISABLE KEYS */;
INSERT INTO `donaturs` VALUES (1,1);
/*!40000 ALTER TABLE `donaturs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payouts`
--

DROP TABLE IF EXISTS `payouts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `payouts` (
  `id` varchar(191) NOT NULL,
  `tahap` tinyint(3) unsigned NOT NULL,
  `project_id` bigint(20) unsigned DEFAULT NULL,
  `peneliti_id` bigint(20) unsigned DEFAULT NULL,
  `status` longtext NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_projects_payout` (`project_id`),
  KEY `fk_penelitis_payout` (`peneliti_id`),
  CONSTRAINT `fk_penelitis_payout` FOREIGN KEY (`peneliti_id`) REFERENCES `penelitis` (`id`),
  CONSTRAINT `fk_projects_payout` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payouts`
--

LOCK TABLES `payouts` WRITE;
/*!40000 ALTER TABLE `payouts` DISABLE KEYS */;
/*!40000 ALTER TABLE `payouts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `penelitis`
--

DROP TABLE IF EXISTS `penelitis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `penelitis` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `n_ip` varchar(20) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_users_peneliti` (`user_id`),
  CONSTRAINT `fk_users_peneliti` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `penelitis`
--

LOCK TABLES `penelitis` WRITE;
/*!40000 ALTER TABLE `penelitis` DISABLE KEYS */;
INSERT INTO `penelitis` VALUES (4,'00003',9),(5,'010101010',5);
/*!40000 ALTER TABLE `penelitis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pengajuans`
--

DROP TABLE IF EXISTS `pengajuans`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pengajuans` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(20) DEFAULT NULL,
  `desc` longtext DEFAULT NULL,
  `link_wa` varchar(120) DEFAULT NULL,
  `link_panduan` varchar(120) DEFAULT NULL,
  `icon_url` varchar(120) DEFAULT NULL,
  `closed_at` datetime(3) DEFAULT NULL,
  `penyunting_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_pengajuans_title` (`title`),
  KEY `idx_pengajuans_deleted_at` (`deleted_at`),
  KEY `fk_penyuntings_pengajuan` (`penyunting_id`),
  CONSTRAINT `fk_penyuntings_pengajuan` FOREIGN KEY (`penyunting_id`) REFERENCES `penyuntings` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pengajuans`
--

LOCK TABLES `pengajuans` WRITE;
/*!40000 ALTER TABLE `pengajuans` DISABLE KEYS */;
INSERT INTO `pengajuans` VALUES (2,'2024-06-06 19:26:47.601','2024-06-06 19:26:47.602',NULL,'Teknologi Komputer d','skema penelitian ini bertujuan untuk meneliti teknoloi Komputer dan AI','web.whatsapp.com','','https://developers.xendit.co/api-reference/images/logo.png','2024-07-09 19:41:53.000',1),(3,'2024-07-22 10:46:26.392','2024-07-22 10:46:26.392',NULL,'Kesehatan dan Medis','skema penelitian ini bertujuan untuk kesehatan','web.whatsapp.com','','https://developers.xendit.co/api-reference/images/logo.png','2024-10-22 10:46:26.387',1);
/*!40000 ALTER TABLE `pengajuans` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `penyuntings`
--

DROP TABLE IF EXISTS `penyuntings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `penyuntings` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `n_ip` varchar(20) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_users_penyunting` (`user_id`),
  CONSTRAINT `fk_users_penyunting` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `penyuntings`
--

LOCK TABLES `penyuntings` WRITE;
/*!40000 ALTER TABLE `penyuntings` DISABLE KEYS */;
INSERT INTO `penyuntings` VALUES (1,'1234567890123456789',8);
/*!40000 ALTER TABLE `penyuntings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `progresses`
--

DROP TABLE IF EXISTS `progresses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `progresses` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `file_url` longtext NOT NULL,
  `desc` longtext DEFAULT NULL,
  `tahap` tinyint(3) unsigned NOT NULL DEFAULT 1,
  `status` tinyint(4) NOT NULL DEFAULT 0,
  `pesan_revisi` longtext DEFAULT NULL,
  `project_id` bigint(20) unsigned DEFAULT NULL,
  `validasi` tinyint(3) unsigned NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `validated_at` datetime(3) DEFAULT NULL,
  `validator_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_projects_progress` (`project_id`),
  KEY `fk_progresses_penyunting` (`validator_id`),
  CONSTRAINT `fk_progresses_penyunting` FOREIGN KEY (`validator_id`) REFERENCES `penyuntings` (`id`),
  CONSTRAINT `fk_projects_progress` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `progresses`
--

LOCK TABLES `progresses` WRITE;
/*!40000 ALTER TABLE `progresses` DISABLE KEYS */;
/*!40000 ALTER TABLE `progresses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `projects`
--

DROP TABLE IF EXISTS `projects`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `projects` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(50) NOT NULL,
  `desc` longtext DEFAULT NULL,
  `proposal_url` varchar(170) DEFAULT NULL,
  `klirens_url` varchar(170) DEFAULT NULL,
  `fund_until` datetime(3) NOT NULL,
  `dead_line` datetime(3) NOT NULL,
  `milestone` tinyint(4) NOT NULL DEFAULT 1,
  `tkt_level` tinyint(4) NOT NULL DEFAULT 1,
  `cost` float NOT NULL,
  `status` smallint(6) NOT NULL,
  `pengajuan_id` bigint(20) unsigned NOT NULL,
  `peneliti_id` bigint(20) unsigned NOT NULL,
  `pesan_revisi` longtext DEFAULT NULL,
  `fraud` tinyint(1) NOT NULL DEFAULT 0,
  `collected_fund` float NOT NULL,
  `admin_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_projects_deleted_at` (`deleted_at`),
  KEY `fk_pengajuans_project` (`pengajuan_id`),
  KEY `fk_penelitis_project` (`peneliti_id`),
  KEY `fk_projects_penyunting` (`admin_id`),
  CONSTRAINT `fk_penelitis_project` FOREIGN KEY (`peneliti_id`) REFERENCES `penelitis` (`id`),
  CONSTRAINT `fk_pengajuans_project` FOREIGN KEY (`pengajuan_id`) REFERENCES `pengajuans` (`id`),
  CONSTRAINT `fk_projects_penyunting` FOREIGN KEY (`admin_id`) REFERENCES `penyuntings` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `projects`
--

LOCK TABLES `projects` WRITE;
/*!40000 ALTER TABLE `projects` DISABLE KEYS */;
INSERT INTO `projects` VALUES (5,'2024-06-08 21:58:06.426','2024-06-08 21:58:06.427',NULL,'Prediksi Cuaca dengan Algoritma Naive Bayes','Cuaca merupakan aspek penting dalam berbagai kehidupan manusia. Maka memprediksi cuaca dengan akurat sangat penting. Untuk itu perlu adanya model yang dapat memprediksi cuaca dengan akurat',NULL,NULL,'2024-11-05 21:58:06.426','2025-06-29 10:15:59.000',4,9,13000000,-2,2,5,NULL,0,0,NULL),(6,'2024-06-16 11:04:34.750','2024-06-17 20:23:13.916',NULL,'Weather Forecasting dengan svm','Cuaca merupakan aspek penting dalam berbagai kehidupan manusia. Maka memprediksi cuaca dengan akurat sangat penting. Untuk itu perlu adanya model yang dapat memprediksi cuaca dengan akurat','./file/klirens/84gfh62.pdf','./file/klirens/84gfh62.pdf','2024-07-01 22:07:55.000','2027-01-01 11:04:34.745',4,9,13000000,2,2,5,NULL,0,0,NULL),(10,'2024-06-19 05:58:19.863','2024-06-19 05:58:19.863',NULL,'Testing Timestamp','Cuaca merupakan aspek penting dalam berbagai kehidupan manusia. Maka memprediksi cuaca dengan akurat sangat penting. Untuk itu perlu adanya model yang dapat memprediksi cuaca dengan akurat',NULL,NULL,'2024-06-19 05:58:19.861','2028-01-07 05:58:19.860',4,9,13000000,-2,2,5,NULL,0,0,NULL),(12,'2024-06-19 06:06:14.294','2024-07-22 11:38:01.318',NULL,'Testing Timestamp','mengukur akurasi perhitungan waktu pada field fund_until',NULL,'./file/klirens/hf4fc75.pdf','2025-08-25 14:06:14.292','2028-01-07 06:06:14.291',4,9,13000000,-2,2,5,NULL,0,0,NULL),(13,'2024-07-22 10:47:49.995','2024-07-22 11:38:26.805',NULL,'pembuatan obat penyakit xyz','xyz ..... sfjmefuei','./file/proposal/38gbf29.pdf','./file/klirens/hf4fc75.pdf','2025-02-08 02:47:49.982','2026-03-16 10:47:49.982',2,9,1300000,-2,3,5,NULL,0,0,NULL);
/*!40000 ALTER TABLE `projects` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tahapans`
--

DROP TABLE IF EXISTS `tahapans`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tahapans` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `cost_percent` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `start` datetime(3) NOT NULL,
  `end` datetime(3) NOT NULL,
  `project_id` bigint(20) unsigned DEFAULT NULL,
  `tahap` tinyint(3) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `fk_projects_tahapan` (`project_id`),
  CONSTRAINT `fk_projects_tahapan` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tahapans`
--

LOCK TABLES `tahapans` WRITE;
/*!40000 ALTER TABLE `tahapans` DISABLE KEYS */;
INSERT INTO `tahapans` VALUES (1,30,'2024-06-16 19:38:48.123','2024-07-16 19:38:48.123',6,1),(2,40,'2024-07-22 19:38:48.123','2024-07-30 19:38:48.123',13,1),(3,40,'2024-07-22 19:38:48.123','2024-07-30 19:38:48.123',13,2);
/*!40000 ALTER TABLE `tahapans` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `token_lists`
--

DROP TABLE IF EXISTS `token_lists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `token_lists` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `token` varchar(191) NOT NULL,
  `expired_date` datetime(3) DEFAULT NULL,
  `blacklisted` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_token_lists_token` (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `token_lists`
--

LOCK TABLES `token_lists` WRITE;
/*!40000 ALTER TABLE `token_lists` DISABLE KEYS */;
/*!40000 ALTER TABLE `token_lists` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `f_name` longtext NOT NULL,
  `email` varchar(20) NOT NULL,
  `password` longtext NOT NULL,
  `phone` varchar(191) NOT NULL,
  `role` varchar(191) NOT NULL,
  `institute` longtext NOT NULL,
  `institute_addr` longtext NOT NULL,
  `post_code` varchar(7) NOT NULL,
  `bank` longtext DEFAULT NULL,
  `no_rek` longtext DEFAULT NULL,
  `profile_url` longtext DEFAULT NULL,
  `is_verfied` tinyint(1) NOT NULL DEFAULT 0,
  `is_block` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_email` (`email`),
  UNIQUE KEY `uni_users_phone` (`phone`),
  KEY `idx_users_role` (`role`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Bagas','bagasmipa3@gmail.com','$2a$10$hZTy0WjQieaozWEs1E/S3e/Gn5KDo2m06sTmfL4s4/deMwmdUVgpG','+6282112897283','donatur','Udinus','jl. Imam Bonjol','13340',NULL,NULL,NULL,1,0),(5,'User Suka Spam','spamtest123@gmail.co','$2a$10$8u3nRFsU2lL192Y8/v8H8Of/yoBNZxiWZCca70IfQeV6EnvRDr0ni','081212345678','peneliti','Udinus','jl. Imam Bonjol','13440_l',NULL,NULL,NULL,1,0),(8,'admin 01','admin01@gmail.com','$2y$04$pmbmfq/R4x7t/qAJj1x2z.Tgjx5f4v91Alo3eqnVG2EVT/19KpCpi','12345678','penyunting','Udinus','jl. Imam Bonjol','13440','BSI','7257491159',NULL,1,0),(9,'Orang Gabut','bagasa11.14715@gmail','$2a$10$Cr0UXNHjXWbJiguK.COTOOzigtdkUyFmuMHdBLkUBtH8Q8mV3SXQq','+6281213244567','peneliti','Udinus','jl. Imam Bonjol','13454',NULL,NULL,NULL,1,0);
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

-- Dump completed on 2024-07-23 21:37:57
