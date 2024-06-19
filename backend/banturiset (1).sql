-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 19, 2024 at 04:27 AM
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
-- Table structure for table `balances`
--

CREATE TABLE `balances` (
  `id` bigint(20) UNSIGNED NOT NULL,
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
  `pengajuan_id` bigint(20) UNSIGNED NOT NULL,
  `peneliti_id` bigint(20) UNSIGNED NOT NULL,
  `balance_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `budget_details`
--

CREATE TABLE `budget_details` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `deskripsi` longtext DEFAULT NULL,
  `tahap` tinyint(3) UNSIGNED NOT NULL DEFAULT 1,
  `cost` float NOT NULL DEFAULT 0,
  `project_id` bigint(20) UNSIGNED DEFAULT NULL,
  `percent` tinyint(3) UNSIGNED NOT NULL DEFAULT 0,
  `start` datetime(3) DEFAULT NULL,
  `end` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `budget_details`
--

INSERT INTO `budget_details` (`id`, `deskripsi`, `tahap`, `cost`, `project_id`, `percent`, `start`, `end`) VALUES
(1, 'Membeli Radar Cuaca: Radar cuaca akan digunakan untuk mengambil data cuaca', 1, 1000000, 6, 0, NULL, NULL),
(2, 'Biaya Observasi Lapangan', 1, 500000, 6, 0, NULL, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `donasis`
--

CREATE TABLE `donasis` (
  `id` varchar(191) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `status` longtext NOT NULL,
  `jml` float NOT NULL,
  `fee` float NOT NULL,
  `method` varchar(191) NOT NULL DEFAULT 'OVO',
  `donatur_id` bigint(20) UNSIGNED DEFAULT NULL,
  `balance_id` bigint(20) UNSIGNED DEFAULT NULL,
  `project_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `donasis`
--

INSERT INTO `donasis` (`id`, `created_at`, `updated_at`, `deleted_at`, `status`, `jml`, `fee`, `method`, `donatur_id`, `balance_id`, `project_id`) VALUES
('invoice-383fg1e', '2024-06-17 22:42:19.780', '2024-06-17 22:42:19.780', NULL, 'PENDING', 25000, 834.75, 'OVO', 1, NULL, 6),
('invoice-5641heb', '2024-06-17 22:49:14.367', '2024-06-17 22:49:14.366', NULL, 'PENDING', 25000, 834.75, 'OVO', 1, NULL, 6),
('invoice-72g3883', '2024-06-17 22:37:21.597', '2024-06-17 22:37:21.596', NULL, 'PENDING', 25000, 834.75, 'OVO', 1, NULL, 6),
('invoice-bf2eh75', '2024-06-17 22:49:18.411', '2024-06-17 22:49:18.411', NULL, 'PENDING', 25000, 834.75, 'OVO', 1, NULL, 6),
('invoice-bfhb26c', '2024-06-17 22:48:28.826', '2024-06-17 22:48:28.825', NULL, 'PENDING', 25000, 834.75, 'OVO', 1, NULL, 6),
('invoice-cf37a3b', '2024-06-17 22:37:15.414', '2024-06-17 22:37:15.413', NULL, 'PENDING', 25000, 834.75, 'OVO', 1, NULL, 6);

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
(4, '00003', 9),
(5, '010101010', 5);

-- --------------------------------------------------------

--
-- Table structure for table `pengajuans`
--

CREATE TABLE `pengajuans` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(20) DEFAULT NULL,
  `desc` longtext DEFAULT NULL,
  `link_wa` varchar(120) DEFAULT NULL,
  `link_panduan` varchar(120) DEFAULT NULL,
  `icon_url` varchar(120) DEFAULT NULL,
  `closed_at` datetime(3) DEFAULT NULL,
  `penyunting_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `pengajuans`
--

INSERT INTO `pengajuans` (`id`, `created_at`, `updated_at`, `deleted_at`, `title`, `desc`, `link_wa`, `link_panduan`, `icon_url`, `closed_at`, `penyunting_id`) VALUES
(2, '2024-06-06 19:26:47.601', '2024-06-06 19:26:47.602', NULL, 'Teknologi Komputer d', 'skema penelitian ini bertujuan untuk meneliti teknoloi Komputer dan AI', 'web.whatsapp.com', '', 'https://developers.xendit.co/api-reference/images/logo.png', '2024-07-09 19:41:53.000', 1);

-- --------------------------------------------------------

--
-- Table structure for table `penyuntings`
--

CREATE TABLE `penyuntings` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `n_ip` varchar(20) DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `penyuntings`
--

INSERT INTO `penyuntings` (`id`, `n_ip`, `user_id`) VALUES
(1, '1234567890123456789', 8);

-- --------------------------------------------------------

--
-- Table structure for table `progresses`
--

CREATE TABLE `progresses` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `file_url` longtext NOT NULL,
  `desc` longtext DEFAULT NULL,
  `tahap` tinyint(3) UNSIGNED NOT NULL DEFAULT 1,
  `status` tinyint(4) NOT NULL DEFAULT 0,
  `pesan_revisi` longtext DEFAULT NULL,
  `project_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `projects`
--

CREATE TABLE `projects` (
  `id` bigint(20) UNSIGNED NOT NULL,
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
  `pengajuan_id` bigint(20) UNSIGNED NOT NULL,
  `peneliti_id` bigint(20) UNSIGNED NOT NULL,
  `pesan_revisi` longtext DEFAULT NULL,
  `fraud` tinyint(1) NOT NULL DEFAULT 0,
  `collected_fund` float NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `projects`
--

INSERT INTO `projects` (`id`, `created_at`, `updated_at`, `deleted_at`, `title`, `desc`, `proposal_url`, `klirens_url`, `fund_until`, `dead_line`, `milestone`, `tkt_level`, `cost`, `status`, `pengajuan_id`, `peneliti_id`, `pesan_revisi`, `fraud`, `collected_fund`) VALUES
(5, '2024-06-08 21:58:06.426', '2024-06-08 21:58:06.427', NULL, 'Prediksi Cuaca dengan Algoritma Naive Bayes', 'Cuaca merupakan aspek penting dalam berbagai kehidupan manusia. Maka memprediksi cuaca dengan akurat sangat penting. Untuk itu perlu adanya model yang dapat memprediksi cuaca dengan akurat', NULL, NULL, '2024-11-05 21:58:06.426', '2025-06-29 10:15:59.000', 4, 9, 13000000, -2, 2, 5, NULL, 0, 0),
(6, '2024-06-16 11:04:34.750', '2024-06-17 20:23:13.916', NULL, 'Weather Forecasting dengan svm', 'Cuaca merupakan aspek penting dalam berbagai kehidupan manusia. Maka memprediksi cuaca dengan akurat sangat penting. Untuk itu perlu adanya model yang dapat memprediksi cuaca dengan akurat', './file/klirens/84gfh62.pdf', './file/klirens/84gfh62.pdf', '2024-07-01 22:07:55.000', '2027-01-01 11:04:34.745', 4, 9, 13000000, 2, 2, 5, NULL, 0, 0),
(10, '2024-06-19 05:58:19.863', '2024-06-19 05:58:19.863', NULL, 'Testing Timestamp', 'Cuaca merupakan aspek penting dalam berbagai kehidupan manusia. Maka memprediksi cuaca dengan akurat sangat penting. Untuk itu perlu adanya model yang dapat memprediksi cuaca dengan akurat', NULL, NULL, '2024-06-19 05:58:19.861', '2028-01-07 05:58:19.860', 4, 9, 13000000, -2, 2, 5, NULL, 0, 0),
(12, '2024-06-19 06:06:14.294', '2024-06-19 06:06:14.294', NULL, 'Testing Timestamp', 'mengukur akurasi perhitungan waktu pada field fund_until', NULL, NULL, '2025-08-25 14:06:14.292', '2028-01-07 06:06:14.291', 4, 9, 13000000, -2, 2, 5, NULL, 0, 0);

-- --------------------------------------------------------

--
-- Table structure for table `tahapans`
--

CREATE TABLE `tahapans` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `cost_percent` tinyint(3) UNSIGNED NOT NULL DEFAULT 0,
  `start` datetime(3) NOT NULL,
  `end` datetime(3) NOT NULL,
  `project_id` bigint(20) UNSIGNED DEFAULT NULL,
  `tahap` tinyint(3) UNSIGNED NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `tahapans`
--

INSERT INTO `tahapans` (`id`, `cost_percent`, `start`, `end`, `project_id`, `tahap`) VALUES
(1, 30, '2024-06-16 19:38:48.123', '2024-07-16 19:38:48.123', 6, 1);

-- --------------------------------------------------------

--
-- Table structure for table `token_lists`
--

CREATE TABLE `token_lists` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `token` varchar(191) NOT NULL,
  `expired_date` datetime(3) DEFAULT NULL,
  `blacklisted` tinyint(1) DEFAULT NULL
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
  `phone` varchar(191) NOT NULL,
  `role` varchar(191) NOT NULL,
  `institute` longtext NOT NULL,
  `institute_addr` longtext NOT NULL,
  `post_code` varchar(7) NOT NULL,
  `bank` longtext DEFAULT NULL,
  `no_rek` longtext DEFAULT NULL,
  `profile_url` longtext DEFAULT NULL,
  `is_verfied` tinyint(1) NOT NULL DEFAULT 0,
  `is_block` tinyint(1) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `f_name`, `email`, `password`, `phone`, `role`, `institute`, `institute_addr`, `post_code`, `bank`, `no_rek`, `profile_url`, `is_verfied`, `is_block`) VALUES
(1, 'Bagas', 'bagasmipa3@gmail.com', '$2a$10$hZTy0WjQieaozWEs1E/S3e/Gn5KDo2m06sTmfL4s4/deMwmdUVgpG', '+6282112897283', 'donatur', 'Udinus', 'jl. Imam Bonjol', '13340', NULL, NULL, NULL, 1, 0),
(5, 'User Suka Spam', 'spamtest123@gmail.co', '$2a$10$8u3nRFsU2lL192Y8/v8H8Of/yoBNZxiWZCca70IfQeV6EnvRDr0ni', '081212345678', 'peneliti', 'Udinus', 'jl. Imam Bonjol', '13440_l', NULL, NULL, NULL, 1, 0),
(8, 'admin 01', 'admin01@gmail.com', '$2a$10$Pwx4jyFAiQ8fr0DAnukAnOhfYE5po4QZyDcdOPDPGYXbi7X52oqvm', '12345678', 'penyunting', 'Udinus', 'jl. Imam Bonjol', '13440', 'BSI', '7257491159', NULL, 1, 0),
(9, 'Orang Gabut', 'bagasa11.14715@gmail', '$2a$10$Cr0UXNHjXWbJiguK.COTOOzigtdkUyFmuMHdBLkUBtH8Q8mV3SXQq', '+6281213244567', 'peneliti', 'Udinus', 'jl. Imam Bonjol', '13454', NULL, NULL, NULL, 1, 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `balances`
--
ALTER TABLE `balances`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_balances_deleted_at` (`deleted_at`),
  ADD KEY `fk_balances_pengajuan` (`pengajuan_id`),
  ADD KEY `fk_balances_peneliti` (`peneliti_id`);

--
-- Indexes for table `budget_details`
--
ALTER TABLE `budget_details`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_projects_budget_details` (`project_id`);

--
-- Indexes for table `donasis`
--
ALTER TABLE `donasis`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_donasis_deleted_at` (`deleted_at`),
  ADD KEY `fk_donasis_donatur` (`donatur_id`),
  ADD KEY `fk_balances_donasi` (`balance_id`),
  ADD KEY `fk_projects_donasi` (`project_id`);

--
-- Indexes for table `donaturs`
--
ALTER TABLE `donaturs`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_users_donatur` (`user_id`);

--
-- Indexes for table `penelitis`
--
ALTER TABLE `penelitis`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_users_peneliti` (`user_id`);

--
-- Indexes for table `pengajuans`
--
ALTER TABLE `pengajuans`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_pengajuans_title` (`title`),
  ADD KEY `idx_pengajuans_deleted_at` (`deleted_at`),
  ADD KEY `fk_penyuntings_pengajuan` (`penyunting_id`);

--
-- Indexes for table `penyuntings`
--
ALTER TABLE `penyuntings`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_users_penyunting` (`user_id`);

--
-- Indexes for table `progresses`
--
ALTER TABLE `progresses`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_projects_progress` (`project_id`);

--
-- Indexes for table `projects`
--
ALTER TABLE `projects`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_projects_deleted_at` (`deleted_at`),
  ADD KEY `fk_pengajuans_project` (`pengajuan_id`),
  ADD KEY `fk_penelitis_project` (`peneliti_id`);

--
-- Indexes for table `tahapans`
--
ALTER TABLE `tahapans`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_projects_tahapan` (`project_id`);

--
-- Indexes for table `token_lists`
--
ALTER TABLE `token_lists`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_token_lists_token` (`token`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_users_email` (`email`),
  ADD UNIQUE KEY `uni_users_phone` (`phone`),
  ADD KEY `idx_users_role` (`role`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `balances`
--
ALTER TABLE `balances`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `budget_details`
--
ALTER TABLE `budget_details`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `donaturs`
--
ALTER TABLE `donaturs`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `penelitis`
--
ALTER TABLE `penelitis`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `pengajuans`
--
ALTER TABLE `pengajuans`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `penyuntings`
--
ALTER TABLE `penyuntings`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `progresses`
--
ALTER TABLE `progresses`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `projects`
--
ALTER TABLE `projects`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `tahapans`
--
ALTER TABLE `tahapans`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `token_lists`
--
ALTER TABLE `token_lists`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `balances`
--
ALTER TABLE `balances`
  ADD CONSTRAINT `fk_balances_peneliti` FOREIGN KEY (`peneliti_id`) REFERENCES `penelitis` (`id`),
  ADD CONSTRAINT `fk_balances_pengajuan` FOREIGN KEY (`pengajuan_id`) REFERENCES `pengajuans` (`id`);

--
-- Constraints for table `budget_details`
--
ALTER TABLE `budget_details`
  ADD CONSTRAINT `fk_projects_budget_details` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`);

--
-- Constraints for table `donasis`
--
ALTER TABLE `donasis`
  ADD CONSTRAINT `fk_balances_donasi` FOREIGN KEY (`balance_id`) REFERENCES `balances` (`id`),
  ADD CONSTRAINT `fk_donasis_donatur` FOREIGN KEY (`donatur_id`) REFERENCES `donaturs` (`id`),
  ADD CONSTRAINT `fk_projects_donasi` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`);

--
-- Constraints for table `donaturs`
--
ALTER TABLE `donaturs`
  ADD CONSTRAINT `fk_users_donatur` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `penelitis`
--
ALTER TABLE `penelitis`
  ADD CONSTRAINT `fk_users_peneliti` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `pengajuans`
--
ALTER TABLE `pengajuans`
  ADD CONSTRAINT `fk_penyuntings_pengajuan` FOREIGN KEY (`penyunting_id`) REFERENCES `penyuntings` (`id`);

--
-- Constraints for table `penyuntings`
--
ALTER TABLE `penyuntings`
  ADD CONSTRAINT `fk_users_penyunting` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `progresses`
--
ALTER TABLE `progresses`
  ADD CONSTRAINT `fk_projects_progress` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`);

--
-- Constraints for table `projects`
--
ALTER TABLE `projects`
  ADD CONSTRAINT `fk_penelitis_project` FOREIGN KEY (`peneliti_id`) REFERENCES `penelitis` (`id`),
  ADD CONSTRAINT `fk_pengajuans_project` FOREIGN KEY (`pengajuan_id`) REFERENCES `pengajuans` (`id`);

--
-- Constraints for table `tahapans`
--
ALTER TABLE `tahapans`
  ADD CONSTRAINT `fk_projects_tahapan` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
