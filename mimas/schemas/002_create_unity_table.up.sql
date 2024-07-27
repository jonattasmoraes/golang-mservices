CREATE TABLE IF NOT EXISTS Unit (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL
);

INSERT INTO Unit (ID, Name)
VALUES 
    (1, 'Litro'),
    (2, 'Kg'),
    (3, 'Unidade'),
    (4, 'Pacote'),
    (5, 'Lata'),
    (6, 'Grama'),
    (7, 'Colher'),
    (8, 'Rolo');