CREATE TABLE Content (
    contentId INT NOT NULL AUTO_INCREMENT,
    contentModel VARCHAR(255) NOT NULL,
    contentModelId INT NOT NULL,
    contentModelTypeId TINYINT NOT NULL,
    contentLang VARCHAR(255) NOT NULL,
    contentValue TEXT NOT NULL,
    contentStatus TINYINT DEFAULT 0, 
    PRIMARY KEY (content_id),
    CONSTRAINT uidx_Content_contentModel_contentModelId_contentModelTypeId_contentLang UNIQUE (contentModel, contentModelId, contentModelTypeId, contentLang)
);

-- TO CONST 
-- content_status 0 awaiting_submission, 1 awaiting_revision, 2 accepted 
-- content_model_type_id -> in each model

CREATE TABLE ContentSuggestion (
    contentSuggestionId INT NOT NULL AUTO_INCREMENT,
    contentId VARCHAR(255) NOT NULL,
    contentSuggestionBalue TEXT NOT NULL,
    PRIMARY KEY (contentSuggestionId),
    CONSTRAINT fkey_ContentSuggestion_contentId FOREIGN KEY (contentId)
    REFERENCES Content (contentId)
);