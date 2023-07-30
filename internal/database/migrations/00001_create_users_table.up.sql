CREATE TABLE IF NOT EXISTS `users` (
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    `first_name` varchar(191) NOT NULL,
    `last_name` varchar(191) NOT NULL,
    `username` varchar(191) NOT NULL,
    `password` varchar(191) NOT NULL,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;