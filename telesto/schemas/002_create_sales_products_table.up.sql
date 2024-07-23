CREATE TABLE sale_products (
    id SERIAL PRIMARY KEY,
    sale_id VARCHAR REFERENCES sales(id) ON DELETE CASCADE,
    product_id VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL,
    price INTEGER NOT NULL
);