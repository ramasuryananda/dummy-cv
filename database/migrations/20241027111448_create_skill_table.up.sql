CREATE TABLE `skill` (
  `id` bigint unsigned PRIMARY KEY AUTO_INCREMENT,
  `profile_code` bigint unsigned NOT NULL,
  `skill` varchar(20) NOT NULL,
  `level` varchar(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp
);

CREATE INDEX `skill_index_5` ON `skill` (`profile_code`);


ALTER TABLE `skill` ADD CONSTRAINT `user_skill` FOREIGN KEY (`profile_code`) REFERENCES `profile` (`profile_code`);
