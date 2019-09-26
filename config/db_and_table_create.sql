DROP DATABASE IF EXISTS locrep;
CREATE DATABASE locrep;
use locrep;
CREATE TABLE maven_artifacts (
    artRowID INT NOT NULL AUTO_INCREMENT,
    groupID TINYTEXT,
    artifactID TINYTEXT,
    version VARCHAR(10),
    dateTime DATETIME DEFAULT CURRENT_TIMESTAMP,
    primary key (artRowID),
    CONSTRAINT artifact UNIQUE (groupID,artifactID,version)

);

CREATE TABLE maven_artifact_file (
    id INT NOT NULL AUTO_INCREMENT,
    artRowID INT NOT NULL,
    filename TINYTEXT NOT NULL,
    checksum VARCHAR(40),
    primary key(id)
);
