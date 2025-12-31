/* 
SCRIPT IS FOR MySQL
Generate tables and data for the program 
*/

DROP DATABASE IF EXISTS `Songs_Database`;
CREATE DATABASE `Songs_Database`;
USE `Songs_Database`;

CREATE TABLE songs (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    artist VARCHAR(50),
    song_name VARCHAR(50),
    song_file_name VARCHAR(50),
    /* duration is in seconds */
    duration INT
);

INSERT INTO songs (artist, song_name, song_file_name, duration) VALUES
('Yorushika', 'Nautilus', 'nautilus.mp4', 262);