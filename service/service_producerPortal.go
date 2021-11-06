package service

import (
	"net/http"
	"sync"
)

type producerPortal struct {
	Dao *Dao
}

var ppOnce sync.Once
var pp producerPortal

// ProducerPortal exportable singleton
func ProducerPortal(dao *Dao) *producerPortal {
	ppOnce.Do(func() {
		pp = producerPortal{dao}
		ExecuteStatements(pp.Dao.DB, ppInitSql)
	})
	return &pp
}

// ProducerPortalActions exportable
func (m producerPortal) Actions() (ac Actions, err error) {
	ac = map[string]Action{
		"createProducer": func(w http.ResponseWriter, r *http.Request) {
			// reqBody := u.Read(w, r)
			// resBody, err := subscribe(reqBody, m)
			// if err != nil {
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			// return
			// }
			// u.Write(w, r, resBody)
			w.WriteHeader(500)
		},
	}
	return
}

// -- TO CONST
// -- producerStatus 0 awaitingSubmission, 1 awaitingRevision, 2 accepted

var ppInitSql = []string{
	`CREATE TABLE Producer (
		ProducerId INT NOT NULL AUTO_INCREMENT,
		ProducerEmail VARCHAR(255) NOT NULL,
		ProducerPassword VARCHAR(255),
		ProducerCreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		ProducerStatus TINYINT DEFAULT 0,
		PRIMARY KEY (ProducerId),
		CONSTRAINT uidx_Producer_ProducerEmail UNIQUE (ProducerEmail)
	);`,
	`CREATE TABLE Product (
		ProductId INT NOT NULL AUTO_INCREMENT,
		ProducerId INT NOT NULL,
		ProductStatus TINYINT DEFAULT 0,
		PRIMARY KEY (ProductId),
		CONSTRAINT fkey_Product_ProducerId FOREIGN KEY (ProducerId)
		REFERENCES Producer (ProducerId)
	);`,
	`CREATE TABLE ProductAvailability (
		ProductAvailabilityId INT NOT NULL AUTO_INCREMENT,
		ProductId INT NOT NULL,
		ProductAvailabilityQuantity INT NOT NULL,
		ProductAvailabilityStart DATETIME DEFAULT CURRENT_TIMESTAMP,
		ProductAvailabilityEnd DATETIME DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (ProductAvailabilityId),
		CONSTRAINT fkey_ProductAvailability_ProductId FOREIGN KEY (ProductId)
		REFERENCES Product (ProductId)
	);`,
	`CREATE TABLE ProductFullfilmentOption (
		ProductFullfilmentOptionId INT NOT NULL AUTO_INCREMENT,
		ProductFullfilmentOptionTypeId TINYINT NOT NULL,
		ProductId INT NOT NULL,
		PRIMARY KEY (ProductFullfilmentOptionId),
		CONSTRAINT fkey_ProductFullfilmentOption_ProductId FOREIGN KEY (ProductId)
		REFERENCES Product (ProductId)
	);`,
}
