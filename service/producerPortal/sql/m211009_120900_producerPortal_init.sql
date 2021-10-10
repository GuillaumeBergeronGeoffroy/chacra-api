CREATE TABLE Producer (
    producerId INT NOT NULL AUTO_INCREMENT,
    producerEmail VARCHAR(255) NOT NULL,
    producerPassword VARCHAR(255) NOT NULL,
    producerCreated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    producerStatus TINYINT DEFAULT 0,
    PRIMARY KEY (producerId),
    CONSTRAINT uidx_Producer_producerEmail UNIQUE (producerEmail)
);

-- TO CONST
-- producerStatus 0 awaitingSubmission, 1 awaitingRevision, 2 accepted 

CREATE TABLE Product (
    productId INT NOT NULL AUTO_INCREMENT,
    producerId INT NOT NULL,
    productCreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    productStatus TINYINT DEFAULT 0,
    PRIMARY KEY (productId),
    CONSTRAINT fkey_Productt_producerId FOREIGN KEY (producerId)
    REFERENCES Producer (producerId)
);

-- TO CONST
-- producerStatus 0 awaitingSubmission, 1 awaitingRevision, 2 accepted 

CREATE TABLE ProductAvailability (
    productAvailabilityId INT NOT NULL AUTO_INCREMENT,
    productId INT NOT NULL,
    productAvailabilityQuantity INT NOT NULL,
    productAvailabilityStart INT NOT NULL,
    productAvailabilityCreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (productAvailabilityId),
    CONSTRAINT fkey_ProductAvailability_productId FOREIGN KEY (productId)
    REFERENCES Product (productId)
);

CREATE TABLE ProductFullfilmentOption (
    productFullfilment_option_id INT NOT NULL AUTO_INCREMENT,
    productId INT NOT NULL,
    productFullfilmentOptionTypeId INT NOT NULL,
    PRIMARY KEY (product_fullfilment_option_id),
    CONSTRAINT fkey_ProductFullfilmentOption_productId FOREIGN KEY (productId)
    REFERENCES Product (productId)
);