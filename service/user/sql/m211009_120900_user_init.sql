CREATE TABLE user (
    user_id INT NOT NULL AUTO_INCREMENT,
    user_email VARCHAR(255) NOT NULL,
    user_password VARCHAR(255) NOT NULL,
    user_created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id),
    CONSTRAINT uidx_user_user_email UNIQUE user_email)
);