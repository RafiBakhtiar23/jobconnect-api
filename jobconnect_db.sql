-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jun 24, 2026 at 06:33 AM
-- Server version: 8.0.30
-- PHP Version: 8.3.29

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `jobconnect_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `jobs`
--

CREATE TABLE `jobs` (
  `id` bigint UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `judul` longtext,
  `deskripsi` longtext,
  `company` longtext,
  `lokasi` longtext,
  `gaji` bigint DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `jobs`
--

INSERT INTO `jobs` (`id`, `created_at`, `updated_at`, `deleted_at`, `judul`, `deskripsi`, `company`, `lokasi`, `gaji`) VALUES
(7, '2026-06-16 22:27:13.309', '2026-06-17 21:38:39.614', NULL, 'Senior Golang Developer', 'Update Lowongan', 'PT Teknologi Maju', 'Bandung', 15000000),
(11, '2026-06-16 22:31:57.465', '2026-06-16 22:31:57.465', NULL, 'Backend Developer', 'Membuat API', 'PT ABC', 'Jakarta', 10000000),
(12, '2026-06-16 22:33:22.274', '2026-06-16 22:33:22.274', NULL, 'Backend Developer', 'Membuat API', 'PT ABC', 'Jakarta', 10000000),
(13, '2026-06-16 22:33:52.585', '2026-06-16 22:33:52.585', NULL, '', '', '', '', 0),
(14, '2026-06-17 21:03:33.617', '2026-06-17 21:06:52.381', '2026-06-17 21:08:34.523', 'Senior Golang Developer', 'Update Lowongan', 'PT Teknologi Maju', 'Bandung', 15000000),
(15, '2026-06-17 21:37:28.050', '2026-06-17 21:40:28.981', '2026-06-17 21:41:32.955', 'Senior Golang Developer', 'Update Lowongan', 'PT Teknologi Maju', 'Bandung', 15000000);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `email` longtext,
  `password` longtext,
  `role` longtext
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `email`, `password`, `role`) VALUES
(5, '2026-06-16 21:50:25.831', '2026-06-16 21:50:25.831', NULL, 'Rafi Bakhtiar Cahyadi', '', '$2a$10$iLXZjTu646FHW645BxKnDORJdf29eE5IHkyHt.DM2D7n1DjysjbGK', 'user'),
(6, '2026-06-16 21:52:17.802', '2026-06-16 21:52:17.802', NULL, 'Rafi Bakhtiar C', 'raficoba2@gmail.com', '$2a$10$EshRRmmljuit1dSFTGq6rOzeKDbYJmOpKG7vVpLy1uoWXA6FLCdtu', 'admin\r\n'),
(7, '2026-06-16 22:28:24.455', '2026-06-16 22:28:24.455', NULL, 'user', 'user@gmail.com', '$2a$10$RJtdM.xa/29aesx/BrmLPuytkMSKPtCmqLYb0XOGIqjj4tyLn.BXi', 'user'),
(8, '2026-06-17 21:33:26.392', '2026-06-17 21:33:26.392', NULL, '', 'testuser@gmail.com', '$2a$10$EvmZssOo33iEljEQJuqyVumHP685dGT4XzIpALD.j7yfp7D5sbdfK', 'user');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `jobs`
--
ALTER TABLE `jobs`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_jobs_deleted_at` (`deleted_at`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_users_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `jobs`
--
ALTER TABLE `jobs`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
