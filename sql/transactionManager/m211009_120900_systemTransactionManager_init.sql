CREATE TABLE SystemTransaction (
    systemTransactionId INT NOT NULL AUTO_INCREMENT,
    systemTransactionTypeId TINYINT NOT NULL, 
    userId INT,
    producerId INT,
    employeeId INT,
    systemTransactionCreatedAt DATETIME DEFAULT CURRENTTIMESTAMP,
    PRIMARY KEY (transactionId)
);

CREATE INDEX SystemTransaction_userId ON transaction (userId);
CREATE INDEX SystemTransaction_producerId ON transaction (producerId);
CREATE INDEX SystemTransaction_employeeId ON transaction (employeeId);
CREATE INDEX SystemTransaction_transactionCreatedAt ON transaction (transactionCreatedAt);

CREATE TABLE Order (
    orderId INT NOT NULL AUTO_INCREMENT,
    productId INT NOT NULL, 
    orderStatus TINYINT DEFAULT 0, 
    PRIMARY KEY (orderId)
);

CREATE INDEX Order_productId ON Order (productId);
-- CREATE INDEX order_orderCreatedAt ON order (orderCreatedAt);

CREATE TABLE OrderFullfilment (
    orderFullfilmentId INT NOT NULL AUTO_INCREMENT,
    orderFullfilmentOptionTypeId TINYINT DEFAULT 0,
    orderId INT NOT NULL,
    orderFullfilmentStatus TINYINT DEFAULT 0, 
    PRIMARY KEY (order_fullfilmentId),
    CONSTRAINT fkey_OrderFullfilment_orderId FOREIGN KEY (orderId)
    REFERENCES Order (orderId)
);

-- TO CONST
-- orderFullfilmentStatus 0 awaiting, 1 onThe_way, 2 confirmed 