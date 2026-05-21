-- 若 user.password 为 varchar(20)，bcrypt 哈希（约60字符）无法写入，执行本脚本一次即可。
USE `book_borrow`;
ALTER TABLE `user` MODIFY COLUMN `password` VARCHAR(255) NOT NULL COMMENT '密码(bcrypt)';
