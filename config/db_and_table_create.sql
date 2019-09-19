DROP DATABASE IF EXISTS locrep;
CREATE DATABASE locrep;
use locrep;
CREATE TABLE maven_artifacts (
    artRowId INT NOT NULL AUTO_INCREMENT,
    groupID TINYTEXT,
    artifactID TINYTEXT,
    version VARCHAR(10),
    dateTime DATETIME DEFAULT CURRENT_TIMESTAMP,
    primary key (artRowId),
    uniq (groupid,artifactid,version)
    
);

CREATE TABLE maven_artifact_file (
    id AUTO_INCREMENT
    artRowId INT NOT NULL,
    filename TINYTEXT NOT NULL,
    checksum VARCHAR(40)
);