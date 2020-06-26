-- phpMyAdmin SQL Dump
-- version 4.9.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 26, 2020 at 10:35 AM
-- Server version: 10.4.10-MariaDB
-- PHP Version: 7.3.12
CREATE DATABASE /*!32312 IF NOT EXISTS*/`db_testing` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `db_testing`;

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_testing`
--

-- --------------------------------------------------------

--
-- Table structure for table `publisher`
--

DROP TABLE IF EXISTS `publisher`;

CREATE TABLE `publisher` (
  `publisher_id` int(11) NOT NULL,
  `publisher_name` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `input_date` date DEFAULT NULL,
  `last_update` date DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Dumping data for table `publisher`
--

INSERT INTO `publisher` (`publisher_id`, `publisher_name`, `input_date`, `last_update`) VALUES
(1111, 'Scarlet', '2020-11-11', '2020-12-12'),
(1112, 'Altair ', '2020-10-10', '2020-10-01'),
(1113, 'Verlina', '2020-09-10', '2020-11-01');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `publisher`
--
ALTER TABLE `publisher`
  ADD PRIMARY KEY (`publisher_id`),
  ADD UNIQUE KEY `publisher_name` (`publisher_name`),
  ADD KEY `publisher_name_2` (`publisher_name`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `publisher`
--
ALTER TABLE `publisher`
  MODIFY `publisher_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1114;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
