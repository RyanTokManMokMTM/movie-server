CREATE DATABASE IF NOT EXISTS `movie`;

DROP TABLE IF EXISTS `movie`.`users`;
CREATE TABLE `movie`.`users`(
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL DEFAULT '' COMMENT 'user name',
    `email` varchar(255) NOT NULL DEFAULT '' COMMENT 'user email',
    `password` varchar(255) NOT NULL DEFAULT '' COMMENT 'user password',
    `avatar` varchar(255) NULL DEFAULT 'default.jpg' COMMENT 'user avatar',
    `cover` varchar(255) NULL DEFAULT  'cover.jpg' COMMENT 'user background cover',
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
  `name` varchar(255) CHARACTER SET utf8mb4 NOT NULL,
  `created_at` timestamp default current_timestamp,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`genre_id`) USING BTREE,
  INDEX `idx_genre_infos_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 ;

-- ----------------------------
-- Table structure for movie_infos
-- ----------------------------
DROP TABLE IF EXISTS `movie`.`movie_infos`;
CREATE TABLE `movie`.`movie_infos`  (
  `adult` tinyint(1) NOT NULL,
  `backdrop_path` varchar(255) CHARACTER SET utf8mb4 NOT NULL,
  `movie_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `original_language` varchar(255)  CHARACTER SET utf8mb4  NOT NULL,
  `original_title` varchar(255) CHARACTER SET utf8mb4 NOT NULL,
  `overview` longtext  CHARACTER SET utf8mb4 NOT NULL,
  `popularity` double NOT NULL,
  `poster_path` varchar(255)  CHARACTER SET utf8mb4  NOT NULL,
  `release_date` varchar(255) CHARACTER SET utf8mb4  NOT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4  NOT NULL,
  `run_time` bigint NOT NULL,
  `video` tinyint(1) NOT NULL,
  `vote_average` double NOT NULL,
  `vote_count` bigint NOT NULL,
  PRIMARY KEY (`movie_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 ;


-- ----------------------------
-- Table structure for person_infos
-- ----------------------------
DROP TABLE IF EXISTS `movie`.`person_infos`;
CREATE TABLE `movie`.`person_infos`  (
  `adult` tinyint(1) NOT NULL,
  `gender` bigint NOT NULL,
  `person_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `department` varchar(255)  CHARACTER SET utf8mb4  NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4  NOT NULL,
  `popularity` double NOT NULL,
  `profile_path` varchar(255) CHARACTER SET utf8mb4  NOT NULL,
  PRIMARY KEY (`person_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 ;

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
) ENGINE = InnoDB CHARACTER SET = utf8mb4;

-- ----------------------------
-- Table structure for movie_characters
-- ----------------------------
DROP TABLE IF EXISTS `movie`.`movie_characters`;
CREATE TABLE `movie`.`movie_characters`  (
  `person_id` bigint UNSIGNED NOT NULL,
  `movie_id` bigint NOT NULL,
  `movie_character_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `character` varchar(255) CHARACTER SET utf8mb4  NOT NULL,
  `credit_id` varchar(255) CHARACTER SET utf8mb4 NOT NULL,
  `order` bigint NOT NULL,
  PRIMARY KEY (`movie_character_id`) USING BTREE,
  INDEX `fk_person_infos_movie_character`(`person_id` ASC) USING BTREE,
  CONSTRAINT `fk_person_infos_movie_character` FOREIGN KEY (`person_id`) REFERENCES `movie`.`person_infos` (`person_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4;

-- ----------------------------
-- Table structure for movie_video_infos
-- ----------------------------
DROP TABLE IF EXISTS `movie`.`movie_video_infos`;
CREATE TABLE `movie`.`movie_video_infos`  (
  `movie_id` bigint UNSIGNED NOT NULL,
  `file_path` varchar(191) CHARACTER SET utf8mb4 NOT NULL,
  `trailer_name` varchar(255)  CHARACTER SET utf8mb4  NULL,
  `release_time` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`movie_id`, `file_path`) USING BTREE,
  CONSTRAINT `fk_movie_infos_movie_video` FOREIGN KEY (`movie_id`) REFERENCES `movie`.`movie_infos` (`movie_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ;


-- ----------------------------
-- Table structure for person_crews
-- ----------------------------
DROP TABLE IF EXISTS `movie`.`person_crews`;
CREATE TABLE `movie`.`person_crews`  (
  `person_id` bigint UNSIGNED NOT NULL,
  `movie_id` bigint NOT NULL,
  `person_crew_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `credit_id` varchar(255) CHARACTER SET utf8mb4  NOT NULL,
  `department` varchar(255) CHARACTER SET utf8mb4  NOT NULL,
  PRIMARY KEY (`person_crew_id`) USING BTREE,
  INDEX `fk_person_infos_person_crew`(`person_id` ASC) USING BTREE,
  CONSTRAINT `fk_person_infos_person_crew` FOREIGN KEY (`person_id`) REFERENCES `movie`.`person_infos` (`person_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 ;

SET FOREIGN_KEY_CHECKS = 1;

DROP TABLE IF EXISTS `movie`.`posts`;
CREATE TABLE `movie`.`posts`(
		`post_id` bigint unsigned NOT NULL AUTO_INCREMENT comment 'post id',
    `post_title` varchar(255) NOT NULL comment 'post title',
    `post_desc` LONGTEXT NOT NULL comment 'post desc',
    `user_id` bigint unsigned NOT NULL  comment 'who posted the post',
    `movie_id` bigint unsigned NOT NULL  comment 'relevant movie info',
    `post_like` int NOT NULL DEFAULT 0 comment 'post like ',
		`create_time` timestamp DEFAULT current_timestamp,
    `update_time` timestamp DEFAULT current_timestamp ON UPDATE current_timestamp,
	PRIMARY KEY(`post_id`),
	FOREIGN KEY(`movie_id`) REFERENCES `movie`.`movie_infos`(`movie_id`) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY(`user_id`) REFERENCES `movie`.`users`(`id`) ON UPDATE CASCADE ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `movie`.`lists`;
CREATE TABLE `movie`.`lists`(
	`list_id` bigint unsigned NOT NULL AUTO_INCREMENT ,
    `list_title` varchar(255) NOT NULL,
    `user_id` bigint unsigned NOT NULL,
    `create_time` timestamp DEFAULT current_timestamp,
    `update_time` timestamp DEFAULT current_timestamp ON UPDATE current_timestamp,
--     `list_last_update` timestamp on update current_timestamp,
	PRIMARY KEY(`list_id`),
    FOREIGN KEY(`user_id`) REFERENCES `movie`.`users`(`id`) ON UPDATE CASCADE ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `movie`.`lists_movies`;
CREATE TABLE `movie`.`lists_movies`(
	`lists_movie_id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `list_id` bigint unsigned NOT NULL,
    `movie_id` bigint unsigned NOT NULL,
    `movie_poster_path` varchar(255) DEFAULT '',
    `user_feeling` varchar(255) DEFAULT '',
    `user_ratetext` varchar(255) DEFAULT '',
    `create_time` timestamp DEFAULT current_timestamp,
    `update_time` timestamp DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`lists_movie_id`),
	FOREIGN KEY(`list_id`) REFERENCES `movie`.`lists`(`list_id`) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY(`movie_id`) REFERENCES `movie`.`movie_infos`(`movie_id`) ON UPDATE CASCADE ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `movie`.`like_articles`;
CREATE TABLE `movie`.`like_articles`(
	`like_article_id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id` bigint unsigned NOT NULL,
    `article_id` bigint unsigned NOT NULL,
	`update_time` timestamp DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`like_article_id`),
    FOREIGN KEY(`user_id`) REFERENCES `movie`.`users`(`id`) ON UPDATE CASCADE ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `movie`.`liked_movies`;
CREATE TABLE `movie`.`liked_movies`(
	`liked_movie_id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id` bigint unsigned NOT NULL,
    `movie_id` bigint unsigned NOT NULL,
--     `movie_title` varchar(255) NOT NULL DEFAULT '',
--     `movie_poster_path` varchar(255) DEFAULT '',
    `create_time` timestamp DEFAULT current_timestamp ,
	`update_time` timestamp DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`liked_movie_id`),
    UNIQUE KEY( `user_id`, `movie_id`),
	FOREIGN KEY(`user_id`) REFERENCES `movie`.`users`(`id`) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY(`movie_id`) REFERENCES `movie`.`movie_infos`(`movie_id`) ON UPDATE CASCADE ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `movie`.`comments`;
CREATE TABLE `movie`.`comments`(
		`comment_id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `comment_text` LONGTEXT NOT NULL  comment 'comment to the post',
    `user_id` bigint unsigned NOT NULL comment 'who commented',
    `post_id` bigint unsigned NOT NULL comment 'which post',
    `comment_like` int NOT NULL DEFAULT 0 comment 'comment like',
    `create_time` timestamp DEFAULT current_timestamp,
    `update_time` timestamp DEFAULT current_timestamp ON UPDATE current_timestamp,
	PRIMARY KEY(`comment_id`),
	FOREIGN KEY(`user_id`) REFERENCES `movie`.`users`(`id`) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY(`post_id`) REFERENCES `movie`.`posts`(`post_id`) ON UPDATE CASCADE ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `movie`.`follow_user`;
CREATE TABLE `movie`.`follow_user`(
	`follow_id` BIGINT NOT NULL AUTO)INCREMENT,
	`user_id` BEGIN NOT NULL COMMENT 'who follow some one',
	`followed_user_id` BIGINT NOT NULL COMMENT 'who is being followed'
	PRIMARY KEY(`follow_id`),
	UNIQUE KEY `idx_userID_followedUserID`(`user_id`,`followed_user_id`) # 1 person will only follow one same person
);