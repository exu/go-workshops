CREATE DATABASE IF NOT EXISTS `gotraining`;
USE `gotraining`;

DROP TABLE  IF EXISTS `squares`;
CREATE TABLE `squares` (
  `num` int(11) NOT NULL,
  `square` int(11) NOT NULL,
  PRIMARY KEY (`num`)
);


DROP TABLE  IF EXISTS `items`;
CREATE TABLE `items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` longtext NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO `items` (name, description) VALUES ('mechanical keyboard', 'good for writing code');
INSERT INTO `items` (name, description) VALUES ('coffe', 'oh yummy!');
INSERT INTO `items` (name, description) VALUES ('emacs', 'your favourite text editor');
