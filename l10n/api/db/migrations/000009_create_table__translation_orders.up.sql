CREATE TABLE IF NOT EXISTS translation_orders (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITHOUT TIME ZONE,
    created_by BIGINT NOT NULL,
    updated_by BIGINT NOT NULL,
    deleted_by BIGINT,
    file_version_id BIGINT NOT NULL,
    subtotal INTEGER NOT NULL,
    total INTEGER NOT NULL,
    translations JSONB
);