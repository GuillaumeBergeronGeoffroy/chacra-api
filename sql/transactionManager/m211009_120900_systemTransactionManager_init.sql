CREATE TABLE SystemTransaction (
    SystemTransactionId INT NOT NULL AUTO_INCREMENT,
    SystemTransactionTypeId TINYINT NOT NULL, 
    UserId INT,
    ProducerId INT,
    EmployeeId INT,
    SystemTransactionCreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (SystemTransactionId)
);

CREATE INDEX SystemTransaction_UserId ON transaction (UserId);
CREATE INDEX SystemTransaction_ProducerId ON transaction (UserId);
CREATE INDEX SystemTransaction_EmployeeId ON transaction (EmployeeId);
CREATE INDEX SystemTransaction_TransactionCreatedAt ON transaction (TransactionCreatedAt);

CREATE TABLE `Order` (
    OrderId INT NOT NULL AUTO_INCREMENT,
    ProductId INT NOT NULL, 
    OrderStatus TINYINT DEFAULT 0, 
    PRIMARY KEY (OrderId)
);

CREATE INDEX Order_ProductId ON `Order` (ProductId);
-- CREATE INDEX order_orderCreatedAt ON order (orderCreatedAt);

CREATE TABLE OrderFullfilment (
    OrderFullfilmentId INT NOT NULL AUTO_INCREMENT,
    OrderFullfilmentOptionTypeId TINYINT DEFAULT 0,
    OrderId INT NOT NULL,
    OrderFullfilmentStatus TINYINT DEFAULT 0, 
    PRIMARY KEY (OrderFullfilmentId),
    CONSTRAINT fkey_OrderFullfilment_OrderId FOREIGN KEY (OrderId)
    REFERENCES `Order` (OrderId)
);

-- TO CONST
-- orderFullfilmentStatus 0 awaiting, 1 onThe_way, 2 confirmed 