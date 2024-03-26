-- +goose Up
-- +goose StatementBegin
    CREATE TABLE widgets (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        color TEXT
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE widgets;
-- +goose StatementEnd
