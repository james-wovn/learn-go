CREATE TABLE IF NOT EXISTS folders (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITHOUT TIME ZONE,
    created_by BIGINT NOT NULL,
    updated_by BIGINT NOT NULL,
    deleted_by BIGINT,
    project_id BIGINT NOT NULL,
    parent_id BIGINT,
    attributes JSONB
);