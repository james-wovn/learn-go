CREATE TABLE IF NOT EXISTS file_segments (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITHOUT TIME ZONE,
    created_by BIGINT NOT NULL,
    updated_by BIGINT NOT NULL,
    deleted_by BIGINT,
    file_version_id BIGINT NOT NULL,
    source_language_code CHARACTER(2) NOT NULL,
    source_language_text TEXT NOT NULL,
    source_language_unit_count INTEGER NOT NULL,
    translations JSONB
);