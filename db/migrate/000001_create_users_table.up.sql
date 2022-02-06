CREATE TABLE `users` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `first_name` VARCHAR(255) NOT NULL,
    `last_name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255),
    `created_at` TIMESTAMP ,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`)
    
) ENGINE=InnoDB CHARSET=utf8mb4;
