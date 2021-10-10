CREATE TABLE User (
    userId INT NOT NULL AUTO_INCREMENT,
    userEmail VARCHAR(255) NOT NULL,
    userPassword VARCHAR(255) NOT NULL,
    userCreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (userId),
    CONSTRAINT uidx_User_userEmail UNIQUE userEmail)
);