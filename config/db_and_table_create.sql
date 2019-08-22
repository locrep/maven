DROP DATABASE IF EXISTS locrep;
CREATE DATABASE locrep;
use locrep;
CREATE TABLE maven_artifacts (
    artRowId INT NOT NULL AUTO_INCREMENT,
    //filename TINYTEXT NOT NULL,
    version VARCHAR(10),
    //md5Hash VARCHAR(32),
    //sha1Hash VARCHAR(40),
    groupID TINYTEXT,
    artifactID TINYTEXT,
    dateTime DATETIME DEFAULT CURRENT_TIMESTAMP,
    primary key (artRowId),
    uniq (groupid,artifactid,version)
    
);

CREATE TABLE maven_artifact_file (
    artRowId INT NOT NULL,
    filename TINYTEXT NOT NULL,
    checksum VARCHAR(40)
);