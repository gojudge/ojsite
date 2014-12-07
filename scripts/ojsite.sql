-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.6.21 - MySQL Community Server (GPL)
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  9.1.0.4867
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- 导出 goj 的数据库结构
CREATE DATABASE IF NOT EXISTS `goj` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `goj`;


-- 导出  表 goj.oauth 结构
CREATE TABLE IF NOT EXISTS `oauth` (
  `id` int(11) NOT NULL,
  `uid` int(11) NOT NULL,
  `type` enum('github') NOT NULL,
  `token` varchar(128) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `token_unique` (`type`,`token`),
  UNIQUE KEY `uid_unique` (`uid`,`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='oauth第三方登录信息';

-- 数据导出被取消选择。


-- 导出  表 goj.problem 结构
CREATE TABLE IF NOT EXISTS `problem` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `type` enum('assert','io') NOT NULL DEFAULT 'assert' COMMENT '题目类型',
  `description` text COMMENT '题目描述',
  `pre_code` text COMMENT '预代码[assert]',
  `io_data` text COMMENT '输入输出样本，json类型[io]',
  `tags` varchar(256) DEFAULT NULL COMMENT '分类标签，逗号分割',
  `level` enum('public','private') NOT NULL DEFAULT 'public' COMMENT '题目类型',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='题库表';

-- 数据导出被取消选择。


-- 导出  表 goj.student 结构
CREATE TABLE IF NOT EXISTS `student` (
  `id` int(11) NOT NULL COMMENT 'ID',
  `userid` int(11) NOT NULL COMMENT 'User ID',
  `student_number` int(10) NOT NULL COMMENT '学号',
  `class` varchar(50) NOT NULL COMMENT '班级',
  PRIMARY KEY (`id`),
  UNIQUE KEY `userid` (`userid`),
  UNIQUE KEY `student_number` (`student_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='学生表';

-- 数据导出被取消选择。


-- 导出  表 goj.user 结构
CREATE TABLE IF NOT EXISTS `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `username` varchar(50) NOT NULL COMMENT '用户名',
  `password` varchar(128) NOT NULL COMMENT '密码',
  `salt` varchar(8) NOT NULL COMMENT '盐',
  `email` varchar(128) NOT NULL COMMENT 'Email',
  `level` enum('user','student','teacher','admin') NOT NULL DEFAULT 'user' COMMENT '用户级别',
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
