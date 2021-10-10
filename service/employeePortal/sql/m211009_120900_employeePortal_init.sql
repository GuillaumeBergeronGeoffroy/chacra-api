CREATE TABLE Employee (
    employeeId INT NOT NULL AUTO_INCREMENT,
    employeeEmail VARCHAR(255) NOT NULL,
    employeePassword VARCHAR(255) NOT NULL,
    employeeCreated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (employeeId),
    CONSTRAINT uidx_Employee_employeeEmail UNIQUE (employeeEmail)
);