-- CREATE TABLE `user`(
-- 	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
--     `name` varchar(255) NOT NULL DEFAULT '' COMMENT 'user name',
--     `gender` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT 'user gender',
--     `mobile` varchar(255) NOT NULL DEFAULT '' COMMENT 'user mobile number',
--     `password` varchar(255) NOT NULL DEFAULT '' COMMENT 'user password',
--     `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
--     `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP 
-- 		ON UPDATE CURRENT_TIMESTAMP, # update as current time 
-- 	PRIMARY KEY(`id`),
--     UNIQUE KEY `idx_mobile_unique` (`moblie`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE DATABASE IF NOT EXISTS `movie`;

DROP TABLE IF EXISTS `movie`.`users`;
CREATE TABLE `movie`.`users`(
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL DEFAULT '' COMMENT 'user name',
    `email` varchar(255) NOT NULL DEFAULT '' COMMENT 'user email',
    `password` varchar(255) NOT NULL DEFAULT '' COMMENT 'user password',
    `avatar` varchar(255) NULL DEFAULT 'default.jpg' COMMENT 'user avatar',
    `create_time` timestamp DEFAULT current_timestamp,
    `update_time` timestamp DEFAULT current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY(`id`),
    UNIQUE `idx_email_unique`(`email`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for genre_infos
-- ----------------------------
DROP TABLE IF EXISTS `movie`.`genre_infos`;
CREATE TABLE `movie`.`genre_infos`  (
  `genre_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `created_at` timestamp default current_timestamp,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`genre_id`) USING BTREE,
  INDEX `idx_genre_infos_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for movie_infos
-- ----------------------------
DROP TABLE IF EXISTS `movie`.`movie_infos`;
CREATE TABLE `movie`.`movie_infos`  (
  `adult` tinyint(1) NOT NULL,
  `backdrop_path` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `movie_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `original_language` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `original_title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `overview` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `popularity` double NOT NULL,
  `poster_path` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `release_date` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `run_time` bigint NOT NULL,
  `video` tinyint(1) NOT NULL,
  `vote_average` double NOT NULL,
  `vote_count` bigint NOT NULL,
  PRIMARY KEY (`movie_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for person_infos
-- ----------------------------
DROP TABLE IF EXISTS `movie`.`person_infos`;
CREATE TABLE `movie`.`person_infos`  (
  `adult` tinyint(1) NOT NULL,
  `gender` bigint NOT NULL,
  `person_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `department` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `popularity` double NOT NULL,
  `profile_path` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`person_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for genres_movies
-- ----------------------------
DROP TABLE IF EXISTS `movie`.`genres_movies`;
CREATE TABLE `movie`.`genres_movies`  (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`genre_info_genre_id` bigint UNSIGNED NOT NULL,
	`movie_info_movie_id` bigint UNSIGNED NOT NULL,
--   PRIMARY KEY (`movie_info_id`, `genre_info_id`) USING BTREE,
--   INDEX `fk_genres_movies_genre_info`(`genre_info_id` ASC) USING BTREE,
	FOREIGN KEY (`genre_info_genre_id`) REFERENCES `movie`.`genre_infos` (`genre_id`) ON DELETE cascade ON UPDATE cascade,
	FOREIGN KEY (`movie_info_movie_id`) REFERENCES `movie`.`movie_infos` (`movie_id`) ON DELETE cascade ON UPDATE cascade,
	UNIQUE KEY( `genre_info_genre_id`, `movie_info_movie_id`),
    PRIMARY KEY(`id`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for movie_characters
-- ----------------------------
DROP TABLE IF EXISTS `movie`.`movie_characters`;
CREATE TABLE `movie`.`movie_characters`  (
  `person_id` bigint UNSIGNED NOT NULL,
  `movie_id` bigint NOT NULL,
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `character` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `credit_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `order` bigint NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_person_infos_movie_character`(`person_id` ASC) USING BTREE,
  CONSTRAINT `fk_person_infos_movie_character` FOREIGN KEY (`person_id`) REFERENCES `movie`.`person_infos` (`person_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for movie_video_infos
-- ----------------------------
DROP TABLE IF EXISTS `movie`.`movie_video_infos`;
CREATE TABLE `movie`.`movie_video_infos`  (
  `movie_id` bigint UNSIGNED NOT NULL,
  `file_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `trailer_name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `release_time` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`movie_id`, `file_path`) USING BTREE,
  CONSTRAINT `fk_movie_infos_movie_video` FOREIGN KEY (`movie_id`) REFERENCES `movie`.`movie_infos` (`movie_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for person_crews
-- ----------------------------
DROP TABLE IF EXISTS `movie`.`person_crews`;
CREATE TABLE `movie`.`person_crews`  (
  `person_id` bigint UNSIGNED NOT NULL,
  `movie_id` bigint NOT NULL,
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `credit_id` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `department` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_person_infos_person_crew`(`person_id` ASC) USING BTREE,
  CONSTRAINT `fk_person_infos_person_crew` FOREIGN KEY (`person_id`) REFERENCES `movie`.`person_infos` (`person_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

DROP TABLE IF EXISTS `movie`.`articles`;
CREATE TABLE `movie`.`articles`(
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `article_title` varchar(255) NOT NULL,
    `user_id` bigint unsigned NOT NULL,
    `movie_id` bigint unsigned NOT NULL,
    `article_like_count` int NOT NULL DEFAULT 0,
	`create_time` timestamp DEFAULT current_timestamp,
    `update_time` timestamp DEFAULT current_timestamp ON UPDATE current_timestamp,
	PRIMARY KEY(`id`),
	FOREIGN KEY(`movie_id`) REFERENCES `movie`.`movie_infos`(`movie_id`) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY(`user_id`) REFERENCES `movie`.`users`(`id`) ON UPDATE CASCADE ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `movie`.`lists`;
CREATE TABLE `movie`.`lists`(
	`id` bigint unsigned NOT NULL AUTO_INCREMENT ,
    `list_title` varchar(255) NOT NULL,
    `user_id` bigint unsigned NOT NULL,
    `create_time` timestamp DEFAULT current_timestamp,
    `update_time` timestamp DEFAULT current_timestamp ON UPDATE current_timestamp,
--     `list_last_update` timestamp on update current_timestamp,
	PRIMARY KEY(`id`),
    FOREIGN KEY(`user_id`) REFERENCES `movie`.`users`(`id`) ON UPDATE CASCADE ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `movie`.`lists_movies`;
CREATE TABLE `movie`.`lists_movies`(
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `list_id` bigint unsigned NOT NULL,
    `movie_id` bigint unsigned NOT NULL,
    `movie_poster_path` varchar(255) DEFAULT '',
    `user_feeling` varchar(255) DEFAULT '',
    `user_ratetext` varchar(255) DEFAULT '',
    `create_time` timestamp DEFAULT current_timestamp,
    `update_time` timestamp DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
	FOREIGN KEY(`list_id`) REFERENCES `movie`.`lists`(`id`) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY(`movie_id`) REFERENCES `movie`.`movie_infos`(`movie_id`) ON UPDATE CASCADE ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `movie`.`like_articles`;
CREATE TABLE `movie`.`like_articles`(
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id` bigint unsigned NOT NULL,
    `article_id` bigint unsigned NOT NULL,
	`update_time` timestamp DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    FOREIGN KEY(`user_id`) REFERENCES `movie`.`users`(`id`) ON UPDATE CASCADE ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `movie`.`like_movies`;
CREATE TABLE `movie`.`like_movies`(
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id` bigint unsigned NOT NULL,
    `movie_id` bigint unsigned NOT NULL,
    `movie_ti
    tle` varchar(255) NOT NULL DEFAULT '',
    `movie_poster_path` varchar(255) DEFAULT '',
	`update_time` timestamp DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
	FOREIGN KEY(`user_id`) REFERENCES `movie`.`users`(`id`) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY(`movie_id`) REFERENCES `movie`.`movie_infos`(`movie_id`) ON UPDATE CASCADE ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `movie`.`comments`;
CREATE TABLE `movie`.`comments`(
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `comment_text` varchar(255) NOT NULL DEFAULT '',
    `user_id` bigint unsigned NOT NULL,
    `article_id` bigint unsigned NOT NULL,
    `comment_like_count` int NOT NULL DEFAULT 0,
    `create_time` timestamp DEFAULT current_timestamp,
    `update_time` timestamp DEFAULT current_timestamp ON UPDATE current_timestamp,
	PRIMARY KEY(`id`),
	FOREIGN KEY(`user_id`) REFERENCES `movie`.`users`(`id`) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY(`user_id`) REFERENCES `movie`.`articles`(`id`) ON UPDATE CASCADE ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
