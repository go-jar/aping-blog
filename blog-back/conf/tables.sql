CREATE DATABASE IF NOT EXISTS blog DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

use blog;
set names UTF8;

CREATE TABLE IF NOT EXISTS `id_gen` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `max_id` bigint(20) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
    `password` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
    `role` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
    `created_time` datetime NOT NULL,
    `updated_time` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    UNIQUE KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

CREATE TABLE IF NOT EXISTS `category` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `category_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
    `category_index` int(11) NOT NULL DEFAULT 0,
    `created_time` datetime NOT NULL,
    `updated_time` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    UNIQUE KEY (`category_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

INSERT INTO id_gen (name, max_id) VALUES ('category', 0);

CREATE TABLE IF NOT EXISTS `tag` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `tag_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
    `tag_index` int(11) NOT NULL DEFAULT 0,
    `created_time` datetime NOT NULL,
    `updated_time` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    UNIQUE KEY (`tag_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

INSERT INTO id_gen (name, max_id) VALUES ('tag', 0);

CREATE TABLE IF NOT EXISTS `article` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
    `category_id` bigint(20) NOT NULL DEFAULT 0,
    `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
    `read_count` bigint(20) NOT NULL DEFAULT 0,
    `created_time` datetime NOT NULL,
    `updated_time` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

INSERT INTO id_gen (name, max_id) VALUES ('article', 0);

CREATE TABLE IF NOT EXISTS `article_tag` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `article_id` bigint(20) NOT NULL DEFAULT 0,
    `tag_id` bigint(20) NOT NULL DEFAULT 0,
    `created_time` datetime NOT NULL,
    `updated_time` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

INSERT INTO id_gen (name, max_id) VALUES ('article_tag', 0);

INSERT INTO id_gen (name, max_id) VALUES ('cos', 0);

CREATE TABLE IF NOT EXISTS `remark` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `article_id` bigint(20) NOT NULL DEFAULT 0,
    `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
    `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
    `init_remark_id` bigint(20) NOT NULL DEFAULT 0,
    `nickname_replied` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
    `created_time` datetime NOT NULL,
    `updated_time` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

INSERT INTO id_gen (name, max_id) VALUES ('remark', 1);
