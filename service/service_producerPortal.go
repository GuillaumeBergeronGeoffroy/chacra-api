package service

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	model "github.com/GuillaumeBergeronGeoffroy/chacra-api/model"
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
			producer := &model.Producer{}
			reqBody := Read(w, r)
			if err = json.Unmarshal([]byte(reqBody), producer); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			producer.ProducerCreatedAt = time.Now().Format("2006-01-02 15:04:05")
			if err = SaveModel(producer, m.Dao.DB); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(ComposeResponse(w, map[string]string{
				"message": "Votre place est réservé!",
				"success": "true",
			}))
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
		ProducerName VARCHAR(255), 
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
