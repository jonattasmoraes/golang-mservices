CREATE TABLE IF NOT EXISTS sale_reports (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    sale_id VARCHAR(255) NOT NULL,
    product_id VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    unit VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    price INT NOT NULL,
    total INT NOT NULL,
    payment_type VARCHAR(255) NOT NULL,
    sale_date TIMESTAMP NOT NULL
);

CREATE INDEX idx_sale_reports_user_id ON sale_reports(user_id);
CREATE INDEX idx_sale_reports_sale_id ON sale_reports(sale_id);
CREATE INDEX idx_sale_reports_product_id ON sale_reports(product_id);

