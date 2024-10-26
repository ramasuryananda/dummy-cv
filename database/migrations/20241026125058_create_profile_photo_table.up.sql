CREATE TABLE `profile_photo` (
  `id` bigint unsigned PRIMARY KEY AUTO_INCREMENT,
  `profile_code` bigint unsigned NOT NULL,
  `photo_url` text,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp
);
ALTER TABLE `profile_photo` ADD CONSTRAINT `name_optional` FOREIGN KEY (`profile_code`) REFERENCES `profile` (`profile_code`);
