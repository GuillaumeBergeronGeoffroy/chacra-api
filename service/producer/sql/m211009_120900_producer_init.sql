CREATE TABLE producer (
    producer_id INT NOT NULL AUTO_INCREMENT,
    producer_email VARCHAR(255) NOT NULL,
    producer_password VARCHAR(255) NOT NULL,
    producer_created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    producer_status TINYINT DEFAULT 0,
    PRIMARY KEY (producer_id),
    CONSTRAINT uidx_producer_producer_email UNIQUE (producer_email)
);

-- TO CONST
-- producer_status 0 awaiting_submission, 1 awaiting_revision, 2 accepted 

CREATE TABLE product (
    product_id INT NOT NULL AUTO_INCREMENT,
    producer_id INT NOT NULL,
    product_created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    product_status TINYINT DEFAULT 0,
    PRIMARY KEY (product_id),
    CONSTRAINT fkey_product_producer_id FOREIGN KEY (producer_id)
    REFERENCES producer (producer_id)
);

-- TO CONST
-- producer_status 0 awaiting_submission, 1 awaiting_revision, 2 accepted 

CREATE TABLE product_availability (
    product_availability_id INT NOT NULL AUTO_INCREMENT,
    product_id INT NOT NULL,
    product_availability_quantity INT NOT NULL,
    product_availability_start INT NOT NULL,
    product_availability_created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (product_availability_id),
    CONSTRAINT fkey_product_delivery_option_product_id FOREIGN KEY (product_id)
    REFERENCES product (product_id)
);

CREATE TABLE product_fullfilment_option (
    product_fullfilment_option_id INT NOT NULL AUTO_INCREMENT,
    product_id INT NOT NULL,
    product_fullfilment_option_type_id INT NOT NULL,
    PRIMARY KEY (product_fullfilment_option_id),
    CONSTRAINT fkey_product_fullfilment_option_id_product_id FOREIGN KEY (product_id)
    REFERENCES product (product_id)
);