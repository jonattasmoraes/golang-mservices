CREATE TABLE IF NOT EXISTS Category (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL
);

INSERT INTO Category (Name) VALUES ('Bebidas'), ('Doces'), ('Lanches');