CREATE TABLE `users`
(
    `id`                              INT           NOT NULL AUTO_INCREMENT,
    `first_name`                      VARCHAR(255)  NOT NULL,
    `last_name`                       VARCHAR(255)  NOT NULL,
    `email`                           VARCHAR(255)  NOT NULL,
    `password`                        VARCHAR(5000) NOT NULL,
    `password_salt`                   VARCHAR(5000),
    `password_reset_token`            VARCHAR(255),
    `password_reset_token_expires_at` DATETIME,
    `role`                            INT           NOT NULL,
    `is_email_verified`               BOOLEAN       NOT NULL,
    `current_country`                 VARCHAR(255),
    `country_of_birth`                VARCHAR(255),
    `gender`                          VARCHAR(255),
    `timezone`                        VARCHAR(255),
    `birthday`                        DATE,
    `photo`                           TEXT,
    `email_two_fa_code`               VARCHAR(255),
    PRIMARY KEY (`id`),
    UNIQUE KEY `email_unique` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
