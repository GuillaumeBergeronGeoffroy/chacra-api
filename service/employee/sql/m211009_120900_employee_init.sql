CREATE TABLE employee (
    employee_id INT NOT NULL AUTO_INCREMENT,
    employee_email VARCHAR(255) NOT NULL,
    employee_password VARCHAR(255) NOT NULL,
    employee_created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (employee_id)
);

CREATE UNIQUE INDEX uidx_employee_employee_email
ON employee (employee_email);