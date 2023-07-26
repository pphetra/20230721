-- +goose Up
-- +goose StatementBegin
CREATE TABLE members (
    id SERIAL PRIMARY KEY,
    name1 VARCHAR(255) NOT NULL,
    name2 VARCHAR(255) NOT NULL,
    member_type int NOT NULL,    
    email VARCHAR(255) NOT NULL,
    address_line1 VARCHAR(255) NOT NULL,
    address_line2 VARCHAR(255) NOT NULL,
    address_postal_code VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE members;
-- +goose StatementEnd
