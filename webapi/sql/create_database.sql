CREATE DATABASE IF NOT EXISTS `genshindb`;

DROP USER IF EXISTS `paimon`@`%`;
CREATE USER paimon IDENTIFIED BY 'paimon';
GRANT ALL PRIVILEGES ON genshindb.* TO 'paimon'@'%';