CREATE TABLE `employment` (
  `id` bigint unsigned PRIMARY KEY AUTO_INCREMENT,
  `profile_code` bigint unsigned NOT NULL,
  `job_title` varchar(50) NOT NULL,
  `employer` varchar(50) NOT NULL,
  `start_date` date NOT NULL,
  `end_date` date,
  `city` varchar(50) NOT NULL,
  `description` text,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp
);
CREATE INDEX `employment_index_3` ON `employment` (`profile_code`);

ALTER TABLE `employment` ADD CONSTRAINT `user_employment` FOREIGN KEY (`profile_code`) REFERENCES `profile` (`profile_code`);