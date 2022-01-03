-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.5.5-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             11.0.0.5919
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Dumping database structure for shorturl
DROP DATABASE IF EXISTS `shorturl`;
CREATE DATABASE IF NOT EXISTS `shorturl` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `shorturl`;

-- Dumping structure for procedure shorturl.sp_AddShortURL
DROP PROCEDURE IF EXISTS `sp_AddShortURL`;
DELIMITER //
CREATE PROCEDURE `sp_AddShortURL`(
	IN `shortenedURL` varchar(1000),
	IN `idd` integer
)
BEGIN
	update tbl_urls set ShortURL=shortenedURL where Id = idd;
END//
DELIMITER ;

-- Dumping structure for procedure shorturl.sp_AddURLAndGetID
DROP PROCEDURE IF EXISTS `sp_AddURLAndGetID`;
DELIMITER //
CREATE PROCEDURE `sp_AddURLAndGetID`(
	IN `url` varchar(1000)
)
BEGIN

	insert into tbl_urls(OriginalURL) values (url);
	select Id from tbl_links where OriginalURL = url;

END//
DELIMITER ;

-- Dumping structure for procedure shorturl.sp_GetURL
DROP PROCEDURE IF EXISTS `sp_GetURL`;
DELIMITER //
CREATE PROCEDURE `sp_GetURL`(
	IN `shortlink` varchar(100)
)
BEGIN

	select OriginalURL from tbl_urls where ShortURL = shortlink;

END//
DELIMITER ;

-- Dumping structure for table shorturl.tbl_urls
DROP TABLE IF EXISTS `tbl_urls`;
CREATE TABLE IF NOT EXISTS `tbl_urls` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `ShortURL` varchar(50) DEFAULT NULL,
  `OriginalURL` varchar(1000) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=10005 DEFAULT CHARSET=utf8;