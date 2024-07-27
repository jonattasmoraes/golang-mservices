CREATE TABLE IF NOT EXISTS Category (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL
);

INSERT INTO Category (ID, Name)
VALUES 
    (1, 'Bebidas'),
    (2, 'Alimentos'),
    (3, 'Limpeza'),
    (4, 'Mercearia'),
    (5, 'Higiene'),
    (6, 'Laticínios'),
    (7, 'Congelados'),
    (8, 'Diversos');