CREATE TABLE transaction (
    transaction_id INT NOT NULL AUTO_INCREMENT,
    transaction_type_id TINYINT NOT NULL, 
    transaction_created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    user_id INT,
    producer_id INT,
    employee_id INT,
    PRIMARY KEY (order_id)
);

CREATE TABLE order (
    order_id INT NOT NULL AUTO_INCREMENT,
    product_id INT NOT NULL, 
    order_status TINYINT DEFAULT 0, 
    order_created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (order_id)
);

CREATE INDEX order_product_id ON order (product_id);

CREATE TABLE order_fullfilment (
    order_fullfilment_id INT NOT NULL AUTO_INCREMENT,
    order_id INT NOT NULL,
    order_fullfilment_status TINYINT DEFAULT 0, 
    order_fullfilment_option_type_id TINYINT DEFAULT 0,
    order_fullfilment_created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (order_fullfilment_id),
    CONSTRAINT fkey_order_fullfilment_order_id FOREIGN KEY (order_id)
    REFERENCES order (order_id)
);

-- TO CONST
-- order_fullfilment_status 0 awaiting, 1 on_the_way, 2 confirmed 