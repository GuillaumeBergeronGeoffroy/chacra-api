package service

import (
	"sync"
)

type contentManager struct {
	Dao *Dao
}

var cmOnce sync.Once
var cm contentManager

// ContentManager singleton exportable
func ContentManager(dao *Dao) *contentManager {
	cmOnce.Do(func() {
		cm = contentManager{dao}
		ExecuteStatements(cm.Dao.DB, cmInitSql)
	})
	return &cm
}

// ContentManagerActions exportable
func (m contentManager) Actions() (ac Actions, err error) {
	return
}

// -- TO CONST
// -- ContentStatus 0 awaitingSubmission, 1 awaitingRevision, 2 accepted
// -- ContentModelTypeId -> defined in each model

var cmInitSql = []string{
	`CREATE TABLE Content (
		ContentId INT NOT NULL AUTO_INCREMENT,
		ContentModel VARCHAR(255) NOT NULL,
		ContentModelId INT NOT NULL,
		ContentModelTypeId TINYINT NOT NULL,
		ContentLang VARCHAR(255) NOT NULL,
		ContentValue TEXT NOT NULL,
		ContentStatus TINYINT DEFAULT 0, 
		PRIMARY KEY (ContentId),
		CONSTRAINT uidx_Content_ContentLang UNIQUE (ContentModel, ContentModelId, ContentModelTypeId, ContentLang)
	);`,
	`CREATE TABLE ContentSuggestion (
		ContentSuggestionId INT NOT NULL AUTO_INCREMENT,
		ContentId INT NOT NULL,
		ContentSuggestionValue TEXT NOT NULL,
		PRIMARY KEY (ContentSuggestionId),
		CONSTRAINT fkey_ContentSuggestion_ContentId FOREIGN KEY (ContentId)
		REFERENCES Content (ContentId)
	);`,
}
