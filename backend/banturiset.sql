-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 03, 2024 at 06:24 AM
-- Server version: 10.4.27-MariaDB
-- PHP Version: 8.1.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `banturiset`
--

-- --------------------------------------------------------

--
-- Table structure for table `donaturs`
--

CREATE TABLE `donaturs` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `donaturs`
--

INSERT INTO `donaturs` (`id`, `user_id`) VALUES
(1, 1);

-- --------------------------------------------------------

--
-- Table structure for table `kategoris`
--

CREATE TABLE `kategoris` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(20) DEFAULT NULL,
  `desc` longtext DEFAULT NULL,
  `link` varchar(120) DEFAULT NULL,
  `penyunting_id` bigint(20) UNSIGNED DEFAULT NULL,
  `n_ip` varchar(20) DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `penelitis`
--

CREATE TABLE `penelitis` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `n_ip` varchar(20) DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `penelitis`
--

INSERT INTO `penelitis` (`id`, `n_ip`, `user_id`) VALUES
(1, '12345', 3),
(2, '12345', 4),
(3, '00001', 5);

-- --------------------------------------------------------

--
-- Table structure for table `penyuntings`
--

CREATE TABLE `penyuntings` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `n_ip` varchar(20) DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `f_name` longtext NOT NULL,
  `email` varchar(20) NOT NULL,
  `password` longtext NOT NULL,
  `phone` longtext NOT NULL,
  `role` varchar(191) NOT NULL,
  `institute` longtext NOT NULL,
  `institute_addr` longtext NOT NULL,
  `post_code` varchar(7) NOT NULL,
  `bank` longtext DEFAULT NULL,
  `no_rek` longtext DEFAULT NULL,
  `profile_url` longtext DEFAULT NULL,
  `is_verfied` tinyint(1) NOT NULL DEFAULT 0,
  `isb_block` tinyint(1) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `f_name`, `email`, `password`, `phone`, `role`, `institute`, `institute_addr`, `post_code`, `bank`, `no_rek`, `profile_url`, `is_verfied`, `isb_block`) VALUES
(1, 'Bagas', 'bagasmipa3@gmail.com', '$2a$10$hZTy0WjQieaozWEs1E/S3e/Gn5KDo2m06sTmfL4s4/deMwmdUVgpG', '12345678', 'donatur', 'Udinus', 'jl. Imam Bonjol', '13340', NULL, NULL, NULL, 0, 0),
(2, 'Bagas RS', 'bagas123@gmail.com', '$2a$10$Pwx4jyFAiQ8fr0DAnukAnOhfYE5po4QZyDcdOPDPGYXbi7X52oqvm', '12345678', 'donatur', 'Udinus', 'jl. Imam Bonjol', '13-34x0', NULL, NULL, NULL, 0, 0),
(3, 'Bagas RS', 'bagas321@gmail.com', '$2a$10$QbZuzkbsZheUcNnhVrKCWuofVJLmU8atCOY/.DvLCpZWEglwYeQuq', '12345678', 'peneliti', 'Udinus', 'jl. Imam Bonjol', '13-34x0', NULL, NULL, NULL, 0, 0),
(4, 'Bagas Rayhan', 'bagas11@gmail.com', '$2a$10$6bSv5oJU40Umtpf0VZ7m/uF5mM15hyWbYwC0qY6AepYKBGi11jJse', '12345678', 'peneliti', 'Udinus', 'jl. Imam Bonjol', '13440_l', NULL, NULL, NULL, 0, 0),
(5, 'User Suka Spam', 'spamtest123@gmail.co', '$2a$10$8u3nRFsU2lL192Y8/v8H8Of/yoBNZxiWZCca70IfQeV6EnvRDr0ni', '12345678', 'peneliti', 'Udinus', 'jl. Imam Bonjol', '13440_l', NULL, NULL, NULL, 0, 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `donaturs`
--
ALTER TABLE `donaturs`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_users_donatur` (`user_id`);

--
-- Indexes for table `kategoris`
--
ALTER TABLE `kategoris`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_kategoris_title` (`title`),
  ADD KEY `idx_kategoris_deleted_at` (`deleted_at`),
  ADD KEY `fk_penyuntings_kategori` (`penyunting_id`);

--
-- Indexes for table `penelitis`
--
ALTER TABLE `penelitis`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_users_peneliti` (`user_id`);

--
-- Indexes for table `penyuntings`
--
ALTER TABLE `penyuntings`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_users_penyunting` (`user_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_users_email` (`email`),
  ADD KEY `idx_users_role` (`role`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `donaturs`
--
ALTER TABLE `donaturs`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `kategoris`
--
ALTER TABLE `kategoris`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `penelitis`
--
ALTER TABLE `penelitis`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `penyuntings`
--
ALTER TABLE `penyuntings`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `donaturs`
--
ALTER TABLE `donaturs`
  ADD CONSTRAINT `fk_users_donatur` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `kategoris`
--
ALTER TABLE `kategoris`
  ADD CONSTRAINT `fk_kategoris_kategori` FOREIGN KEY (`penyunting_id`) REFERENCES `kategoris` (`id`),
  ADD CONSTRAINT `fk_penyuntings_kategori` FOREIGN KEY (`penyunting_id`) REFERENCES `penyuntings` (`id`);

--
-- Constraints for table `penelitis`
--
ALTER TABLE `penelitis`
  ADD CONSTRAINT `fk_users_peneliti` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `penyuntings`
--
ALTER TABLE `penyuntings`
  ADD CONSTRAINT `fk_users_penyunting` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
