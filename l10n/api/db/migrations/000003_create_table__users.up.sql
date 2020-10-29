CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITHOUT TIME ZONE,
    created_by BIGINT NOT NULL,
    updated_by BIGINT NOT NULL,
    deleted_by BIGINT,
    display_language_code CHARACTER(2) NOT NULL DEFAULT 'en',
    file_language_code CHARACTER(2) NOT NULL DEFAULT 'en',
    email CHARACTER VARYING NOT NULL,
    attributes JSONB
);