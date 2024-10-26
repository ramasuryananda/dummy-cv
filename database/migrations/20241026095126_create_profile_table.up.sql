CREATE TABLE `profile` (
  `profile_code` bigint unsigned PRIMARY KEY AUTO_INCREMENT,
  `wanted_job_title` varchar(255) NOT NULL,
  `first_name` varchar(50) NOT NULL,
  `last_name` varchar(50),
  `email` varchar(50),
  `phone` varchar(15),
  `country` varchar(20),
  `city` varchar(20),
  `address` text,
  `postal_code` varchar(20),
  `driving_license` varchar(30),
  `nationality` varchar(20) NOT NULL,
  `place_of_birth` varchar(30) NOT NULL,
  `date_of_birth` date NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp
);

CREATE INDEX `profile_index_0` ON `profile` (`profile_code`);
