-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS files_metadata (
    id VARCHAR(100) PRIMARY KEY UNIQUE,
    owner_user_id varchar(100) REFERENCES users(id),
    file_name VARCHAR(100) NOT NULL,
    file_size bigint NOT NULL,
    path VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC') NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE files_metadata;
-- +goose StatementEnd
