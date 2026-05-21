<?php
/**
 * 图书馆管理系统：通过浏览器访问本页，自动创建数据库 book_borrow 及 user / book / borrow 三张表。
 *
 * 使用步骤（XAMPP）：
 * 1. 启动 Apache + MySQL
 * 2. 将本文件复制到 XAMPP 的 htdocs，例如：/Applications/XAMPP/htdocs/setup_book_borrow.php
 * 3. 浏览器打开：http://localhost/setup_book_borrow.php
 *
 * 若 root 有密码，请修改下面 $db_password。
 */

header('Content-Type: text/html; charset=utf-8');

$db_host = '127.0.0.1';
$db_port = 3306;
$db_user = 'root';
$db_password = ''; // XAMPP 默认常为空；有密码则写成 '你的密码'

// ---------- ① 连接 MySQL 服务器（不指定库，才能执行 CREATE DATABASE）----------
$mysqli = @new mysqli($db_host, $db_user, $db_password, '', $db_port);
if ($mysqli->connect_errno) {
    exit('连接 MySQL 失败：' . htmlspecialchars($mysqli->connect_error, ENT_QUOTES, 'UTF-8'));
}
$mysqli->set_charset('utf8mb4');

// ---------- ② 创建数据库 ----------
$sql_create_db = "CREATE DATABASE IF NOT EXISTS `book_borrow`
    DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci";
if (!$mysqli->query($sql_create_db)) {
    exit('创建数据库失败：' . htmlspecialchars($mysqli->error, ENT_QUOTES, 'UTF-8'));
}
if (!$mysqli->select_db('book_borrow')) {
    exit('选择数据库 book_borrow 失败：' . htmlspecialchars($mysqli->error, ENT_QUOTES, 'UTF-8'));
}

// ---------- ③ 建表（先删后建，与 create.sql 一致）----------
$statements = [
    "DROP TABLE IF EXISTS `borrow`",
    "DROP TABLE IF EXISTS `book`",
    "DROP TABLE IF EXISTS `user`",

    "CREATE TABLE `user` (
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
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户'",

    "CREATE TABLE `book` (
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
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='图书'",

    "CREATE TABLE `borrow` (
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
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='借阅记录'",
];

foreach ($statements as $sql) {
    if (!$mysqli->query($sql)) {
        exit('执行 SQL 失败：<pre>' . htmlspecialchars($sql, ENT_QUOTES, 'UTF-8') . '</pre>错误：' . htmlspecialchars($mysqli->error, ENT_QUOTES, 'UTF-8'));
    }
}

$mysqli->close();
echo '<p>完成：数据库 <strong>book_borrow</strong> 及表 <strong>user</strong>、<strong>book</strong>、<strong>borrow</strong> 已就绪。</p>';
echo '<p>Go 项目请在 <code>manifest/config/config.yaml</code> 中配置相同的库名与账号（见 README 或实验说明）。</p>';
