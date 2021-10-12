CREATE TABLE Employee (
    EmployeeId INT NOT NULL AUTO_INCREMENT,
    EmployeeEmail VARCHAR(255) NOT NULL,
    EmployeePassword VARCHAR(255) NOT NULL,
    PRIMARY KEY (EmployeeId),
    CONSTRAINT uidx_Employee_EmployeeEmail UNIQUE (EmployeeEmail)
);