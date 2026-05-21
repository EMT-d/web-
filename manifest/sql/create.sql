-- 图书馆管理系统：在 XAMPP phpMyAdmin 或 mysql 客户端中执行本脚本
-- 数据库名与 manifest/config/config.yaml 中 link 的路径一致（默认 book_borrow）

CREATE DATABASE IF NOT EXISTS `book_borrow` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `book_borrow`;

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `borrow`;
DROP TABLE IF EXISTS `book`;
DROP TABLE IF EXISTS `user`;

-- 用户表（注册 / 登录抓包实验）
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户主键ID',
  `student_no` varchar(32) NOT NULL COMMENT '学号',
  `username` varchar(64) NOT NULL COMMENT '学生姓名',
  `password` varchar(255) NOT NULL COMMENT '密码(bcrypt)',
  `phone` varchar(32) DEFAULT NULL COMMENT '手机号',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_student_no` (`student_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户';

-- 图书表
CREATE TABLE `book` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '图书主键ID',
  `isbn` varchar(32) NOT NULL COMMENT 'ISBN',
  `book_name` varchar(255) NOT NULL COMMENT '图书名称',
  `author` varchar(128) NOT NULL COMMENT '作者',
  `publisher` varchar(128) DEFAULT NULL COMMENT '出版社',
  `stock` int unsigned NOT NULL DEFAULT '0' COMMENT '可借库存',
  `total` int unsigned NOT NULL DEFAULT '0' COMMENT '馆藏总量',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_isbn` (`isbn`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='图书';

-- 借阅表（实验要求表名 borrow）
CREATE TABLE `borrow` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '借阅记录主键ID',
  `user_id` bigint unsigned NOT NULL COMMENT '借阅用户ID',
  `book_id` bigint unsigned NOT NULL COMMENT '借阅图书ID',
  `borrow_time` datetime DEFAULT NULL COMMENT '借阅时间',
  `expect_return_time` datetime DEFAULT NULL COMMENT '预计归还时间',
  `actual_return_time` datetime DEFAULT NULL COMMENT '实际归还时间',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '1借阅中 2已归还 3超期',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user` (`user_id`),
  KEY `idx_book` (`book_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='借阅记录';
