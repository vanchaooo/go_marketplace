-- +goose Up
CREATE TABLE cart_items (
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    product_id BIGINT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    quantity BIGINT NOT NULL CHECK (quantity > 0),
    PRIMARY KEY(user_id, product_id)
);

-- +goose Down
DROP TABLE cart_items;
