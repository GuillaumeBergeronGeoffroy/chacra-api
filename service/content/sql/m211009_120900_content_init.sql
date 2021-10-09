CREATE TABLE content (
    content_id INT NOT NULL AUTO_INCREMENT,
    content_model VARCHAR(255) NOT NULL,
    content_model_id INT NOT NULL,
    content_model_type_id TINYINT NOT NULL,
    content_lang VARCHAR(255) NOT NULL,
    content_value TEXT NOT NULL,
    content_status TINYINT DEFAULT 0, 
    PRIMARY KEY (content_id),
    CONSTRAINT uidx_content_content_model_content_model_id_content_lang UNIQUE (content_model, content_model_id, content_lang)
);

-- TO CONST 
-- content_status 0 awaiting_submission, 1 awaiting_revision, 2 accepted 
-- content_model_type_id -> in each model

CREATE TABLE content_suggestion (
    content_suggestion_id INT NOT NULL AUTO_INCREMENT,
    content_id VARCHAR(255) NOT NULL,
    content_suggestion_value TEXT NOT NULL,
    PRIMARY KEY (content_suggestion_id)
);