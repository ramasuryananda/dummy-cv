CREATE TABLE `education` (
  `id` bigint unsigned PRIMARY KEY AUTO_INCREMENT,
  `profile_code` bigint unsigned NOT NULL,
  `school` varchar(100) NOT NULL,
  `degree` varchar(10) NOT NULL,
  `start_date` date NOT NULL,
  `end_date` date,
  `city` varchar(50) NOT NULL,
  `description` text,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp
);


CREATE INDEX `education_index_4` ON `education` (`profile_code`);


ALTER TABLE `education` ADD CONSTRAINT `user_education` FOREIGN KEY (`profile_code`) REFERENCES `profile` (`profile_code`);