DROP DATABASE IF EXISTS locrep;
CREATE DATABASE locrep;
use locrep;
CREATE TABLE maven (
    id INT NOT NULL AUTO_INCREMENT,
    filename TINYTEXT NOT NULL,
    version VARCHAR(10),
    md5Hash VARCHAR(32),
    sha1Hash VARCHAR(40),
    groupid TINYTEXT,
    artifactid TINYTEXT,
    dateTime DATETIME DEFAULT CURRENT_TIMESTAMP,
    primary key (id)
    
);