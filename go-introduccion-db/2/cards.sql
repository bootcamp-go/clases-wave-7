DROP DATABASE IF EXISTS bank_db;
CREATE DATABASE IF NOT EXISTS bank_db;

USE bank_db;

-- DDL
CREATE TABLE IF NOT EXISTS `owner` (
	`id` int unsigned NOT NULL AUTO_INCREMENT,
	`name` varchar(100) NOT NULL,
    PRIMARY KEY(`id`)
);

CREATE TABLE IF NOT EXISTS `cards` (
	`id` int unsigned NOT NULL AUTO_INCREMENT,
	`number` varchar(19) NOT NULL,
    `type` varchar(6) NOT NULL,
    `owner_id` int unsigned NOT NULL,
    PRIMARY KEY(`id`),
    KEY `idx_fk_owner_id` (`owner_id`),
    CONSTRAINT `idx_fk_owner_id` FOREIGN KEY(`owner_id`) REFERENCES `owner` (`id`)
);

-- DATASET
INSERT INTO owner (`name`) VALUES ("John Doe");
INSERT INTO owner (`name`) VALUES ("Jane Doe");
INSERT INTO owner (`name`) VALUES ("Addison");

INSERT INTO cards (`number`, `type`, `owner_id`) VALUES ("4444-5555-6666-7777", "DEBIT", 1);
INSERT INTO cards (`number`, `type`, `owner_id`) VALUES ("1111-1111-1111-1111", "DEBIT", 2);
INSERT INTO cards (`number`, `type`, `owner_id`) VALUES ("2222-2222-2222-2222", "CREDIT", 2);

