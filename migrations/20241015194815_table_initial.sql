-- +goose Up
-- +goose StatementBegin
-- CREATE TABLE product_images(
--     product_id INTEGER,
--     url TEXT UNIQUE,

-- )

CREATE TABLE product(
    id SERIAL PRIMARY KEY,
    tittle VARCHAR,
    cost float,
    amount INTEGER,
    description TEXT
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE product;
-- +goose StatementEnd
