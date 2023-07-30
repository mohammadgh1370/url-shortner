CREATE TABLE IF NOT EXISTS `views` (
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    `link_id` bigint UNSIGNED NOT NULL,
    `ip` varchar(191) NOT NULL,
    `referer` varchar(191) DEFAULT NULL,
    `user_agent` longtext,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `fk_views_link` (`link_id`),
    CONSTRAINT `fk_views_link` FOREIGN KEY (`link_id`) REFERENCES `links` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
