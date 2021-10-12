CREATE TABLE User (
    UserId INT NOT NULL AUTO_INCREMENT,
    UserEmail VARCHAR(255) NOT NULL,
    UserPassword VARCHAR(255) NOT NULL,
    PRIMARY KEY (UserId),
    CONSTRAINT uidx_User_UserEmail UNIQUE (UserEmail)
);