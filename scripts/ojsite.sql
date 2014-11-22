-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.6.21 - MySQL Community Server (GPL)
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  8.3.0.4803
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- 导出 ojsite 的数据库结构
CREATE DATABASE IF NOT EXISTS `ojsite` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `ojsite`;


-- 导出  表 ojsite.problem 结构
CREATE TABLE IF NOT EXISTS `problem` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `type` enum('assert','io') NOT NULL DEFAULT 'assert' COMMENT '题目类型',
  `description` text COMMENT '题目描述',
  `pre_code` text COMMENT '预代码[assert]',
  `io_data` text COMMENT '输入输出样本，json类型[io]',
  `tags` varchar(256) DEFAULT NULL COMMENT '分类标签，逗号分割',
  `level` enum('public','private') NOT NULL DEFAULT 'public' COMMENT '题目类型',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='题库表';

-- 数据导出被取消选择。


-- 导出  表 ojsite.user 结构
CREATE TABLE IF NOT EXISTS `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `username` varchar(50) NOT NULL COMMENT '用户名',
  `password` varchar(128) NOT NULL COMMENT '密码',
  `salt` varchar(8) NOT NULL COMMENT '盐',
  `email` varchar(128) NOT NULL COMMENT 'Email',
  `level` enum('user','student','admin') NOT NULL DEFAULT 'user' COMMENT '用户级别',
  `registor_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  `nickname` varchar(50) NOT NULL COMMENT '昵称',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `nickname` (`nickname`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

-- 数据导出被取消选择。
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
