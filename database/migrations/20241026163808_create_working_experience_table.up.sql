CREATE TABLE `working_experience` (
  `id` bigint unsigned PRIMARY KEY AUTO_INCREMENT,
  `profile_code` bigint unsigned NOT NULL,
  `working_experience` text,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp
);

CREATE INDEX `working_experience_index_2` ON `working_experience` (`profile_code`);

ALTER TABLE `working_experience` ADD CONSTRAINT `user_working_experience` FOREIGN KEY (`profile_code`) REFERENCES `profile` (`profile_code`);