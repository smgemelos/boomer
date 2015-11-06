DROP DATABASE IF EXISTS `caprica`;
DROP USER IF EXIXTS 'caprica'@'%';

CREATE DATABASE `caprica`;
CREATE USER 'caprica'@'%' IDENTIFIED BY 'caprica';
GRANT ALL PRIVILEGES ON caprica.* TO 'caprica'@'%';
FLUSH PRIVILEGES;
USE caprica;

