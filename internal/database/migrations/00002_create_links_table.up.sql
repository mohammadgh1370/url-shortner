CREATE TABLE IF NOT EXISTS `links` (
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` bigint UNSIGNED NOT NULL,
    `url` longtext NOT NULL,
    `hash` varchar(191) NOT NULL,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    `unique_hash_md5` char(32) GENERATED ALWAYS AS (md5(concat(`hash`, ifnull(`deleted_at`, '')))) VIRTUAL,
    `unique_url_user_id_md5` char(32) GENERATED ALWAYS AS (md5(concat(`url`, `user_id`, ifnull(`deleted_at`, '')))) VIRTUAL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_hash_md5` (`unique_hash_md5`),
    UNIQUE KEY `unique_url_user_id_md5` (`unique_url_user_id_md5`),
    KEY `fk_links_user` (`user_id`),
    CONSTRAINT `fk_links_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
