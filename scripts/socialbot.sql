CREATE DATABASE  IF NOT EXISTS `socialbot` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `socialbot`;
-- MySQL dump 10.13  Distrib 8.0.16, for Win64 (x86_64)
--
-- Host: 149.28.210.141    Database: socialbot
-- ------------------------------------------------------
-- Server version	8.0.17

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `s_admin`
--

DROP TABLE IF EXISTS `s_admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `s_admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(65) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(65) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `s_article`
--

DROP TABLE IF EXISTS `s_article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `s_article` (
  `id` bigint(20) NOT NULL,
  `mid` bigint(20) unsigned NOT NULL DEFAULT '0',
  `content` varchar(10000) COLLATE utf8mb4_unicode_ci NOT NULL,
  `create_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `update_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_mid` (`mid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `s_category`
--

DROP TABLE IF EXISTS `s_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `s_category` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `cover` varchar(145) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` varchar(85) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `short_name` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `sort` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `create_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `update_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `s_commission_product`
--

DROP TABLE IF EXISTS `s_commission_product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `s_commission_product` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `mid` bigint(20) unsigned NOT NULL DEFAULT '0',
  `promote_link` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `detail_link` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `content` varchar(10000) COLLATE utf8mb4_unicode_ci NOT NULL,
  `create_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `update_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_mid` (`mid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `s_config`
--

DROP TABLE IF EXISTS `s_config`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `s_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(145) COLLATE utf8mb4_unicode_ci NOT NULL,
  `key_mark` varchar(65) COLLATE utf8mb4_unicode_ci NOT NULL,
  `value` varchar(5000) COLLATE utf8mb4_unicode_ci NOT NULL,
  `create_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `update_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_key_mark` (`key_mark`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `s_copywriter`
--

DROP TABLE IF EXISTS `s_copywriter`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `s_copywriter` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `description` varchar(800) COLLATE utf8mb4_unicode_ci NOT NULL,
  `create_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `update_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_title` (`title`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `s_crawler`
--

DROP TABLE IF EXISTS `s_crawler`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `s_crawler` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(65) COLLATE utf8mb4_unicode_ci NOT NULL,
  `last_page` varchar(65) COLLATE utf8mb4_unicode_ci NOT NULL,
  `script` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `crawler_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT 'o created  1 running  2 failed   3 stop ',
  `create_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `update_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `s_crawler_item`
--

DROP TABLE IF EXISTS `s_crawler_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `s_crawler_item` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uri` varchar(65) COLLATE utf8mb4_unicode_ci NOT NULL,
  `crwid` int(10) unsigned NOT NULL DEFAULT '0',
  `title` varchar(300) COLLATE utf8mb4_unicode_ci NOT NULL,
  `cover` varchar(245) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL,
  `content_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0 imglist split with ,',
  `content` varchar(10000) COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0 created   10 published',
  `create_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `update_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_crwid` (`crwid`),
  KEY `idx_uri` (`uri`)
) ENGINE=InnoDB AUTO_INCREMENT=335574 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `s_media`
--

DROP TABLE IF EXISTS `s_media`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `s_media` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) unsigned NOT NULL DEFAULT '0',
  `cid` int(10) unsigned NOT NULL DEFAULT '0',
  `uri` bigint(20) unsigned NOT NULL DEFAULT '0',
  `title` varchar(800) COLLATE utf8mb4_unicode_ci NOT NULL,
  `cover` varchar(245) COLLATE utf8mb4_unicode_ci NOT NULL,
  `recommend` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `comment_num` int(10) unsigned NOT NULL DEFAULT '0',
  `media_num` smallint(5) unsigned NOT NULL DEFAULT '0',
  `view_num` int(10) unsigned NOT NULL DEFAULT '0',
  `like_num` int(10) unsigned NOT NULL DEFAULT '0',
  `media_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '2 social media  3 article   4 commission  product',
  `media_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0 created   10 published',
  `publish_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `create_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `update_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`uid`) /*!80000 INVISIBLE */,
  KEY `idx_cid` (`cid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `s_media_source`
--

DROP TABLE IF EXISTS `s_media_source`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `s_media_source` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `mid` bigint(20) unsigned NOT NULL DEFAULT '0',
  `uri` bigint(20) unsigned NOT NULL DEFAULT '0',
  `url` varchar(245) COLLATE utf8mb4_unicode_ci NOT NULL,
  `cid` int(10) unsigned NOT NULL DEFAULT '0',
  `source_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0 image  1 video  2 music  ',
  `source_from` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0 local ',
  `create_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `update_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_mid` (`mid`),
  KEY `idx_uri` (`uri`),
  KEY `idx_cid` (`cid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `s_media_source_tag`
--

DROP TABLE IF EXISTS `s_media_source_tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `s_media_source_tag` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `msid` bigint(20) unsigned NOT NULL DEFAULT '0',
  `mid` bigint(20) unsigned NOT NULL DEFAULT '0',
  `cid` int(10) unsigned NOT NULL DEFAULT '0',
  `tid` int(10) unsigned NOT NULL DEFAULT '0',
  `create_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `update_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_msid` (`msid`) /*!80000 INVISIBLE */,
  KEY `idx_tid` (`tid`),
  KEY `idx_mid` (`mid`),
  KEY `idx_cid` (`cid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `s_media_tag`
--

DROP TABLE IF EXISTS `s_media_tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `s_media_tag` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `mid` bigint(20) unsigned NOT NULL DEFAULT '0',
  `cid` int(10) unsigned NOT NULL DEFAULT '0',
  `tid` int(10) unsigned NOT NULL DEFAULT '0',
  `media_status` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `media_publish_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `create_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `update_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_mid` (`mid`),
  KEY `idx_tid` (`tid`),
  KEY `idx_cid` (`cid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `s_pinbot`
--

DROP TABLE IF EXISTS `s_pinbot`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `s_pinbot` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sid` int(11) NOT NULL COMMENT 'server_id',
  `name` varchar(65) COLLATE utf8mb4_unicode_ci NOT NULL,
  `config` varchar(10000) COLLATE utf8mb4_unicode_ci NOT NULL,
  `pinbot_status` tinyint(4) NOT NULL COMMENT '0 created  1 running  2 destroyed',
  `create_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `update_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_sid` (`sid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `s_robot_server`
--

DROP TABLE IF EXISTS `s_robot_server`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `s_robot_server` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `subid` varchar(65) COLLATE utf8mb4_unicode_ci NOT NULL,
  `api_key` varchar(105) COLLATE utf8mb4_unicode_ci NOT NULL,
  `full_server_info` varchar(10000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `robot_agent_status` tinyint(4) NOT NULL,
  `update_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `create_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_subid` (`subid`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `s_tag`
--

DROP TABLE IF EXISTS `s_tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `s_tag` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `cid` int(10) unsigned NOT NULL DEFAULT '0',
  `title` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `board_name` varchar(85) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'relation to pinterest  border',
  `description` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `short_name` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `create_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `update_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_cid` (`cid`) /*!80000 INVISIBLE */,
  KEY `idx_title` (`title`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `s_user`
--

DROP TABLE IF EXISTS `s_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `s_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `uri` bigint(20) unsigned NOT NULL DEFAULT '0',
  `nickname` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'nickname',
  `avatar` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `intro` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `register_ip` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0 default  ',
  `identity` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0 common user   10 editor',
  `register_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `create_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `update_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_uri` (`uri`),
  KEY `idx_email` (`email`),
  KEY `idx_register_ip` (`register_ip`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `s_user_like_media`
--

DROP TABLE IF EXISTS `s_user_like_media`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `s_user_like_media` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `mid` bigint(20) unsigned NOT NULL DEFAULT '0',
  `uid` bigint(20) unsigned NOT NULL DEFAULT '0',
  `create_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `update_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_mid` (`mid`) /*!80000 INVISIBLE */,
  KEY `idx_uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `server`
--

DROP TABLE IF EXISTS `server`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `server` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `subid` varchar(65) COLLATE utf8mb4_unicode_ci NOT NULL,
  `apiKey` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `server_full_info` varchar(10000) COLLATE utf8mb4_unicode_ci NOT NULL,
  `robot_agent_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0 created 1 success connect 2 connect faild',
  `robot_agent_port` int(10) unsigned NOT NULL DEFAULT '0',
  `create_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `update_at` bigint(20) unsigned NOT NULL DEFAULT '0',
  `is_del` tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-08-20 14:32:16
