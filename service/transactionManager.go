package service

import (
	"sync"
)

type transactionManager struct {
	Dao *Dao
}

var tmOnce sync.Once
var tm transactionManager

// TransactionManager exportable singleton
func TransactionManager(dao *Dao) *transactionManager {
	tmOnce.Do(func() {
		tm = transactionManager{dao}
		InitServiceSqlDB(tm.Dao.DB, tmInitSql)
	})
	return &tm
}

// TransactionManagerActions exportable
func (m transactionManager) Actions() (ac Actions, err error) {
	return
}

// -- TO CONST
// -- orderFullfilmentStatus 0 awaiting, 1 onThe_way, 2 confirmed

var tmInitSql = []string{
	`CREATE TABLE SystemTransaction (
		SystemTransactionId INT NOT NULL AUTO_INCREMENT,
		SystemTransactionTypeId TINYINT NOT NULL, 
		UserId INT,
		ProducerId INT,
		EmployeeId INT,
		SystemTransactionCreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (SystemTransactionId)
	);`,
	`CREATE INDEX SystemTransaction_UserId ON SystemTransaction (UserId);`,
	`CREATE INDEX SystemTransaction_ProducerId ON SystemTransaction (ProducerId);`,
	`CREATE INDEX SystemTransaction_EmployeeId ON SystemTransaction (EmployeeId);`,
	`CREATE INDEX SystemTransaction_SystemTransactionCreatedAt ON SystemTransaction (SystemTransactionCreatedAt);`,
	`CREATE TABLE ` + "`Order`" + ` (
		OrderId INT NOT NULL AUTO_INCREMENT,
		ProductId INT NOT NULL, 
		OrderStatus TINYINT DEFAULT 0, 
		PRIMARY KEY (OrderId)
	);`,
	`CREATE INDEX Order_ProductId ON ` + "`Order`" + ` (ProductId);`,
	// -- CREATE INDEX order_orderCreatedAt ON order (orderCreatedAt);
	`CREATE TABLE OrderFullfilment (
		OrderFullfilmentId INT NOT NULL AUTO_INCREMENT,
		OrderFullfilmentOptionTypeId TINYINT DEFAULT 0,
		OrderId INT NOT NULL,
		OrderFullfilmentStatus TINYINT DEFAULT 0, 
		PRIMARY KEY (OrderFullfilmentId),
		CONSTRAINT fkey_OrderFullfilment_OrderId FOREIGN KEY (OrderId)
		REFERENCES ` + "`Order`" + ` (OrderId)
	);`,
	`SELECT * FROM OrderFullfilment;`,
}
