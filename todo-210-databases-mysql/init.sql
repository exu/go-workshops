CREATE DATABASE `sqlx_test`;
USE `sqlx_test`;
CREATE TABLE `items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` longtext NOT NULL,
  PRIMARY KEY (`id`)
);
INSERT INTO `items` (name, description) VALUES ('mechanical keyboard', 'good for writing code');
INSERT INTO `items` (name, description) VALUES ('coffe', 'oh yummy!');
INSERT INTO `items` (name, description) VALUES ('emacs', 'your favourite text editor');
