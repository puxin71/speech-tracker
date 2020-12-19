CREATE DATABASE IF NOT EXISTS test_db;
Use test_db;

DROP TABLE IF EXISTS talk CASCADE;

CREATE TABLE IF NOT EXISTS talk (
  id INT NOT NULL AUTO_INCREMENT,
  title	VARCHAR(255),
  abstract TINYTEXT,
  room SMALLINT UNSIGNED,
  PRIMARY KEY (id)
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS attendant (
  id	INT	NOT NULL AUTO_INCREMENT PRIMARY KEY,
  fname	VARCHAR(255) 	NOT NULL,
  lname	VARCHAR(255)	NOT NULL,
  company	VARCHAR(255) NOT NULL,
  email		VARCHAR(255) NOT NULL,
  local_register_time DATETIME,
  utc_register_time DATETIME,
  tz VARCHAR(64),	
  role ENUM('speaker', 'attendee') NOT NULL,
  bio	TINYTEXT,
  talk_id	INT	NOT NULL,
  
  INDEX t_ind 	(talk_id),
  
  FOREIGN KEY (talk_id)
  	REFERENCES talk(id)
  	ON DELETE CASCADE
) ENGINE = INNODB;