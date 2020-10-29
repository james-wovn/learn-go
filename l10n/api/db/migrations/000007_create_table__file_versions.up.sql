CREATE TABLE IF NOT EXISTS file_versions (
     id BIGSERIAL PRIMARY KEY,
     created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
     updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
     deleted_at TIMESTAMP WITHOUT TIME ZONE,
     created_by BIGINT NOT NULL,
     updated_by BIGINT NOT NULL,
     deleted_by BIGINT,
     file_id BIGINT NOT NULL,
     version SERIAL,
     is_current_version BOOLEAN NOT NULL DEFAULT FALSE,
     is_machine_translated BOOLEAN NOT NULL DEFAULT FALSE,
     is_human_translated BOOLEAN NOT NULL DEFAULT FALSE,
     attributes JSONB
);